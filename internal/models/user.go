package models

type User struct {
	ID       int    `json="id"`
	Nom      string `json="nom"`
	Prenom   string `json="prenom"`
	Passowrd string `json="password"`
	Email    string `json="email"`
}
