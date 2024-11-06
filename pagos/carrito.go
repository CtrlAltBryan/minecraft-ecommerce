package pagos

import (
	"gorm.io/gorm"
	"moduloecommerce/productos"  // Importa el paquete productos
)

type Carrito struct {
	ID         uint              `gorm:"primaryKey"`
	UsuarioID  uint              `gorm:"index"`
	ProductoID uint              `gorm:"index"`
	Cantidad   int
	Producto   productos.Producto `gorm:"foreignKey:ProductoID"`  // Ahora reconoce productos.Producto
}

// AgregarAlCarrito agrega un producto al carrito de un usuario.
func AgregarAlCarrito(db *gorm.DB, usuarioID, productoID, cantidad int) error {
	var carrito Carrito
	if err := db.Where("usuario_id = ? AND producto_id = ?", usuarioID, productoID).First(&carrito).Error; err == nil {
		carrito.Cantidad += cantidad
		if err := db.Save(&carrito).Error; err != nil {
			return err
		}
	} else {
		newCarrito := Carrito{
			UsuarioID:  uint(usuarioID),
			ProductoID: uint(productoID),
			Cantidad:   cantidad,
		}
		if err := db.Create(&newCarrito).Error; err != nil {
			return err
		}
	}
	return nil
}

// ListarCarrito devuelve todos los productos en el carrito de un usuario.
func ListarCarrito(db *gorm.DB, usuarioID uint) ([]Carrito, error) {
	var carritos []Carrito
	if err := db.Preload("Producto").Where("usuario_id = ?", usuarioID).Find(&carritos).Error; err != nil {
		return nil, err
	}
	return carritos, nil
}
