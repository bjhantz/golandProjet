package repositories

import (
	"database/sql"

	"github.com/apiGO/internal/models"
)

func GetAllClient(db *sql.DB) ([]models.Client, error) {
	rows, err := db.Query("SELECT id, nom, tel, email FROM Client")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var client models.Client

		err := rows.Scan(&client.ID, &client.Nom, &client.Tel, &client.Email)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func GetOneClient(db *sql.DB, id int) (models.Client, error) {
	var client models.Client
	err := db.QueryRow("SELECT id, nom, tel, email FROM Client where id=$1", id).
		Scan(&client.ID, &client.Nom, &client.Tel, &client.Email)
	if err != nil {
		return client, err
	}

	return client, nil
}

func CreateClient(db *sql.DB, client models.Client) error {
	_, err := db.Exec("INSERT INTO Client (nom, tel, email) VALUES($1, $2)", client.Nom, client.Tel, client.Email)

	return err
}

func UpdateClient(db *sql.DB, client models.Client, id int) error {
	_, err := db.Exec("UPDATE Client SET nom=$1, tel=$2, email=$3 WHERE id=$4", client.Nom, client.Tel, client.Email, id)

	return err
}

func DeleteClient(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM Client Where id=$1", id)

	return err
}
