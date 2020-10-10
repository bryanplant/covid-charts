package src

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type ChartSettings struct {
	Data          []*Line    `json:"data" firestore:"data"`
	StartDate     *time.Time `json:"-" firestore:"start_date"`
	StartDateJson *string    `json:"start_date" firestore:"_"`
	EndDate       *time.Time `json:"-" firestore:"end_date"`
	EndDateJson   *string    `json:"end_date" firestore:"_"`
	XStat         *string    `json:"x_stat" firestore:"x_stat"`
	YStat         *string    `json:"y_stat" firestore:"y_stat"`
}

func (cs *ChartSettings) populateDates() {
	date, err := parseDate(*cs.StartDateJson)
	if err != nil {
		panic(fmt.Sprintf("Could not parse start date: " + *cs.StartDateJson))
	}
	cs.StartDate = date

	date, err = parseDate(*cs.EndDateJson)
	if err != nil {
		panic(fmt.Sprintf("Could not parse end date: " + *cs.EndDateJson))
	}
	cs.EndDate = date
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
	body.Settings.populateDates()

	log.Println("Save Chart")

	ctx := context.Background()
	client := getFirebaseClient(ctx)

	id := uuid.New().String()
	_, err := client.Collection("charts").Doc(id).Set(ctx, map[string]interface{}{
		"data":       body.Settings.Data,
		"start_date": body.Settings.StartDate,
		"end_date":   body.Settings.EndDate,
		"x_stat":     body.Settings.XStat,
		"y_stat":     body.Settings.YStat,
	})
	if err != nil {
		panic("Failed saving chart to database: " + err.Error())
	}

	doc, err := client.Collection("charts").Doc(id).Get(ctx)
	if err != nil {
		panic("a")
	}
	var load ChartSettings
	err = doc.DataTo(&load)
	if err != nil {
		panic("b")
	}

	_, err = w.Write([]byte("unchartd.io/?chart=" + id))
	if err != nil {
		panic("Failed to return chart id")
	}

	elapsed := time.Since(start)
	log.Printf("Save Chart took %s", elapsed)
}
