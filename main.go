package main

import (
	"log"

	"github.com/temp-go-dev/sample-swagger/restapi"
	"github.com/temp-go-dev/sample-swagger/service"
)

var apiConfig restapi.Config

func main() {
	apiConfig.Debug = true
	apiConfig.Address = "localhost:8080"
	apiConfig.InsecureHTTP = true

	// serviceパッケージのSvcにインターフェースを実装してNewServerに渡す
	svc := &service.Svc{}
	api := restapi.NewServer(svc, &apiConfig)

	// サーバ起動
	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
