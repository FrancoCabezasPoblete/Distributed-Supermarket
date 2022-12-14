package entity

type Client struct {
	IdClient int    `json:"id_cliente"`
	Name     string `json:"nombre,omitempty"`
	Password string `json:"contrasena"`
}
