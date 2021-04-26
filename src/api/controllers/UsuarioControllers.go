package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

type error interface {
	Error() string
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := models.Usuario{}
	json.NewDecoder(r.Body).Decode(&usuario)

	usuario, err := services.CrearUsuario(usuario)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(usuario)
}

func AutenticarUsuario(w http.ResponseWriter, r *http.Request) {
	usuario := &models.Usuario{}
	err := json.NewDecoder(r.Body).Decode(usuario)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := services.AutenticarUsuario(usuario.Email, usuario.Password)
	json.NewEncoder(w).Encode(resp)
}

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, err := services.ObtenerUsuarios()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(usuarios)
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := services.EliminarUsuario(id); err != nil {
		var resp = map[string]interface{}{"status": false, "message": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var resp = map[string]interface{}{"status": true, "message": "Usuario eliminado satisfactoriamente"}
	json.NewEncoder(w).Encode(resp)
}
