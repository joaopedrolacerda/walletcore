package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("john Doe", "j@j")
	Account1 := NewAccount(client1)

	client2, _ := NewClient("john doe 2", "j@j2")
	Account2 := NewAccount(client2)

	Account1.Credit(1000)
	Account2.Credit(1000)

	Transaction, err := newTransaction(Account1, Account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, Transaction)
	assert.Equal(t, 1100.0, Account2.Balance)
	assert.Equal(t, 900.0, Account1.Balance)
}

func TestClientTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("john Doe", "j@j")
	Account1 := NewAccount(client1)

	client2, _ := NewClient("john doe 2", "j@j2")
	Account2 := NewAccount(client2)

	Account1.Credit(1000)
	Account2.Credit(1000)

	Transaction, err := newTransaction(Account1, Account2, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insuficient funds")
	assert.Nil(t, Transaction)
	assert.Equal(t, 1000.0, Account2.Balance)
	assert.Equal(t, 1000.0, Account1.Balance)
}
