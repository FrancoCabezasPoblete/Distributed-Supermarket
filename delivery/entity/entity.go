package entity

type Registro struct {
	IdDespacho int `json:"id_despacho"`

	Estado string `json:"estado"`

	IdCompra int `json:"id_compra"`
}
