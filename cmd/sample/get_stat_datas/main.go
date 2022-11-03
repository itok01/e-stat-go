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

	data, err := ac.GetStatsDatas(ctx, estatgo.ParamsGetStatsDatas{
		MetaGetFlg: "Y",
	}, []estatgo.StatsDatasSpec{
		{
			StatsDataId:   "0003090287",
			StartPosition: 1,
			Limit:         10,
		},
		{
			StatsDataId:   "0003010900",
			StartPosition: 1,
			Limit:         10,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", data)
}
