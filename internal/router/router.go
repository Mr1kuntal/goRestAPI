package router

import (
	"net/http"
	"crudAPI/internal/handler"
)


func Router(h *handler.Handler) {
	http.HandleFunc("GET /users",h.GetUsers)
	http.HandleFunc("GET /users/{id}",h.GetUserbyid)
	http.HandleFunc("POST /users",h.CreateUsers)
	http.HandleFunc("PATCH /users/{id}", h.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", h.DeleteUser)
}



