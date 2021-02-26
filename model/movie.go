package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	Id     string `gorm:"varchar(64);primary_key; comment:'movie id'"`
	Title  string `gorm:"varchar(64); comment:'movie title'"`
	Year   string `gorm:"varchar(64); comment:'publish year'"`
	Info   string `gorm:"varchar(64); comment:'movie info'"`
	Rating string `gorm:"varchar(64); comment:'movie rating'"`
	Url    string `gorm:"varchar(64); comment:'movie url'"`
}

func SaveMovieInfo(db *gorm.DB, id string, title string, year string,
	info string, rating string, url string) {
	movieInfo := Movie{}
	movieInfo.Id = id
	movieInfo.Title = title
	movieInfo.Rating = rating
	movieInfo.Url = url
	movieInfo.Year = year
	err := db.Save(&movieInfo)
	if err != nil {
		fmt.Printf("save failed at %s, erro:%s", id, err.Error)
	}
	fmt.Println("save success", id)
}
