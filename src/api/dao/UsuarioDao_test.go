package dao

import (
	"os"
	"testing"

	"github.com/tesis/API-Usuario/src/api/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("TEST", "1")
}

func TestObtenerUsuario(t *testing.T) {
	assert := assert.New(t)

	usuario := models.Usuario{
		Nombre:        "Juan",
		Apellido:      "Fuentes",
		Email:         "juancito@gmail.com",
		Password:      "juancito",
		Username:      "jfuentes",
		Rol:           "DIRECTOR",
		Imagen_url:    "imagen.jpg",
		Telefono:      "2664258546",
		Especialidad:  "-",
		InstitucionID: 1,
	}

	usuario, err := AltaUsuario(usuario)
	assert.Nil(err)
	assert.NotNil(usuario.ID)
	assert.Equal("Juan", usuario.Nombre)
	assert.Equal("Fuentes", usuario.Apellido)

	usuario, err = ObtenerUsuario("juancito@gmail.com")
	assert.Nil(err)
	assert.NotNil(usuario.ID)
	assert.Equal("Juan", usuario.Nombre)
	assert.Equal("Fuentes", usuario.Apellido)
}
