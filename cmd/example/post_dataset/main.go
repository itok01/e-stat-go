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

	data, err := ac.PostDataset(ctx, estatgo.ParamsPostDataset{
		StatsDataId: "0003010900",
		DataSetName: "住宅・土地統計調査　データセット１",
		NarrowingConditon: estatgo.NarrowingConditon{
			CategoryCondition: estatgo.CategoryCondition{
				CodeCat01From: "01",
				CodeCat01To:   "02",
				CodeCat02:     "03",
			},
			AreaCondition: estatgo.AreaCondition{
				LevelArea: "1",
			},
		},
		OpenSpecified: "1",
		ProcessMode:   "E",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", data)
}
