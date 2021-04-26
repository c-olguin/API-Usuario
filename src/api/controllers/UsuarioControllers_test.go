package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

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
	json.Unmarshal(respRec.Body.Bytes(), &resp)

	assert.Equal(http.StatusOK, respRec.Code)
	assert.Equal("Email no encontrado. Error: record not found", resp["message"])
	assert.False(resp["status"].(bool))
}
