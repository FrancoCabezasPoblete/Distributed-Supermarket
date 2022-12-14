package storage

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"

	"tarea2/delivery/entity"
)

var DB *sql.DB

const (
	DB_USER   = "admin"
	DB_PASSWD = "12345678"
	DB_ADDR   = "localhost:3306"
	DB_NAME   = "db_despachos"
)

func ConnectDatabase() *sql.DB {

	cfg := mysql.Config{
		User:   DB_USER,
		Passwd: DB_PASSWD,
		Net:    "tcp",
		Addr:   DB_ADDR,
		DBName: DB_NAME,
	}

	var err error

	// connect to the database with admin
	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingError := DB.Ping()

	if pingError != nil {
		log.Fatal(pingError)
	}

	fmt.Println("Connection Succesful")

	return DB
}

func AddEstado(id int) *entity.Registro {

	registro := &entity.Registro{}

	// get the current state of the delivery with the id
	results, err := DB.Query("SELECT estado FROM despacho WHERE id_despacho=?", id)

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	if !results.Next() {
		return nil
	}

	err = results.Scan(&registro.Estado)
	if err != nil {
		return nil
	}

	return registro

}

func AddRegister(registro entity.Registro) int {

	fmt.Printf(strconv.Itoa(registro.IdDespacho))

	// add register to sql
	insert, err := DB.Exec(
		"INSERT INTO despacho (id_despacho, estado, id_compra) VALUES (?,?,?)",
		registro.IdDespacho, registro.Estado, registro.IdCompra)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := insert.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return int(lastId)

}
