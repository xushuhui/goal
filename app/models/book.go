package models

type Book struct {
	ID   uint   `sql:"id"`
	Name string `sql:"name"`
}
