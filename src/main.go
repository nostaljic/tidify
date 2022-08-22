package main

import (
	"os"
	"tidify/devlog"
	"tidify/models"
	"tidify/router"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	rawEnv := os.Getenv(key)
	if len(rawEnv) == 0 {
		devlog.Debug("empty environment :", key)
		return rawEnv
	}
	return rawEnv
}
func main() {
	//err := godotenv.Load()
	devlog.SetLogLevel("Develop")
	err := godotenv.Load("tidify-api.env")
	USER := getEnv("USERNAME")
	PASS := getEnv("PASS")
	HOST := getEnv("HOST")
	PORT := getEnv("PORT")
	DBNAME := getEnv("DBNAME")
	if err != nil {
		devlog.Panic("File Importation Error Occured. ")
	}
	db := models.DBConnection(USER, PASS, HOST, PORT, DBNAME)

	router.SetupRoutes(db)
}
