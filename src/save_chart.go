package src

import (
	"context"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/speps/go-hashids"
)

type ChartSettings struct {
	Data      []*Line `json:"data" firestore:"data"`
	StartDate *string `json:"start_date" firestore:"start_date"`
	EndDate   *string `json:"end_date" firestore:"end_date"`
	XStat     *string `json:"x_stat" firestore:"x_stat"`
	YStat     *string `json:"y_stat" firestore:"y_stat"`
}

func SaveChart(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body := readRequest(r)
	if body.Settings == nil {
		panic("Failed to parse settings")
	}

	log.Println("Save Chart")

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	counterRef := client.Collection("charts").Doc("counter")
	_, err := counterRef.Update(ctx, []firestore.Update{{Path: "count", Value: firestore.Increment(1)}})
	if err != nil {
		panic("Failed to increment save counter: " + err.Error())
	}

	counter, err := counterRef.Get(ctx)
	if err != nil {
		panic("Failed to get save counter: " + err.Error())
	}

	count := int(counter.Data()["count"].(int64))

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{count})

	_, err = client.Collection("charts").Doc(id).Set(ctx, map[string]interface{}{
		"data":       body.Settings.Data,
		"start_date": body.Settings.StartDate,
		"end_date":   body.Settings.EndDate,
		"x_stat":     body.Settings.XStat,
		"y_stat":     body.Settings.YStat,
	})
	if err != nil {
		panic("Failed saving chart to database: " + err.Error())
	}

	_, err = w.Write([]byte("unchartd.io/?chart=" + id))
	if err != nil {
		panic("Failed to return chart id")
	}

	elapsed := time.Since(start)
	log.Printf("Save Chart took %s", elapsed)
}
