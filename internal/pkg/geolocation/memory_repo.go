package geolocation

import "github.com/google/uuid"

type inmemRepository struct {
	m map[uuid.UUID]*geolocation
}

func newInmemRepository() *inmemRepository {
	var m = map[entity.ID]*entity.Book{}
	return &inmemRepository{
		m: m,
	}
}

func (r *inmemRepository) AddGeolocation(geolocations []*geolocation) error {
	r.m[] = e
	return nil
}

func (r *inmemRepository) GetGeolocationByIp(ipAddress string) (*geolocation, error){
	if r.m[ipAddress] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

