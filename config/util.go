package config

import (
	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
)

//read the config file, helper function
func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func Load() *viper.Viper {d
	// Configs
	Config, err := readConfig("api_config", map[string]interface{}{
		"port":        9090,
		"hostname":    "localhost",
		"environment": "debug",
		"instance": map[string]string{
			"wallet_address": "0x5c9b022d100a8947e614bbfd232136077bc7c456d0",
			"wallet_seed":    "night care ten tomorrow ....",
		},
	})
	if err != nil {
		log.Error("Error when reading config: %v\n", err)
	}
	return Config
}
