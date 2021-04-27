package services

import (
	"fmt"
	"time"

	"github.com/tesis/API-Usuario/src/api/dao"
	"github.com/tesis/API-Usuario/src/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AutenticarUsuario(email, password string) map[string]interface{} {
	var usuario models.Usuario
	var err error

	if usuario, err = dao.ObtenerUsuario(email); err != nil {
		var resp = map[string]interface{}{"status": false, "message": err.Error()}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 3600).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Contraseña incorrecta. Intente nuevamente"}
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

	tokenString, error := token.SignedString([]byte(uuid.NewString()))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": true, "message": "Usuario autenticado"}

	resp["token"] = tokenString //Guarda el token en la respuesta
	resp["usuario"] = usuario
	return resp
}

func ObtenerUsuarios() ([]models.Usuario, error) {
	return dao.ObtenerUsuarios()
}

func EliminarUsuario(id string) error {
	return dao.BajaUsuario(id)
}

func CrearUsuario(usuario models.Usuario) (models.Usuario, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		return usuario, err
	}

	usuario.Password = string(pass)

	if usuario, err = dao.AltaUsuario(usuario); err != nil {
		return usuario, err
	}

	return usuario, err
}
