package usecases

import "pathipat/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
