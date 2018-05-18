package storage

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Client struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	BackendURL string `json:"backend_url"`
}

type Api struct {
	Name     string  `json:"name"`
	Client   *Client `json:"client"`
	ClientId int     `json:"client_id"`
}

type Storage struct {
}

func (t *Storage) GetApi(name string) *Api {

	return apiList[name]
}

var apiList = make(map[string]*Api)

func Init() {

	var (
		err     error
		bytes   []byte
		clients = []*Client{}
		apis    = []*Api{}
	)

	fmt.Println("..........")

	bytes, err = ioutil.ReadFile("./clients.json")
	dealError(err)

	err = json.Unmarshal(bytes, &clients)
	dealError(err)

	bytes, err = ioutil.ReadFile("./apis.json")

	dealError(err)

	json.Unmarshal(bytes, &apis)

	for _, api := range apis {

		for _, client := range clients {
			if api.ClientId == client.Id {
				api.Client = client
				break
			}
		}

		apiList[api.Name] = api
	}

	fmt.Println(apiList)

}

func dealError(err error) {
	if err != nil {
		panic(err)
	}
}
