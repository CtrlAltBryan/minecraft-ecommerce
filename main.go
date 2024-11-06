package main

import (
	"fmt"
	"log"
	"moduloecommerce/autenticacion"
	"moduloecommerce/pagos"
	"moduloecommerce/productos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos exitosa")

	// Migraciones
	db.AutoMigrate(&autenticacion.Usuario{}, &productos.Producto{}, &pagos.Carrito{})

	// Crear un usuario si no existe
	usuario := autenticacion.Usuario{
		Nombre:     "Bryan",
		Email:      "bryangabito@hotmail.com",
		Contraseña: "1234",
		Rol:        "cliente",
	}
	if err := db.Where("nombre = ?", usuario.Nombre).FirstOrCreate(&usuario).Error; err != nil {
		log.Fatalf("Error al crear usuario: %v", err)
	}

	// Crear un producto si no existe
	producto := productos.Producto{
		Nombre:      "Super Plugin",
		Descripcion: "Plugin avanzado para Minecraft",
		Version:     "1.16.5",
		Categoria:   "Plugin",
		Precio:      9.99,
	}
	if err := db.Where("nombre = ?", producto.Nombre).FirstOrCreate(&producto).Error; err != nil {
		log.Fatalf("Error al crear producto: %v", err)
	}

	// Probar autenticación y gestión de productos
	autenticacion.ProbarAutenticacion(db)
	productos.ProbarGestorProductos(db)

	// Agregar producto al carrito (conversión de uint a int)
	if err := pagos.AgregarAlCarrito(db, int(usuario.ID), int(producto.ID), 2); err != nil {
		log.Printf("Error al agregar producto al carrito: %v", err)
	}

	// Listar contenido del carrito
	carrito, err := pagos.ListarCarrito(db, usuario.ID)
	if err != nil {
		log.Printf("Error al listar carrito: %v", err)
	}

	// Mostrar productos en el carrito
	fmt.Println("Productos en el carrito:")
	for _, item := range carrito {
		fmt.Printf("- %s (Cantidad: %d)\n", item.Producto.Nombre, item.Cantidad)
	}

	// Realizar la compra (capturar ambos valores devueltos)
	transaccion, err := pagos.RealizarCompra(db, usuario.ID)
	if err != nil {
		log.Printf("Error al realizar la compra: %v", err)
	} else {
		fmt.Println("Compra realizada con éxito. Total:", transaccion.Total)
	}
}
