package models

type Usuario struct {
	ID            uint   `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Rol           string `json:"rol"`
	Imagen_url    string `json:"imagen_url"`
	Telefono      string `json:"telefono"`
	Especialidad  string `json:"especialidad"`
	InstitucionID int64  `json:"institucion_id,omitempty"`
}

func (Usuario) TableName() string {
	return "usuarios"
}
