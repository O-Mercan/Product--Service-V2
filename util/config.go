package util

//"log"

//"github.com/spf13/viper"

//Config stores all configuration of the application.
//The vvalues are read by viper from a config file or environment variables.

/* type Config struct {
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_HOST     string `mapstructure:"DB_PASSWORD"`
	DB_TABLE    string `mapstructure:"DB_TABLE"`
	DB_PORT     string `mapstructure:"DB_PORT"`
}

//LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading confgi file %s", err)
	}

	err = viper.Unmarshal(&config)
	return
} */
