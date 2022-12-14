package entity

type Product struct {
	IdProduct int    `json:"id_producto,omitempty"`
	Name      string `json:"nombre"`
	Available int    `json:"cantidad_disponible"`
	Price     int    `json:"precio_unitario"`
}
