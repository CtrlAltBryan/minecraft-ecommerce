package pagos

import (
	"errors"
	"gorm.io/gorm"
	"moduloecommerce/productos"
	"time"
)

type Transaccion struct {
	gorm.Model
	UsuarioID uint
	Total     float64
	Fecha     time.Time
	Productos []productos.Producto `gorm:"many2many:transaccion_productos;"`
}

// RealizarCompra procesa la compra del carrito.
func RealizarCompra(db *gorm.DB, usuarioID uint) (*Transaccion, error) {
	var carrito []Carrito
	if err := db.Preload("Producto").Where("usuario_id = ?", usuarioID).Find(&carrito).Error; err != nil {
		return nil, err
	}

	if len(carrito) == 0 {
		return nil, errors.New("el carrito está vacío")
	}

	total := 0.0
	for _, item := range carrito {
		total += float64(item.Cantidad) * item.Producto.Precio
	}

	transaccion := Transaccion{
		UsuarioID: usuarioID,
		Total:     total,
		Fecha:     time.Now(),
	}

	for _, item := range carrito {
		transaccion.Productos = append(transaccion.Productos, item.Producto)
	}

	if err := db.Create(&transaccion).Error; err != nil {
		return nil, err
	}

	if err := db.Where("usuario_id = ?", usuarioID).Delete(&Carrito{}).Error; err != nil {
		return nil, err
	}

	return &transaccion, nil
}
