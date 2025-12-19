package sql

import "gorm.io/gorm"

type Hero struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"unique"`
	Height         int
	Weight         int
	BaseExperience int
	Abilities      []Ability `gorm:"many2many:hero_abilities;"`
}

type Ability struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	PowerIndex int
}

type Battle struct {
	ID      uint `gorm:"primaryKey"`
	Hero1ID uint
	Hero2ID uint
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Hero{}, &Ability{}, &Battle{})
}
