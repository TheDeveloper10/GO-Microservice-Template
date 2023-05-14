package config

import "msvc-template/internal/util"

const (
	ServiceConfigPath = "config/service.json"
)

var (
	Master *MasterConfig
)

func InitConfigs() {
	Master = &MasterConfig{}
	loadConfig(ServiceConfigPath, &Master)

	util.Logger.Info().Msg("Loaded all configs")
}
