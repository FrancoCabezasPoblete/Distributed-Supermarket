package entity

type Statistics struct {
	MostBought  int `json:"producto_mas_vendido"`
	LeastBought int `json:"producto_menos_vendido"`
	MostProfit  int `json:"producto_mayor_ganancia"`
	LeastProfit int `json:"producto_menos_ganacia"`
}
