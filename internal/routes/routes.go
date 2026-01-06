package routes

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/apiGO/internal/handlers"
	
)

func RegisterRoutes(db *sql.DB) http.Handler{
	// http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		handlers.GetEmployees(db)(w, r)
	// 	}

	// 	if r.Method == http.MethodPost {
	// 		handlers.CreateEmploye(db)(w, r)
	// 	}
	// })

	r:= chi.NewRouter()

	r.Get("/employees", handlers.GetEmployees(db))
	r.Put("/employees/{id}", handlers.UpdateEmployee(db))
	r.Post("/employees", handlers.CreateEmploye(db))
	r.Get("/employees/{id}", handlers.GetOneEmployee(db))
	r.Delete("/employees/{id}", handlers.DeleteEmployee(db))
	
	return r
}
