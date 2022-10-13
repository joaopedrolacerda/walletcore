package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("john doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "john doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}
func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("john doe", "j@j.com")
	err := client.Update("John Doe Update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}
func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("john doe", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.Error(t, err, "name is required")
}
