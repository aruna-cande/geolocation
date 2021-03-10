package geolocation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newFixtureGeolocation() *geolocation {
	return &geolocation{
		IpAddress: "200.106.141.15",
		CountryCode:     "si",
		Country:    "SI",
		City:     "DuBuquemouth",
		Latitude:  -84.87503094689836,
		Longitude: 7.206435933364332,
		MysteryValue: 823011346,
	}
}

func Test_ImportGeolocationData(t *testing.T) {
	repo := newInmemRepository()
	m := NewService(repo)
	//u := newFixtureBook()
	_, err := m.CreateBook(u.Title, u.Author, u.Pages, u.Quantity)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}