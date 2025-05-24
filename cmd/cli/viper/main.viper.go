package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstruct:"port"`
	} `mapstruct:"server"`
	Databases []struct {
		User     string `mapstruct:"user"`
		Password string `mapstruct:"password"`
		Host     string `1mapstruct:"host"`
	} `mapstruct:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to folder config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")

	// reaf file config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %s", err)) // panic: stop the program
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Security jwt::", viper.GetString("security.jwt.key"))

	// config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %v", err))
	}

	fmt.Println("Server Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("User: %s, Password: %s, Host: %s\n", db.User, db.Password, db.Host)
	}
}
