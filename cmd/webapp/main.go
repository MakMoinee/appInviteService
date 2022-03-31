package main

import (
	"log"

	"github.com/MakMoinee/appInviteService/cmd/webapp/config"
	"github.com/MakMoinee/appInviteService/cmd/webapp/routes"
	"github.com/MakMoinee/appInviteService/internal/appInviteService/common"
	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {

	log.Println("....")
	config.Set()

	httpService := goserve.NewService(common.SERVER_PORT)
	httpService.EnableProfiling(common.ENABLE_PROFILING) // for profiling code

	routes.Set(httpService)
	log.Println("Server Started in port ")
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
