package geolocation

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/golang/mock/gomock"
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
	contri
	repo := newInmemRepository()
	m := NewService(repo)
	//u := newFixtureBook()
	_, err := m.CreateBook(u.Title, u.Author, u.Pages, u.Quantity)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
}

func Test_getBook(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeBookHandlers(r, *n, service)
	path, err := r.GetRoute("getBook").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/book/{id}", path)
	b := &entity.Book{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetBook(b.ID).
		Return(b, nil)
	handler := getBook(service)
	r.Handle("/v1/book/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/book/" + b.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.Book
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, b.ID, d.ID)
}