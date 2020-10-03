package src

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"io"
	"log"
	"net/http"
	"strings"
)

func getFirebaseClient(ctx context.Context) *firestore.Client {
	// Use the application default credentials
	conf := &firebase.Config{ProjectID: "covid-charts-289323"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func readRequest(r *http.Request) *Request {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := &Request{}
	err = json.Unmarshal([]byte(buf.String()), body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

type Request struct {
	Locations    []string `json:"locations"`
	LocationType *string  `json:"location_type,omitempty"`
	XStat        *string  `json:"x_stat"`
	YStat        *string  `json:"y_stat"`
}
