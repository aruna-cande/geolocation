package geolocation_api

import (
	"context"
	"Geolocation/internal/pkg/geolocation"
	"cloud.google.com/go/firestore"
	"fmt"
	"os"
)

func main(){
	ctx := context.Background()
	fsClient, err := firestore.NewClient(ctx, os.Getenv("test"))

	if err != nil{
		fmt.Sprintln("erro", err)
	}

	repo = geolocation.FirestoreRepository(&fsClient)

}
