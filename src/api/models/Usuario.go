package models

import (
	"github.com/jinzhu/gorm"
) 

type Usuario struct {
	gorm.Model
	ID					 uint				 `json:"id"`
	Nombre               string               `json:"nombre"`
	Apellido             string               `json:"apellido"`
	Email            	 string               `json:"email"`
	Username             string               `json:"username"`
	Password             string               `json:"password"`
	Rol             	 string               `json:"rol"`
	Imagen_url           string               `json:"imagen_url"`
	Telefono             string               `json:"telefono"`
	Especialidad         string               `json:"especialidad"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

