package entity_test

import (
	"github.com/GeovaneCavalcante/api-anime-golang/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	ID := entity.NewID()
	assert.NotNil(t, ID)
	assert.ObjectsAreEqual(entity.ID{}, ID)
}

func TestStringToID(t *testing.T) {
	i := entity.NewID().String()
	ID, err := entity.StringToID(i)
	assert.Nil(t, err)
	assert.NotNil(t, ID)
	assert.ObjectsAreEqual(entity.ID{}, ID)

}
