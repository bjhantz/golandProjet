package repositories

import (
	"database/sql"

	"github.com/apiGO/internal/models"
)

func GetAllEmployees(db *sql.DB) ([]models.Employee, error) {
	rows, err := db.Query("SELECT id, name, position, salary FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee

		err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

func CreateEmploye(db *sql.DB, emp models.Employee) error {
	_, err := db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)",
		emp.Name, emp.Position, emp.Salary)

	return err
}

func GetEmployeeByID(db *sql.DB, id int) (models.Employee, error) {
	var emp models.Employee
	err := db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id=$1", id).
		Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func UpdateEmployee(db *sql.DB, id int, emp models.Employee) error {
	_, err := db.Exec(
		"UPDATE employees SET name=$1, position=$2, salary=$3 WHERE id=$4",
		emp.Name, emp.Position, emp.Salary, id,
	)
	return err
}

func DeleteEmployee(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM employees WHERE id=$1", id)
	return err
}

