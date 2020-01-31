// echosite models
// author: lfreedomdev@gmail.com
// date: 2019/04/21 11:24:30
package echosite

// user status info
type EsUser struct {
	Id    string
	Token string
}

// echosite request pack
type EsReq struct {
	Action string      `json:"Action"`
	Data   interface{} `json:"Data"`
	Server interface{} `json:"Server"`
}

// echosite response pack
type EsResp struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"errcode"`
	Message string      `json:"message"`
}
