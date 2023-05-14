package main

import (
	"msvc-template/internal/api"
	"msvc-template/internal/client"
	"msvc-template/internal/config"
	"msvc-template/internal/util"
)

func main() {
	// Configs
	config.InitConfigs()
	client.InitClients()

	// Initializations
	fiberApp := api.NewFiberApp()

	// HTTP REST API V1
	api.SetUpRESTV1(fiberApp)

	// Starting HTTP server
	util.Logger.Info().Msg("HTTP Server is ON")
	err := fiberApp.Listen(config.Master.HTTP.Address)
	if err != nil {
		util.Logger.Error().Msg(err.Error())
		util.Logger.Panic().Msg("HTTP Server failed!")
	}
}
