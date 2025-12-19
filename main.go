package main

import (
	battleservice "pokemon/battleService"
	"pokemon/database"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.InitDB(
		"host=localhost user=postgres password=123456789 dbname=pokemon port=5432 sslmode=disable",
	); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/heroes", battleservice.ShowHeroesAPI)
	r.POST("/select", battleservice.SelectHeroesAPI)
	r.GET("/battle", battleservice.BattleAPI)

	r.Run(":8081")
}
