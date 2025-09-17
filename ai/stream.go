package ai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/types"
	openai "github.com/sashabaranov/go-openai"
)

// StreamChat 阿里云DashScope流式对话对象
type StreamChat struct {
	apiKey       string
	baseURL      string
	model        string
	client       *openai.Client
	writeHandler func(string)
}

// StreamEvent 流式事件结构
type StreamEvent struct {
	Content string // 增量内容
	Err     error  // 错误信息
}

var aliBaseURL = "https://dashscope.aliyuncs.com/compatible-mode/v1"

func NewAliStreamChat(apiKey, model string, writeHandler ...func(string)) *StreamChat {
	return NewStreamChat(apiKey, aliBaseURL, model, writeHandler...)
}

// NewStreamChat 构造函数
func NewStreamChat(apiKey, baseURL, model string, writeHandler ...func(string)) *StreamChat {
	config := openai.DefaultConfig(apiKey)

	config.BaseURL = types.GetString(baseURL, aliBaseURL)
	config.HTTPClient = &http.Client{
		Timeout: 30 * time.Minute,
	}
	chat := &StreamChat{
		apiKey:  apiKey,
		baseURL: types.GetString(baseURL, aliBaseURL),
		model:   model,
		client:  openai.NewClientWithConfig(config),
	}
	if len(writeHandler) > 0 {
		chat.writeHandler = writeHandler[0]
	}
	return chat
}
func (s *StreamChat) GetResponse(prompt string) (string, error) {
	var result strings.Builder
	events, err := s.SubscribeToMessages(prompt)
	if err != nil {
		return "", err
	}
	for event := range events {
		if event.Err != nil {
			return "", event.Err
		}
		result.WriteString(event.Content)
	}

	return result.String(), err
}

// GetJsonResponse 获取完整响应（阻塞式）
func (s *StreamChat) GetJsonResponse(prompt string) (string, error) {

	data, err := s.GetResponse(prompt)
	if err != nil {
		return "", err
	}

	msg := strings.TrimSpace(data)
	msg = strings.ReplaceAll(msg, "\n", "")
	msg = strings.ReplaceAll(msg, " ", "")
	msg = strings.ReplaceAll(msg, "```json", "")
	msg = strings.ReplaceAll(msg, "```", "")
	msg = strings.TrimSpace(msg)

	return strings.TrimPrefix(msg, "json"), err
}

// SubscribeToMessages 流式订阅方法
func (s *StreamChat) SubscribeToMessages(prompt string) (<-chan StreamEvent, error) {
	ch := make(chan StreamEvent)
	ctx := context.Background()
	req := s.createRequest(prompt)

	stream, err := s.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("连接失败: %v", err)
	}

	go func() {
		defer close(ch)
		defer stream.Close()

		for {
			response, err := stream.Recv()
			switch {
			case errors.Is(err, io.EOF):
				return
			case err != nil:
				ch <- StreamEvent{Err: fmt.Errorf("接收错误: %v", err)}
				return
			default:
				if len(response.Choices) > 0 {
					msg := strings.TrimSpace(response.Choices[0].Delta.Content)
					ch <- StreamEvent{Content: msg}
					if s.writeHandler != nil {
						s.writeHandler(msg)
					}
				}
			}
		}
	}()

	return ch, nil
}

// 创建通用请求
func (s *StreamChat) createRequest(prompt string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: s.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Stream: true,
	}
}
