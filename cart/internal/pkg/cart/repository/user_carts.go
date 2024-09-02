package repository

import "gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"

type UserStorage = map[model.UID][]Cart
type UserCart struct {
	storage UserStorage
}
