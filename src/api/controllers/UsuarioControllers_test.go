package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"../models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCrearUsuario(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios/registrar", CrearUsuario).Methods("POST")

	assert := assert.New(t)
	bodyReq := models.Usuario{
		Email:         "juancito@gmail.com",
		Password:      "Juancito1234",
		Nombre:        "Juan",
		Apellido:      "Fuentes",
		Username:      "jfuentes",
		Especialidad:  "ninguna",
		InstitucionID: 1,
		Imagen_url:    "-",
		Rol:           "DIRECTOR",
		Telefono:      "2553543423",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bodyReq)

	req, err := http.NewRequest("POST", "/usuarios/registrar", b)
	if err != nil {
		t.Fatal("Creating POST, /usuarios/registrar request failed!")
	}

	respRec := httptest.NewRecorder()
	router.ServeHTTP(respRec, req)

	var resp models.Usuario
	json.Unmarshal(respRec.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, respRec.Code)
	assert.Equal(bodyReq.Email, resp.Email)
	assert.NotNil(resp.ID)
}

func TestAutenticarUsuario(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios/autenticar", AutenticarUsuario).Methods("POST")

	assert := assert.New(t)
	bodyReq := map[string]string{
		"email":    "juancito@gmail.com",
		"password": "Juancito1234",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bodyReq)

	req, err := http.NewRequest("POST", "/usuarios/autenticar", b)
	if err != nil {
		t.Fatal("Creating POST, /usuarios/autenticar request failed!")
	}

	respRec := httptest.NewRecorder()
	router.ServeHTTP(respRec, req)

	var resp map[string]interface{}
	json.Unmarshal(respRec.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, respRec.Code)
	assert.Equal("Usuario autenticado", resp["message"])
	assert.True(resp["status"].(bool))

}

func TestAutenticarUsuarioContraseñaIncorrecta(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios/autenticar", AutenticarUsuario).Methods("POST")

	assert := assert.New(t)

	bodyReq := map[string]string{
		"email":    "juancito@gmail.com",
		"password": "otracontraseña",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bodyReq)

	req, err := http.NewRequest("POST", "/usuarios/autenticar", b)
	if err != nil {
		t.Fatal("Creating POST, /usuarios/autenticar request failed!")
	}

	respRec := httptest.NewRecorder()
	router.ServeHTTP(respRec, req)

	var resp map[string]interface{}
	json.Unmarshal(respRec.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, respRec.Code)
	assert.Equal("Contraseña incorrecta. Intente nuevamente", resp["message"])
	assert.False(resp["status"].(bool))
}

func TestAutenticarUsuarioUsuarioIncorrecto(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios/autenticar", AutenticarUsuario).Methods("POST")

	assert := assert.New(t)

	bodyReq := map[string]string{
		"email":    "unmailinexistente@gmail.com",
		"password": "otracontraseña",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bodyReq)

	req, err := http.NewRequest("POST", "/usuarios/autenticar", b)
	if err != nil {
		t.Fatal("Creating POST, /usuarios/autenticar request failed!")
	}

	respRec := httptest.NewRecorder()
	router.ServeHTTP(respRec, req)

	var resp map[string]interface{}
	_ = json.Unmarshal(respRec.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, respRec.Code)
	assert.Equal("Email no encontrado.", resp["message"])
	assert.False(resp["status"].(bool))
}
