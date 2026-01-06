package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/apiGO/internal/models"
	"github.com/apiGO/internal/repositories"
)

func GetEmployees(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := repositories.GetAllEmployees(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}

func CreateEmploye(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var emp models.Employee

		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		err = repositories.CreateEmploye(db, emp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetOneEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := chi.URLParam(r,"id") //On recupere l'id de l'URL
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		emp, err := repositories.GetEmployeeByID(db, id)
		if err != nil {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(emp)
	}
}

func UpdateEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var emp models.Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := repositories.UpdateEmployee(db, id, emp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteEmployee(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := repositories.DeleteEmployee(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent) // 
	}
}