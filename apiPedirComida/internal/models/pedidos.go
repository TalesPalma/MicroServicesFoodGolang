package models

import (
	"gorm.io/gorm"
)

type PedidoStatus string

const (
	StatusPendendte   PedidoStatus = "pendente"
	StatusProcessando PedidoStatus = "processando"
	StatusAprovado    PedidoStatus = "aprovado"
	StatusReprovado   PedidoStatus = "reprovado"
)

type Pedidos struct {
	*gorm.Model
	Nome       string       `json:"nome"`
	Quantidade int          `json:"quantidade"`
	Preco      float64      `json:"preco"`
	Status     PedidoStatus `json:"status"`
}

func NovoPedido(dto PedidosDto) *Pedidos {
	return &Pedidos{
		Nome:       dto.Nome,
		Quantidade: dto.Quantidade,
		Preco:      dto.Preco,
		Status:     StatusPendendte,
	}
}

func (p *Pedidos) AtualizarStatus() {
	switch p.Status {
	case StatusPendendte:
		p.Status = StatusProcessando
	case StatusProcessando:
		p.Status = StatusAprovado
	case StatusAprovado:
		p.Status = StatusReprovado
	}
}
