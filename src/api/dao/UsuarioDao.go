package dao

import (
	"fmt"

	"github.com/tesis/API-Usuario/src/api/config"
	"github.com/tesis/API-Usuario/src/api/models"
)

func ObtenerUsuario(email string) (models.Usuario, error) {
	usuario := models.Usuario{}

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return usuario, fmt.Errorf("Error al conectarse a la base de datos. Error: %v", err)
	}

	if err := db.Debug().Where("email = ?", email).First(&usuario).Error; err != nil {
		return usuario, fmt.Errorf("Email no encontrado. Error: %v", err)
	}

	return usuario, nil
}

func ObtenerUsuarios() ([]models.Usuario, error) {
	usuarios := []models.Usuario{}

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return usuarios, fmt.Errorf("Error al conectarse a la base de datos. Error: %v", err)
	}

	if err := db.Find(&usuarios).Error; err != nil {
		return usuarios, fmt.Errorf("Error al buscar usuarios. Error: %v", err)
	}

	return usuarios, nil
}

func BajaUsuario(id string) error {
	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return fmt.Errorf("Error al conectarse a la base de datos. Error: %v", err)
	}

	res := db.Unscoped().Where("id = ?", id).Delete(models.Usuario{})
	if res.Error != nil {
		return fmt.Errorf("Error eliminando usuario. Error: %v", err)
	}

	return nil
}

func AltaUsuario(usuario models.Usuario) (models.Usuario, error) {
	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return usuario, fmt.Errorf("Error al conectarse a la base de datos. Error: %v", err)
	}

	res := db.Create(&usuario)
	if res.Error != nil {
		return usuario, fmt.Errorf("Error eliminando usuario. Error: %v", err)
	}

	return usuario, nil
}
