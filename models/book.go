package models

import "gorm.io/gorm"

type Book struct {
    gorm.Model
	Id int `json:"id"`
    Title  string `json:"title"`
    User string `json:"author"`
}
