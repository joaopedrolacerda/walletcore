package createclient

import (
	"time"

	"github.com.br/joaopedrolacerda/walletcore/internal/entity"
	"github.com.br/joaopedrolacerda/walletcore/internal/gateway"
)

type CreateClientInputDTO struct {
	name  string
	email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	clientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		clientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.name, input.email)
	if err != nil {
		return nil, err
	}
	err = uc.clientGateway.Save(client)

	if err != nil {
		return nil, err
	}
	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
