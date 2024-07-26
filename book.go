package main

type Book struct {
    Id    string `json:"id"`
    Title string `json:"title"`
    User  string `json:"user"`
}

var Books = []Book{
    {Id: "1", Title: "Book 1", User: "User 1"},
    {Id: "2", Title: "Book 2", User: "User 2"},
}
