package main

import (
	"context"
	"log"

	estatgo "github.com/itok01/e-stat-go/core"
)

const (
	AppID = "520acac534269b12172b75119c6aa11938575a53"
)

func main() {
	ctx := context.Background()

	hc := estatgo.NewClient(true)

	ac := estatgo.NewApiClient(hc, estatgo.CommonParams{
		AppID: AppID,
	})

	data, err := ac.GetMetaInfoList(ctx, estatgo.ParamsGetMetaInfoList{
		StatsDataId: "0003109741",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", data)
}
