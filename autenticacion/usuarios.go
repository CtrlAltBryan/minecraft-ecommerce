// autenticacion/usuarios.go
package autenticacion

import (
	"log"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuario struct {
	ID        uint   `gorm:"primaryKey"`
	Nombre    string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Contraseña string
	Rol       string
}

// AutenticarUsuario busca un usuario por su email y valida la contraseña.
func AutenticarUsuario(db *gorm.DB, email, contraseña string) (Usuario, error) {
	var usuario Usuario
	if err := db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return usuario, errors.New("usuario no encontrado")
	}

	// Comparar contraseñas
	err := bcrypt.CompareHashAndPassword([]byte(usuario.Contraseña), []byte(contraseña))
	if err != nil {
		return usuario, errors.New("contraseña incorrecta")
	}

	return usuario, nil
}

// ProbarAutenticacion: Función para probar la autenticación de un usuario
func ProbarAutenticacion(db *gorm.DB) {
	usuario, err := AutenticarUsuario(db, "bryangabito@hotmail.com", "1234")
	if err != nil {
		log.Println("Error de autenticación:", err)
	} else {
		log.Println("Usuario autenticado:", usuario.Nombre)
	}
}
