package dao

import (
	"fmt"
	"log"

	"../config"
	"../models"
	"gorm.io/gorm"
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

func ObtenerUsuarioPorID(userID string) (models.Usuario, error) {
	usuario := models.Usuario{}

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return usuario, fmt.Errorf("Error al conectarse a la base de datos. Error: %v", err)
	}

	if err := db.Debug().Where("id = ?", userID).First(&usuario).Error; err != nil {
		return usuario, fmt.Errorf("usuario no encontrado. Error: %v", err)
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

	res := db.Debug().Create(&usuario)
	if res.Error != nil {
		return usuario, fmt.Errorf("Error eliminando usuario. Error: %v", err)
	}

	return usuario, nil
}

func ModificarUsuario(userID string, nuevosAtributos map[string]interface{}) (models.Usuario, error) {
	var err error
	var res *gorm.DB
	var usuario models.Usuario

	db, err := config.BibliotecaDigitalDB()
	if err != nil {
		return usuario, err
	}

	usuario, err = ObtenerUsuarioPorID(userID)
	if err != nil {
		log.Println("Error obteniendo usuario", res.Error)
		return usuario, err
	}

	res = db.Model(&usuario).Updates(nuevosAtributos)
	if res.Error != nil {
		log.Println("Error al modificar usuario", res.Error)
		return usuario, res.Error
	}

	return usuario, err
}
