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

	var envVariable string = "MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION"

	serverVariables := []string{"host", "port", "user", "password"}
	queueNames := []string{"incoming", "outgoing"}
	queueVariables := []string{"name"}
	originVariables := []string{"name"}

	viper := viperLib.New()

	//Look for config file location defined as env var
	viper.BindEnv(envVariable)
	configFileLocation = viper.GetString(envVariable)
	if configFileLocation == "" {
		// Get config file from default location
		configFileLocation = "/etc/music-manager/"
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		return config, errors.New(errors.New("Fatal error reading config file: ").Error() + err.Error())
	}

	for _, server_variable := range serverVariables {
		if !viper.IsSet("server." + server_variable) {
			return config, errors.New("Fatal error config: no server " + server_variable + " was found.")
		}
	}

	for _, queue := range queueNames {
		for _, variable := range queueVariables {
			if !viper.IsSet(queue + "." + variable) {
				return config, errors.New("Fatal error config: no " + queue + " server " + variable + " variable was found.")
			}
		}
	}

	for _, origin_variable := range originVariables {
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
