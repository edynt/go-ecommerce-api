package initialize

import (
	"fmt"

	"github.com/edynt/go-ecommerce-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %v", err))
	}

}
