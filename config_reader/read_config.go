package config

import (
	"errors"
	viperLib "github.com/spf13/viper"
)

type Server struct {
	User     string
	Password string
	Host     string
	Port     int
}

type Queue struct {
	Name             string
	Durable          bool
	DeleteWhenUnused bool
	Exclusive        bool
	NoWait           bool
	NoLocal          bool
	AutoACK          bool
}

type Config struct {
	Server   Server
	Incoming Queue
	Outgoing Queue
	Origin   string
}

func ReadConfig() (Config, error) {
	var configFileLocation string
	var config Config

	server_variables := []string{"host", "port", "user", "password"}
	queue_names := []string{"incoming", "outgoing"}
	queue_variables := []string{"name"}
	origin_variables := []string{"name"}

	viper := viperLib.New()

	//Look for config file location defined as env var
	viper.BindEnv("MUSIC_MANAGER_METAL_ARCHIVES_WRAPPER_CONFIG_FILE_LOCATION")
	configFileLocation = viper.GetString("MUSIC_MANAGER_METAL_ARCHIVES_WRAPPER_CONFIG_FILE_LOCATION")
	if configFileLocation == "" {
		// Get config file from default location
		configFileLocation = "/etc/music-manager-metal-archives-wrapper/"
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.New(errors.New("Fatal error reading config file: ").Error() + err.Error())
	}

	for _, server_variable := range server_variables {
		if !viper.IsSet("server." + server_variable) {
			return config, errors.New("Fatal error config: no server " + server_variable + " was found.")
		}
	}

	for _, queue := range queue_names {
		for _, variable := range queue_variables {
			if !viper.IsSet(queue + "." + variable) {
				return config, errors.New("Fatal error config: no " + queue + " server " + variable + " variable was found.")
			}
		}
	}

	for _, origin_variable := range origin_variables {
		if !viper.IsSet("origin." + origin_variable) {
			return config, errors.New("Fatal error config: no origin " + origin_variable + " was found.")
		}
	}

	server := Server{User: viper.GetString("server.user"), Password: viper.GetString("server.password"), Host: viper.GetString("server.host"), Port: viper.GetInt("server.port")}
	incoming := Queue{Name: viper.GetString("incoming.name")}
	outgoing := Queue{Name: viper.GetString("outgoing.name")}

	config.Server = server
	config.Incoming = incoming
	config.Outgoing = outgoing
	config.Origin = viper.GetString("origin.name")

	return config, nil
}
