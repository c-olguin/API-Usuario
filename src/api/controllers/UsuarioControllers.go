package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"../config"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func CrearUsuario(w http.ResponseWriter, r *http.Request) {
	var err error

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return
	}

	usuario := &models.Usuario{}
	json.NewDecoder(r.Body).Decode(usuario)

	pass, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Error en la encriptacion de la contraseña",
		}
		json.NewEncoder(w).Encode(err)
	}

	usuario.Password = string(pass)

	usuarioCreado := db.Create(usuario)
	var errMessage = usuarioCreado.Error

	if usuarioCreado.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(usuarioCreado)
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscando usuario")
	usuario := &models.Usuario{}
	err := json.NewDecoder(r.Body).Decode(usuario)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := findOne(usuario.Email, usuario.Password)
	json.NewEncoder(w).Encode(resp)
}

func findOne(email, password string) map[string]interface{} {
	usuario := &models.Usuario{}

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Error al conectarse a la base de datos"}
		return resp
	}

	if err := db.Where("Email = ?", email).First(usuario).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: usuario.ID,
		Name:   usuario.Nombre,
		Email:  usuario.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	fmt.Println("Usuario Encontrado")
	resp["token"] = tokenString //Guarda el token en la respuesta
	resp["usuario"] = usuario
	return resp
}

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	//db.Preload("auths").Find(&usuarios) puede haber ususarios que no esten registrados?
	db.Find(&usuarios)

	json.NewEncoder(w).Encode(usuarios)
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res := db.Unscoped().Where("id = ?", id).Delete(models.Usuario{})
	if res.Error != nil {
		log.Println("Error al eliminar usuario", res.Error)
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
	}

	var resp = map[string]interface{}{"status": true, "message": "Usuario eliminado satisfactoriamente"}
	json.NewEncoder(w).Encode(resp)
}
