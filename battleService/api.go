package battleservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHeroesAPI(c *gin.Context) {
	ShowHeroList() 
	c.JSON(http.StatusOK, HeroData.Results)
}

func SelectHeroesAPI(c *gin.Context) {
	var req struct {
		FirstHeroID  int `json:"first_hero_id"`
		SecondHeroID int `json:"second_hero_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	FirstHero.HeroId = req.FirstHeroID
	SecondHero.HeroId = req.SecondHeroID

	FindHeroDetails(&FirstHero, &SecondHero)

	c.JSON(http.StatusOK, gin.H{
		"first_hero":  FirstHero,
		"second_hero": SecondHero,
	})
}

func BattleAPI(c *gin.Context) {
	BattleService(&FirstHero)
	BattleService(&SecondHero)

	winner := GetWinner()

	c.JSON(http.StatusOK, gin.H{
		"first_hero":  FirstHero,
		"second_hero": SecondHero,
		"winner":      winner,
	})
}
