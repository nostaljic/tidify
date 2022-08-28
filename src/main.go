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
	devlog.SetLogLevel("Develop")
	path, err := os.Getwd()
	if err != nil {
		devlog.Debug(err)
	}
	devlog.Debug("[PATH]", path)
	//err := godotenv.Load()
	err = godotenv.Load("tidify.env")
	USER := getEnv("USRN")
	PASS := getEnv("PASS")
	HOST := getEnv("HOST")
	PORT := getEnv("PORT")
	DBNAME := getEnv("DBNAME")
	if err != nil {
		devlog.Panic("File Importation Error Occured. ")
	}

	devlog.Debug("[USERINFO]", USER, PASS, HOST, PORT, DBNAME)
	db := models.DBConnection(USER, PASS, HOST, PORT, DBNAME)

	router.SetupRoutes(db)
}
