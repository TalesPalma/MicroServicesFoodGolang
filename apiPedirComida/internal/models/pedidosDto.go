package models

type PedidosDto struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
	Status     string  `json:"status"`
}

func PedidosListToDto(p []Pedidos) []PedidosDto {
	var pedidosDto []PedidosDto = []PedidosDto{}
	for _, v := range p {
		pedidos := PedidosDto{
			Id:         int(v.ID),
			Nome:       v.Nome,
			Quantidade: v.Quantidade,
			Preco:      v.Preco,
			Status:     string(v.Status),
		}
		pedidosDto = append(pedidosDto, pedidos)
	}
	return pedidosDto
}
