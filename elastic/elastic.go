package elastic

import (
	"encoding/json"
	"fmt"

	"github.com/mattbaird/elastigo/lib"
)

//ElasticSearch es组件
type ElasticSearch struct {
	host []string
	conn *elastigo.Conn
}

//ESConfigOption 配置文件
type ESConfigOption struct {
	Host []string `json:"hosts"`
}

//NewJSON 根据JSON配置文件初始化ElasticSearch组件
func NewJSON(config string) (host *ElasticSearch, err error) {
	var esco ESConfigOption
	err = json.Unmarshal([]byte(config), &esco)
	if err != nil {
		return
	}
	return New(esco)
}

//New 根据配置文件初始化ElasticSearch组件
func New(esco ESConfigOption) (host *ElasticSearch, err error) {
	host = &ElasticSearch{}
	host.conn = elastigo.NewConn()
	host.host = esco.Host
	host.conn.SetHosts(esco.Host)
	return
}

//Create 创建索引
func (host *ElasticSearch) Create(name string, typeName string, jsonData string) (id string, err error) {
	response, err := host.conn.Index(name, typeName, "", nil, jsonData)
	if err != nil {
		return
	}
	id = response.Id
	host.conn.Flush()
	if response.Created {
		return
	}
	err = fmt.Errorf("/%s/%s create error:%+v", name, typeName, response)
	return
}

//Update 更新索引
func (host *ElasticSearch) Update(name string, typeName string, id string, jsonData string) (err error) {
	response, err := host.conn.Index(name, typeName, id, nil, jsonData)
	if err != nil {
		return
	}
	host.conn.Flush()
	if response.Ok {
		return
	}
	// err = errors.New(fmt.Sprintf("/%s/%s update error:%+v", name, typeName, response))
	err = fmt.Errorf("/%s/%s update error:%+v", name, typeName, response)
	return
}

//Search 搜索数据
func (host *ElasticSearch) Search(name string, typeName string, query string) (result string, err error) {
	out, err := host.conn.Search(name, typeName, nil, query)
	if err != nil {
		return
	}
	var resultLst []*json.RawMessage
	for i := 0; i < len(out.Hits.Hits); i++ {
		resultLst = append(resultLst, (out.Hits.Hits[i].Source))
	}
	buffer, err := json.Marshal(&resultLst)
	if err != nil {
		return
	}
	result = string(buffer)
	return
}
