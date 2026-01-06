package models

type Client struct {
	ID  int    `json="id"`
	Nom string `json="nom"`
	Tel string `json="tel"`
}
