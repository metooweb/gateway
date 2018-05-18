package main

type Client struct {
	Id          string
	EndPointURL string
	//事件
	//钩子
}

type ApiParam struct {
	Name    string //参数名
	Type    string //参数类型 int | string | bool | map[int]string | map[int]int | map[string]string | map[string]int | array[int] | array[string]
	Require bool   //是否必传
}

type Api struct {
	Name   string
	Params map[string]*ApiParam //api参数
	Client *Client
}

func GetApiByName(name string) *Api {

	return nil
}
