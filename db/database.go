package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:password@tcp(localhost:3306)/tb2_go"

// Guarda la conexion
var db *sql.DB

// Realiza la conexion
func Connect() {
	con, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("conexion exitosa")
	db = con

}

// Cerrar la conexion
func Close() {
	db.Close()
}

// Verificar la conexion
func Ping() {
	Connect()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	Close()
}

// Verifica si una tabla existe
func ExistsTable(tableName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(query)

	if err != nil {
		fmt.Println("Error buscando tabla", tableName, "en la base de datos:", err)
	}
	return rows.Next()

}

// Crear tabla
func CreateTable(schema, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println("Error creando tabla", name, "en la base de datos:", err)
		}
	}
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}

//Reniciar el registro de una tabla

func TruncateTable(tableName string) {
	query := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(query)
}
