package battleservice

import (
	"fmt"
	apiclients "pokemon/apiClients"
	db "pokemon/database"
	heroes "pokemon/heroes"
	sqlpack "pokemon/sql"
)

type SelectedHero struct {
	HeroId         int      `json:"hero_id"`
	FighterName    string   `json:"fighter_name"`
	Powers         []string `json:"powers"`
	ExperienceLeve int      `json:"experience_level"`
	Weight         int      `json:"weight"`
	Height         int      `json:"height"`
	TotalPower     int      `json:"total_power"`
}

var (
	HeroData     heroes.HerosList
	HerosAbiliry heroes.HerosProperties
	Ability      heroes.AllAbilities

	FirstHero  SelectedHero
	SecondHero SelectedHero
	winner     *SelectedHero
	Target     *SelectedHero
)

func ShowHeroList() {
	apiclients.Api("https://pokeapi.co/api/v2/pokemon?offset=0&limit=1350", &HeroData)
}

func FindHeroDetails(first, second *SelectedHero) {
	for index, names := range HeroData.Results {
		if index+1 == first.HeroId || index+1 == second.HeroId {

			apiclients.Api(names.URL, &HerosAbiliry)

			if index+1 == first.HeroId {
				Target = first
			} else {
				Target = second
			}

			Target.FighterName = names.Name
			Target.ExperienceLeve = HerosAbiliry.BaseExperience
			Target.Weight = HerosAbiliry.Weight
			Target.Height = HerosAbiliry.Height

			Target.Powers = nil
			for _, p := range HerosAbiliry.Abilities {
				Target.Powers = append(Target.Powers, p.Ability.Name)
			}

			Target.TotalPower = 0
			for _, p := range Target.Powers {
				Target.TotalPower += findPower(p)
			}

			if err := SaveHeroToDB(Target); err != nil {
				fmt.Println("Error saving hero:", err)
			}
		}
	}
}

func heroesOptions() {
	if Ability.Count == 0 {
		apiclients.Api("https://pokeapi.co/api/v2/ability?offset=0&limit=370", &Ability)
	}
}

func findPower(name string) int {
	heroesOptions()
	for index, item := range Ability.Results {
		if item.Name == name {
			return index + 1
		}
	}
	return 0
}

func BattleService(hero *SelectedHero) {
	hero.TotalPower = 0
	for _, name := range hero.Powers {
		hero.TotalPower += findPower(name) + hero.Weight + hero.ExperienceLeve*5 + hero.Weight
	}
}

func SaveHeroToDB(hero *SelectedHero) error {
	if db.DB == nil {
		return fmt.Errorf("database not initialized")
	}

	var dbHero sqlpack.Hero
	err := db.DB.Where("name = ?", hero.FighterName).First(&dbHero).Error
	if err != nil {
		dbHero = sqlpack.Hero{
			Name:           hero.FighterName,
			Height:         hero.Height,
			Weight:         hero.Weight,
			BaseExperience: hero.ExperienceLeve,
		}
		db.DB.Create(&dbHero)
	}

	for _, abName := range hero.Powers {
		var ability sqlpack.Ability
		err := db.DB.Where("name = ?", abName).First(&ability).Error
		if err != nil {
			ability = sqlpack.Ability{
				Name:       abName,
				PowerIndex: findPower(abName),
			}
			db.DB.Create(&ability)
		}

		db.DB.Model(&dbHero).Association("Abilities").Append(&ability)
	}

	return nil
}

func SaveBattleToDB() error {
	if db.DB == nil {
		return fmt.Errorf("database not initialized")
	}

	var hero1, hero2 sqlpack.Hero
	if err := db.DB.Where("name = ?", FirstHero.FighterName).First(&hero1).Error; err != nil {
		return err
	}
	if err := db.DB.Where("name = ?", SecondHero.FighterName).First(&hero2).Error; err != nil {
		return err
	}

	battle := sqlpack.Battle{
		Hero1ID: hero1.ID,
		Hero2ID: hero2.ID,
	}
	return db.DB.Create(&battle).Error
}

func GetWinner() string {
	if FirstHero.TotalPower > SecondHero.TotalPower {
		winner = &FirstHero
	} else if SecondHero.TotalPower > FirstHero.TotalPower {
		winner = &SecondHero
	} else {
		winner = nil
	}

	if winner != nil {
		return winner.FighterName
	}
	return "draw"
}
