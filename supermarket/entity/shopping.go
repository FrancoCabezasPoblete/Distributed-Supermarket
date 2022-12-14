package entity

type Shopping struct {
	IdClient int      `json:"id_cliente"`
	Products []Single `json:"productos"`
}
