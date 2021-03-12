package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeolocation(t *testing.T) {
	g := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	g1 := NewGeolocation("10.0.0.", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	g2 := NewGeolocation("10.0.0.1", "", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	g3 := NewGeolocation("10.0.0.1", "SI", "", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	g4 := NewGeolocation("10.0.0.1", "SI", "Nepal", "", "-84.87503094689836", "7.206435933364332", "7823011346")
	g5 := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-91.87503094689836", "7.206435933364332", "7823011346")
	g6 := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "", "7.206435933364332", "7823011346")
	g7 := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "181.206435933364332", "7823011346")
	g8 := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "", "7823011346")
	g9 := NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "")

	assert.NotNil(t, g)
	assert.Nil(t, g1)
	assert.Nil(t, g2)
	assert.Nil(t, g3)
	assert.Nil(t, g4)
	assert.Nil(t, g5)
	assert.Nil(t, g6)
	assert.Nil(t, g7)
	assert.Nil(t, g8)
	assert.Nil(t, g9)
}
