package geolocation

type Service interface {
	ImportGeoloacationData(geolocation geolocationModel)
	GetGeolocationByIp(ipAdrress string)
}

type service struct {
	gr FirestoreRepository
}

func newService(repository firestoreRepository) Service  {
	return &service{repository}
}

func (s service) ImportGeoloacationData(geolocation geolocationModel) {
	panic("implement me")
}

func (s service) GetGeolocationByIp(ipAdrress string) {
	panic("implement me")
}

