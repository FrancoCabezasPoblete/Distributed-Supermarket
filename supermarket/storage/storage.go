package storage

import (
	"database/sql"
	"fmt"
	"log"

	"tarea1/golang-api/entity"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	DB_USER   = "admin"
	DB_PASSWD = "12345678"
	DB_ADDR   = "localhost:3306"
	DB_NAME   = "tarea_1_sd"
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
	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingError := DB.Ping()

	if pingError != nil {
		log.Fatal(pingError)
	}

	fmt.Println("Conexion a base de datos")

	return DB
}

func GetProducts() []entity.Product {
	results, err := DB.Query("SELECT * FROM producto")

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	products := []entity.Product{}

	for results.Next() {
		var product entity.Product
		err = results.Scan(&product.IdProduct, &product.Name, &product.Available, &product.Price)

		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	return products
}

func GetProduct(id int) *entity.Product {
	product := &entity.Product{}

	results, err := DB.Query("SELECT * FROM producto WHERE id_producto=?", id)

	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}

	if !results.Next() {
		return nil
	}

	err = results.Scan(&product.IdProduct, &product.Name, &product.Available, &product.Price)
	if err != nil {
		return nil
	}

	return product
}

func AddProduct(product entity.Product) int {
	insert, err := DB.Exec(
		"INSERT INTO producto (nombre, cantidad_disponible, precio_unitario) VALUES (?,?,?)",
		product.Name, product.Available, product.Price)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := insert.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	return int(lastId)
}

func DeleteProduct(id int) int {
	_, err := DB.Exec("DELETE FROM producto WHERE id_producto=?", id)

	if err != nil {
		panic(err.Error())
	}

	return id
}

func UpdateProduct(product *entity.Product) int {
	_, err := DB.Exec("UPDATE producto SET nombre=?, cantidad_disponible=?, precio_unitario=? WHERE id_producto=?",
		product.Name, product.Available, product.Price, product.IdProduct)

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return product.IdProduct
}

func ValidateSession(client entity.Client) bool {
	isValid := true
	results, err := DB.Query("SELECT * FROM cliente WHERE id_cliente=? AND contrasena=?", client.IdClient, client.Password)

	if err != nil || !results.Next() {
		isValid = false
	}

	return isValid
}

func AddShopping(shopping entity.Shopping) int {
	insShopping, _ := DB.Exec(
		"INSERT INTO compra (id_cliente) VALUES (?)", shopping.IdClient)

	lastId, _ := insShopping.LastInsertId()
	products := shopping.Products

	for i := range products {
		_, err := DB.Exec(
			"INSERT INTO detalle (id_compra, id_producto, cantidad) VALUES (?,?,?)",
			lastId, products[i].IdProduct, products[i].Quantity)

		if err != nil {
			panic(err.Error())
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	return int(lastId)
}

func GetStatistics() [4]int {
	statistics := [4]int{GetMostBought(), GetLeastBought(), GetMostProfit(), GetLeastProfit()}

	return statistics
}

func GetMostBought() int {
	results, err := DB.Query("SELECT id_producto FROM detalle GROUP BY id_producto ORDER BY SUM(cantidad) DESC LIMIT 1")

	var mostBought int

	if !results.Next() {
		return 0
	}

	err = results.Scan(&mostBought)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return int(mostBought)
}

func GetLeastBought() int {
	results, err := DB.Query("SELECT id_producto FROM detalle GROUP BY id_producto ORDER BY SUM(cantidad) ASC LIMIT 1")

	var leastBought int

	if !results.Next() {
		return 0
	}

	err = results.Scan(&leastBought)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return int(leastBought)
}

func GetMostProfit() int {
	results, err := DB.Query("SELECT detalle.id_producto FROM detalle, producto WHERE detalle.id_producto = producto.id_producto GROUP BY detalle.id_producto ORDER BY SUM(detalle.cantidad)*producto.precio_unitario DESC LIMIT 1")

	var mostProfit int

	if !results.Next() {
		return 0
	}

	err = results.Scan(&mostProfit)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return int(mostProfit)
}

func GetLeastProfit() int {
	results, err := DB.Query("SELECT detalle.id_producto FROM detalle, producto WHERE detalle.id_producto = producto.id_producto GROUP BY detalle.id_producto ORDER BY SUM(detalle.cantidad)*producto.precio_unitario ASC LIMIT 1")

	var leastProfit int

	if !results.Next() {
		return 0
	}

	err = results.Scan(&leastProfit)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	return int(leastProfit)
}
