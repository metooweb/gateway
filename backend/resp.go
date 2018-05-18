package backend

type Resp struct {
	Code   string      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}
