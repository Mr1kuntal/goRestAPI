package handler

import (
	"crudAPI/internal/database"
	"crudAPI/internal/dto"
	"crudAPI/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	db *database.Usersdbstr
}

func NewHandler(data *database.Usersdbstr) *Handler {
	return &Handler{
		db : data,
	}
}


func (h *Handler) GetUsers(w http.ResponseWriter,r *http.Request) {

	snapshot := make([]model.User,0,len(h.db.Usersdb))

	h.db.Mu.RLock()
	for _,val:= range h.db.Usersdb {
		snapshot = append(snapshot, val)
	}
	h.db.Mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(snapshot); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetUserbyid(w http.ResponseWriter,r *http.Request) {

	idstr := r.PathValue("id")
	id,err:= strconv.Atoi(idstr)
	if err != nil{
		http.Error(w,"invalid user id",http.StatusBadRequest)
		return
	}

	h.db.Mu.RLock()
	tempuser, ok := h.db.Usersdb[id]
	h.db.Mu.RUnlock()
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tempuser); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) CreateUsers(w http.ResponseWriter,r *http.Request){
	var req dto.CreateUser

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w,"invalid json data",http.StatusBadRequest)
		return
	}
	
	if strings.TrimSpace(req.Username) == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	if req.Age <= 0 {
		http.Error(w, "age must be greater than zero", http.StatusBadRequest)
		return
	}
	if len(req.Hobbies)<1{
		http.Error(w, "hobbies required", http.StatusBadRequest)
		return
	}

	h.db.Mu.Lock()
	id := h.db.NextID
	h.db.NextID++
	tempuser := model.User{
		ID:       id,
		Username: req.Username,
		Age:      req.Age,
		Hobbies:  req.Hobbies,
	}
	h.db.Usersdb[id] = tempuser
	h.db.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location",fmt.Sprintf("/users/%d",id))
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tempuser); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) UpdateUser(w http.ResponseWriter,r *http.Request){
	idstr := r.PathValue("id")
	id,err := strconv.Atoi(idstr);
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	
	var req dto.UpdateUserRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req);err!=nil{
		http.Error(w,"invalid json",http.StatusBadRequest)
		return
	}

	if req.Username!=nil {
		if strings.TrimSpace(*req.Username) == "" {
			http.Error(w, "username is required", http.StatusBadRequest)
			return
		}
	}
	if req.Age!=nil{
		if *req.Age <= 0 {
			http.Error(w, "age must be greater than zero", http.StatusBadRequest)
			return
		}
	
	}
	if req.Hobbies!=nil{
		if len(*req.Hobbies)<1{
			http.Error(w, "hobbies required", http.StatusBadRequest)
			return
		}
	}

	h.db.Mu.Lock()
	tempuser,ok:= h.db.Usersdb[id]
	if !ok {
		h.db.Mu.Unlock()
		http.Error(w,"user not found",http.StatusNotFound)
		return
	}
	if req.Username!= nil{tempuser.Username = *req.Username} 
	if req.Age!=nil{tempuser.Age=*req.Age}
	if req.Hobbies!=nil{tempuser.Hobbies=*req.Hobbies}
	h.db.Usersdb[id]=tempuser
	h.db.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tempuser); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) DeleteUser(w http.ResponseWriter,r *http.Request){
	idstr := r.PathValue("id")
	id,err := strconv.Atoi(idstr);
	if err!=nil{
		http.Error(w,"invalid id",http.StatusBadRequest)
		return
	}
	h.db.Mu.Lock()
	_ , ok := h.db.Usersdb[id]
	if !ok {
		h.db.Mu.Unlock()
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	delete(h.db.Usersdb,id)
	h.db.Mu.Unlock()
	
	w.WriteHeader(http.StatusNoContent)
	
}