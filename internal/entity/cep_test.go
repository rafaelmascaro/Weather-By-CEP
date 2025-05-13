package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipcode(t *testing.T) {
	_, err1 := NewCEP("13098401")
	_, err2 := NewCEP("13098-401")
	_, err3 := NewCEP("")

	assert.NoError(t, err1)
	assert.Error(t, err2, "invalid zipcode")
	assert.Error(t, err3, "invalid zipcode")
}
