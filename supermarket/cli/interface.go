package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tarea1/golang-api/entity"
)

var (
	isLogged  = false
	isAdmin   = false
	sessionId = -1
)

type Access struct {
	ValidAccess bool `json:"acceso_valido"`
}

func DisplayInterface() {
	var option string
	fmt.Println("Bienvenido")

	for {
		if !isLogged {
			showMain()
		} else {
			if isAdmin {
				showAdmin()
			} else {
				showClient()
			}
		}

		fmt.Scan(&option)

		if !isLogged {
			if option == "1" {
				var (
					id     string
					passwd string
				)

				fmt.Print("Ingrese su id: ")
				fmt.Scan(&id)

				fmt.Print("Ingrese su contraseña: ")
				fmt.Scan(&passwd)

				parsedId, _ := strconv.Atoi(id)

				values, _ := json.Marshal(map[string]interface{}{
					"id_cliente": parsedId,
					"contrasena": passwd,
				})

				responseBody := bytes.NewBuffer(values)

				URL := "http://localhost:5000/api/clientes/iniciar_sesion"
				resp, _ := http.Post(URL, "application/json", responseBody)

				var cResp Access

				if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
					log.Fatal("Error")
				}

				defer resp.Body.Close()

				if cResp.ValidAccess {
					isLogged = true
					sessionId = parsedId

					fmt.Println("Inicio de sesión exitoso")
				} else {
					fmt.Println("Error, no hay ninguna coincidencia con los datos ingresados.")
				}
			} else if option == "2" {
				var passwd string
				fmt.Print("Ingrese su contraseña: ")
				fmt.Scan(&passwd)

				if passwd == "1234" {
					isLogged = true
					isAdmin = true

					fmt.Println("Inicio de sesión exitoso")
				} else {
					fmt.Println("Error, no hay ninguna coincidencia con los datos ingresados.")
				}
			} else {
				fmt.Println("")
				fmt.Println("Hasta luego!")

				os.Exit(0)
			}
		} else {
			if isAdmin {
				switch option {
				case "1":
					URL := "http://localhost:5000/api/productos"
					res, _ := http.Get(URL)

					var cResp []entity.Product

					if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
						log.Fatal("Error")
					}

					for i := range cResp {
						fmt.Printf("%d;%s;%d por unidad;%d disponibles\n", cResp[i].IdProduct, cResp[i].Name, cResp[i].Price, cResp[i].Available)
					}
				case "2":
					var (
						name      string
						available int
						price     int
					)

					fmt.Print("Ingrese el nombre: ")
					fmt.Scan(&name)

					fmt.Print("Ingrese la disponibilidad: ")
					fmt.Scan(&available)

					fmt.Print("Ingrese el precio unitario: ")
					fmt.Scan(&price)

					values, _ := json.Marshal(entity.Product{
						Available: available,
						Name:      name,
						Price:     price,
					})

					responseBody := bytes.NewBuffer(values)

					URL := "http://localhost:5000/api/productos"
					resp, _ := http.Post(URL, "application/json", responseBody)

					defer resp.Body.Close()

					fmt.Println("Producto ingresado correctamente!")
				case "3":
					var deleteId string

					fmt.Print("Ingrese la id del producto: ")
					fmt.Scan(&deleteId)

					URL := strings.Replace("http://localhost:5000/api/productos/?", "?", deleteId, 1)

					client := &http.Client{}
					req, _ := http.NewRequest("DELETE", URL, nil)

					_, _ = client.Do(req)
					fmt.Println("Producto eliminado correctamente!")
				case "4":
					URL := "http://localhost:5000/api/estadisticas"
					res, _ := http.Get(URL)

					var cResp entity.Statistics

					if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
						log.Fatal("Error")
					}

					fmt.Printf("Producto más vendido: %d\n", cResp.MostBought)
					fmt.Printf("Producto menos vendido: %d\n", cResp.LeastBought)
					fmt.Printf("Producto con mayor ganacia: %d\n", cResp.MostProfit)
					fmt.Printf("Producto con menor ganancia: %d\n", cResp.LeastProfit)
				case "5":
					isLogged = false
					isAdmin = false

					sessionId = -1
					fmt.Println("")
					fmt.Println("Hasta luego!")
				}
			} else {
				switch option {
				case "1":
					URL := "http://localhost:5000/api/productos"
					res, _ := http.Get(URL)

					var cResp []entity.Product

					if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
						log.Fatal("Error")
					}

					for i := range cResp {
						fmt.Printf("%d;%s;%d por unidad;%d disponibles\n", cResp[i].IdProduct, cResp[i].Name, cResp[i].Price, cResp[i].Available)
					}
				case "2":
					var qtyBuy int

					fmt.Print("Ingrese cantidad de productos: ")
					fmt.Scan(&qtyBuy)

					items := []entity.Single{}

					for i := 0; i < qtyBuy; i++ {
						var parityShop string

						fmt.Print("Ingrese producto par id-cantidad: ")
						fmt.Scan(&parityShop)

						splittedComp := strings.Split(parityShop, "-")

						id, _ := strconv.Atoi(splittedComp[0])
						qty, _ := strconv.Atoi(splittedComp[1])

						items = append(items, entity.Single{
							IdProduct: id,
							Quantity:  qty,
						})
					}

					var totalItems int = 0
					var totalCost int = 0

					for m := range items {
						URL := strings.Replace("http://localhost:5000/api/productos/?", "?", strconv.Itoa(items[m].IdProduct), 1)
						res, _ := http.Get(URL)

						var cResp entity.Product

						if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
							fmt.Printf("El producto #%d no se encontro en la base de datos.\n", items[m].IdProduct)
						}

						var actualStock int = items[m].Quantity

						if actualStock > cResp.Available {
							// Solicitar cantidad faltante al proveedor
							var cReq = cResp
							cReq.Available = actualStock - cResp.Available
							values, _ := json.Marshal(cReq)
							responseBody := bytes.NewBuffer(values)
							URLSupplier := "http://localhost:5000/api/proveedor"
							respReq, _ := http.Post(URLSupplier, "application/json", responseBody)

							respReq.Body.Close()
							cResp.Available = 0
						} else {
							cResp.Available = cResp.Available - actualStock
						}

						Json, _ := json.Marshal(cResp)
						clientRequest("PUT", URL, bytes.NewBuffer(Json))

						totalCost += actualStock * cResp.Price
						totalItems += actualStock
					}

					values, _ := json.Marshal(entity.Shopping{
						IdClient: sessionId,
						Products: items,
					})

					responseBody := bytes.NewBuffer(values)
					URL := "http://localhost:5000/api/compras"

					resp, _ := clientRequest("POST", URL, responseBody)
					defer resp.Body.Close()

					var cResp entity.Delivery

					if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
						log.Fatal("Error")
					}

					fmt.Println("Gracias por su compra!")
					fmt.Println("Cantidad de productos comprados: ", totalItems)
					fmt.Println("Monto total de la compra: ", totalCost)
					fmt.Println("El ID del despacho es: ", cResp.IdDespacho)
				case "3":
					var (
						idBuy int
						cResp entity.Delivery
					)

					fmt.Print("Ingrese el ID del despacho: ")
					fmt.Scan(&idBuy)

					URL := strings.Replace("http://10.10.11.204:5000/api/clientes/estado_despacho/?", "?", strconv.Itoa(idBuy), 1)
					res, _ := http.Get(URL)

					if err := json.NewDecoder(res.Body).Decode(&cResp); err != nil {
						log.Fatal("Error")
					}

					fmt.Printf("El estado del despacho es: %s\n", cResp.Estado)

				case "4":
					isLogged = false
					sessionId = -1

					fmt.Println("")
					fmt.Println("Hasta luego!")
				}
			}
		}
	}
}

func clientRequest(reqType string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(reqType, url, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)

	if err != nil {
		return res, err
	}

	return res, nil
}

func showMain() {
	fmt.Println("")

	fmt.Println("Opciones:")
	fmt.Println("1. Iniciar sesión como cliente")
	fmt.Println("2. Iniciar sesión como administrador")
	fmt.Println("3. Salir")

	fmt.Print("Ingrese una opción: ")
}

func showClient() {
	fmt.Println("")

	fmt.Println("Opciones:")
	fmt.Println("1. Ver lista de productos")
	fmt.Println("2. Hacer compra")
	fmt.Println("3. Consultar despacho")
	fmt.Println("4. Salir")

	fmt.Print("Ingrese una opción: ")
}

func showAdmin() {
	fmt.Println("")

	fmt.Println("Opciones:")
	fmt.Println("1. Ver lista de productos")
	fmt.Println("2. Crear producto")
	fmt.Println("3. Eliminar producto")
	fmt.Println("4. Ver estadísticas")
	fmt.Println("5. Salir")

	fmt.Print("Ingrese una opción: ")
}
