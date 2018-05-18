package main

import (
	"github.com/ory/hydra/sdk/go/hydra"
)

var hydraSDK *hydra.CodeGenSDK

//token验证
func initHydraSDK() (err error) {

	hydraSDK, err = hydra.NewSDK(&hydra.Configuration{
		EndpointURL:  config.Client.EndpointURL,
		ClientID:     config.Client.Id,
		ClientSecret: config.Client.Secret,
		Scopes:       []string{"hydra.*"},
	})

	return
}
