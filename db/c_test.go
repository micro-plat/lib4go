package db

import (
	"testing"

	"github.com/micro-plat/lib4go/ut"
)

func TestParse(t *testing.T) {
	list := []struct {
		input  string
		tp     string
		result string
		err    error
	}{
		{
			input:  "hydra:123456@hydra",
			tp:     "oracle",
			result: "hydra/123456@hydra",
			err:    nil,
		}, {
			input:  "hydra:123456@hydra/192.168.0.1",
			tp:     "mysql",
			result: "hydra:123456@tcp(192.168.0.1)/hydra",
			err:    nil,
		},
	}

	for _, item := range list {
		r, err := ParseConnectString(item.input, item.tp)
		ut.Expect(t, r, item.result)
		ut.Expect(t, err, item.err)
	}

}
