package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	prod := Product{
		Price: 1.22,
	}

	v := NewValidation()
	err := v.Validate(prod)
	assert.Len(t, err, 2)
}

func TestProductMissingPriceReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: -1,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 2)
}

// TODO: Add more tests here...
