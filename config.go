package main

import (
	"github.com/pkg/errors"
	"fmt"
	"github.com/go-ini/ini"
)

type Config struct {
	Addr string `ini:"addr"`
	Client struct {
		Id          string `ini:"id"`
		Secret      string `ini:"secret"`
		EndpointURL string `ini:"endpoint_url"`
	} `ini:"client"`
}

var config = new(Config)

func initConfig() (err error) {

	if err = ini.MapTo(config, ROOT_PATH+"/config.ini"); err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(config)

	return

}
