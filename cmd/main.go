package main

import (
	"log"

	userRoutes "github.com/akmyrat/project1/internal/users/routes"
	"github.com/akmyrat/project1/pkg/database/dbcon"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	// Load .env file
	err := godotenv.Load(".env") // Make sure the .env file is correctly loaded
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load configuration
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// Connect to DB
	DB, err := dbcon.ConnectToDB(dbcon.Config{
		Host:     viper.GetString("DB.host"),
		Port:     viper.GetString("DB.port"),
		Username: viper.GetString("DB.username"),
		DBName:   viper.GetString("DB.dbname"),
		SSLMode:  viper.GetString("DB.sslmode"),
		// Password: os.Getenv("DB.password"),
		Password: viper.GetString("DB.password"),
	})

	if err != nil {
		log.Fatalf("failed to initialize DB: %s", err.Error())
	}

	defer DB.Close()

	app := gin.Default()
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api, DB)
	// categoryRoutes.InitCategoryRoutes(api, DB)
	// postRoutes.InitPostRoutes(api, DB)

	log.Println("Starting the app in release mode on localhost:8000")

	if err := app.Run("localhost:8000"); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}

}

func initConfig() error {
	viper.AutomaticEnv()
	viper.AddConfigPath("Configs")
	viper.SetConfigName(("config"))
	return viper.ReadInConfig()
}
