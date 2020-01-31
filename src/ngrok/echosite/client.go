// echosite api client
// author: lfreedomdev@gmail.com
// date: 2019/04/21 11:24:30
package echosite

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
)

var (
	esServer  string = ""
	ngServer  string = ""
	ngVersion string = ""
)

func Request(req EsReq) (*EsResp, error) {

	if req.Action == "" {
		return nil, fmt.Errorf("Action invalid")
	}

	request := EsReq{
		Action: req.Action,
		Data:   req.Data,
		Server: map[string]interface{}{
			"Domain":  ngServer,
			"Version": ngVersion,
		},
	}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(EsResp{}).
		Post(esServer)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("EchoSite Request Err,%d,%s", resp.StatusCode(), resp.Status())
	}

	res := &EsResp{}
	err = json.Unmarshal([]byte(resp.String()), res)
	if err != nil {
		return nil, fmt.Errorf("EchoSite data invalid")
	}

	return res, nil
}

func SetConfig(es_server string, ng_server string, ng_version string) {
	esServer = es_server
	ngServer = ng_server
	ngVersion = ng_version
}
