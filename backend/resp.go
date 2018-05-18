package backend

import (
	"net/http"
	"encoding/json"
)

type Resp struct {
	Code   string      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func (data *Resp) Output(w http.ResponseWriter, ok bool, res interface{}) {

	if ok {
		data.Code = "SUCCESS"
		data.Result = res
	} else {

		switch res.(type) {
		case *ParamError:
			err := res.(*ParamError)
			data.Msg = err.Msg
			data.Code = err.Code()
		case *LogicError:
			err := res.(*LogicError)
			data.Msg = err.Msg
			data.Code = err.Code
		default:
			data.Code = "unknown"
			data.Msg = "system error"
		}
	}

	w.WriteHeader(http.StatusOK)

	bytes, _ := json.Marshal(data)

	w.Write(bytes)

}
