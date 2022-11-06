package core_test

import (
	"context"
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/itok01/e-stat-go/core"
)

func TestGetDataCatalog(t *testing.T) {
	tests := []struct {
		name string
		arg  core.ParamsGetDataCatalog
		want core.ResponseGetDataCatalogRoot
	}{
		{
			name: "Unmarshal",
			arg: core.ParamsGetDataCatalog{
				SurveyYears: "202201",
				Limit:       2,
			},
			want: core.ResponseGetDataCatalogRoot{
				ResponseGetDataCatalog: core.ResponseGetDataCatalog{
					Result: core.ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 3, 23, 46, 29, 946000000, time.Local),
					},
					Parameter: core.ResponseGetDataCatalogParameter{
						CommonParams: core.CommonParams{
							Lang: "J",
						},
						ParamsGetDataCatalog: core.ParamsGetDataCatalog{
							DataType: "XLS",
							Limit:    1,
						},
					},
					DataCatalogList: core.DataCatalogListInf{
						Number: 42198,
						Result: core.ResultInf{FromNumber: 1,
							ToNumber: 1,
							NextKey:  2,
						},
						DataCatalog: []core.DataCatalogInf{
							{
								ID: "000001120179",
								Dataset: core.Dataset{
									StatName: core.StatName{
										Code: "00000002",
										Name: "一般職国家公務員在職状況統計表（人事統計報告）",
									},
									Organization: core.Organization{
										Code: "00000",
										Name: "内閣官房",
									},
									Title: core.DatasetTitle{
										Name: "一般職国家公務員在職状況統計表（平成２１年７月１日現在）_検察官在職状況統計表_2009年度",
										StatisticsNameSpec: core.StatisticsNameSpec{
											TabulationCategory:     "一般職国家公務員在職状況統計表（平成２１年７月１日現在）",
											TabulationSubCategory1: "検察官在職状況統計表",
											TabulationSubCategory2: "",
											TabulationSubCategory3: "",
											TabulationSubCategory4: "",
											TabulationSubCategory5: "",
										},
										Cycle:       "",
										SurveyDate:  "200904-201003",
										CollectArea: "該当なし",
									},
									Publisher:         "内閣官房",
									ContactPoint:      "内閣人事局",
									Creator:           "内閣人事局",
									ReleaseDate:       "2016-11-01",
									LastModifiedDate:  "2016-11-01",
									FrequencyOfUpdate: "",
									LandingPage:       "https://www.e-stat.go.jp/stat-search/files?layout=datalist&toukei=00000002&tstat=000001065638&cycle=0&tclass1=000001065620",
								},
								Resources: []core.Resources{
									{
										Resource: core.Resource{
											ID: "000006912081",
											Title: core.ResourceTitle{
												Name:        "5_省庁別、区分別検察官数（第５表）",
												TableNumber: 5,
												TitleSpec: core.TitleSpec{
													TableName: "省庁別、区分別検察官数（第５表）",
												},
											},
											URL:               "https://www.e-stat.go.jp/stat-search/file-download?&statInfId=000026290978&fileKind=0",
											Format:            "XLS",
											ReleaseDate:       "2016-11-01",
											LastModifiedDate:  "2016-11-01",
											ResourceLicenceID: "Government of Japan Standard Terms of Use",
											Language:          "J",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	ctx := context.Background()
	hc := mockHttpClient{}
	ac := core.NewApiClient(&hc, core.CommonParams{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ac.GetDataCatalog(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetDataCatalog() = %v, want %v", got, tt.want)
			}
		})
	}
}
