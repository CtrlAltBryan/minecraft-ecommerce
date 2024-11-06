// productos/gestor.go
package productos

import (
	"log"
	"gorm.io/gorm"
)

type Producto struct {
	ID          uint    `gorm:"primaryKey"`
	Nombre      string  `gorm:"unique"`
	Descripcion string
	Version     string
	Categoria   string
	Precio      float64
}

// ObtenerProductoPorNombre busca un producto por su nombre.
func ObtenerProductoPorNombre(db *gorm.DB, nombre string) (Producto, error) {
	var producto Producto
	if err := db.Where("nombre = ?", nombre).First(&producto).Error; err != nil {
		return producto, err
	}
	return producto, nil
}

// ProbarGestorProductos: Función para probar la gestión de productos
func ProbarGestorProductos(db *gorm.DB) {
	producto, err := ObtenerProductoPorNombre(db, "Super Plugin")
	if err != nil {
		log.Println("Error al obtener producto:", err)
	} else {
		log.Println("Producto encontrado:", producto.Nombre)
	}
}
