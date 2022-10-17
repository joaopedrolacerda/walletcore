package gateway

import "github.com.br/joaopedrolacerda/walletcore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
