package main

import (
	"path/filepath"
	"os"
	"log"
	"github.com/metooweb/gateway/storage"
)

var ROOT_PATH string

func main() {

	initConfig()
	storage.Init()
	initHydraSDK()
	initServer()

}

func init() {

	ROOT_PATH = filepath.Dir(os.Args[0])
	file, _ := os.OpenFile(ROOT_PATH+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(file)

}
