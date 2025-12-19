package main

import (
	battleservice "pokemon/battleService"
	"pokemon/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	if err := database.InitDB(dsn); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/heroes", battleservice.ShowHeroesAPI)
	r.POST("/select", battleservice.SelectHeroesAPI)
	r.GET("/battle", battleservice.BattleAPI)

	r.Run(":8081")
}
