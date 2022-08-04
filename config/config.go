package config

//import log "github.com/sirupsen/logrus"

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBTable    string
	DBPort     string
}

func NewDBConfig() *DBConfig {
	/* err := godotenv.Load()
	if err != nil {
		log.Error("Error loading env vars")
	} */
	return &DBConfig{
		DBUsername: "postgres",
		DBPassword: "sql1234",
		DBHost:     "localhost",
		DBTable:    "postgres",
		DBPort:     "5432",
	}
}

//asdasdasd
