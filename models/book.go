package models

import "gorm.io/gorm"

type Book struct {
    gorm.Model
	Id string `json:"id"`
    Title  string `json:"title"`
    User string `json:"author"`
}


var Books = []Book{{
    Id: "1",
    Title: "Book 1",
    User: "User 1",
},
{
    Id: "2",
    Title: "Book 2",
    User: "User 2",
},
}