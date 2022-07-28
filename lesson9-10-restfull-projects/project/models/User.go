package models

type User struct{
	ID int
	Name string
	FirstName string
	LastName string
	Interests []Interest
}