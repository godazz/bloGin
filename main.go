package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godazz/goGin/config"
	"github.com/spf13/viper"
)

func main() {
	configs := setConfig()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}

func setConfig() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading the config")
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decote configuration %v\n", err)
	}
	return configs
}
