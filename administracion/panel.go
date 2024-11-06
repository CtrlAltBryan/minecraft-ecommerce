// moduloecommerce/administracion/panel.go
package administracion

import (
    "gorm.io/gorm"
    "moduloecommerce/autenticacion"
    "moduloecommerce/productos"
)

func ListarUsuarios(db *gorm.DB) ([]autenticacion.Usuario, error) {
    var usuarios []autenticacion.Usuario
    if err := db.Find(&usuarios).Error; err != nil {
        return nil, err
    }
    return usuarios, nil
}

func ActualizarUsuario(db *gorm.DB, usuarioID uint, nuevosDatos autenticacion.Usuario) error {
    return db.Model(&autenticacion.Usuario{}).Where("id = ?", usuarioID).Updates(nuevosDatos).Error
}

func EliminarUsuario(db *gorm.DB, usuarioID uint) error {
    return db.Delete(&autenticacion.Usuario{}, usuarioID).Error
}

func ListarProductos(db *gorm.DB) ([]productos.Producto, error) {
    var productosList []productos.Producto
    if err := db.Find(&productosList).Error; err != nil {
        return nil, err
    }
    return productosList, nil
}

func ActualizarProducto(db *gorm.DB, productoID uint, nuevosDatos productos.Producto) error {
    return db.Model(&productos.Producto{}).Where("id = ?", productoID).Updates(nuevosDatos).Error
}

func EliminarProducto(db *gorm.DB, productoID uint) error {
    return db.Delete(&productos.Producto{}, productoID).Error
}
