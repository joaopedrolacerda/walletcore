package createtransaction

import (
	"github.com.br/joaopedrolacerda/walletcore/internal/entity"
	"github.com.br/joaopedrolacerda/walletcore/internal/gateway"
)

type CreatetransactionInputDTO struct {
	AccountIDfrom string
	AccountIDTo   string
	Amount        float64
}

type CreatetransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreatetransactionInputDTO) (*CreatetransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindById(input.AccountIDfrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.AccountGateway.FindById(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreatetransactionOutputDTO{ID: transaction.ID}, nil
}
