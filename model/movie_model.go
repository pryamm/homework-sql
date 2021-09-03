package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"type:char(255);NOT NULL"`
	Slug        string `gorm:"type:char(255);unique;NOT NULL"`
	Description string `gorm:"type:text;NOT NULL"`
	Duration    int    `gorm:"type:integer(5);NOT NULL"`
	Image       string `gorm:"type:char(255);NOT NULL"`
}
