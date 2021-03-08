package geolocation

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

type FirestoreRepository interface {
	AddGeolocation(ctx context.Context, geolocation geolocation) error
	GetGeolocationByIp(ctx context.Context, ipAddress string) (*geolocation, error)
}

type firestoreRepository struct {
	firestoreClient *firestore.Client
}

func NewGeolocationFirestoreRepository(
	firestoreClient *firestore.Client,
) FirestoreRepository {
	return &firestoreRepository{
		firestoreClient: firestoreClient,
	}
}

func (r firestoreRepository) geolocationCollection() *firestore.CollectionRef {
	return r.firestoreClient.Collection("geolocation")
}

func (r *firestoreRepository) AddGeolocation(ctx context.Context, geolocation geolocation) error {
	_, err := r.geolocationCollection().Doc(geolocation.uuid).Set(ctx, geolocation)
	return err
}

func (r *firestoreRepository) GetGeolocationByIp(ctx context.Context, ipAddress string) (*geolocation, error) {
	iter := r.geolocationCollection().Select("ipAddress", "==", ipAddress).Limit(1).Documents(ctx)

	doc, err := iter.Next()

	if err != nil {
		return nil, err
	}
	fmt.Sprintln(doc.Data(), nil)

	return nil, nil
}
