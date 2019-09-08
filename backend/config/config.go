package config

import "github.com/spf13/viper"

var config *viper.Viper
var defaults = map[string]string{
	"env":                    "development",
	"db_host":                "localhost",
	"db_slave_host":          "db",
	"db_name":                "mydb",
	"db_user":                "go_user",
	"db_password":            "gonza1026",
	"payment_api_url":        "http://13.115.34.89/",
	"payment_api_key":        "apikey-test",
	"payment_api_agent":      "veritrans",
	"payment_api_service_id": "dr",
}

func Init() {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetEnvPrefix("simple_event_app")
	v.AutomaticEnv()
	config = v
}

func GetConfig() *viper.Viper {
	return config
}
