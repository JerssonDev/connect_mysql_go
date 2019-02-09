package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // _ significa de manera explicitamente (solo se ejecutara el init)
)

// Persona ...
type Persona struct {
	gorm.Model // ID, CreatedAt, UpdatedAt , DeletedAt
	Nombre     string
	Edad       uint8
	Telefono   uint
}

func main() {

	db, err := gorm.Open("mysql", "user:password@/database?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("Error al conectarse a la base de datos :( " + err.Error())
	}

	defer db.Close()

	fmt.Println("Se conectó a la base de datos :) ")

	db.CreateTable(&Persona{})

	fmt.Println("Tabla creada correctamente :) ")

	p := &Persona{}

	p.Nombre = "Juanito"
	p.Edad = 20
	p.Telefono = 949651237

	db.Create(p)

	fmt.Println("Se insertó una persona :) ")

	var u Persona

	db.First(&u, "ID = ?", "2")

	fmt.Println("Se listará una persona :) ")
	fmt.Println(u)

	db.Model(&u).Update("edad", 20)

	db.First(&u, "ID = ?", "2")

	fmt.Printf("Edad de %s ah sido actualizada = %d :)\n", u.Nombre, u.Edad)

}
