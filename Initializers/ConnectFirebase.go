package initializers

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func ConnectFirebase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./Config/serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	FirebaseApp = app
	log.Println("Firebase initialized successfully")
}
