package gateway

import "github.com.br/joaopedrolacerda/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
