package core

import (
	"context"
	_ "embed"
	"reflect"
	"testing"
	"time"
)

type mockHttpClientGetDataCatalog struct{}

//go:embed testmock/get_data_catalog.xml
var responseGetDataCatalog []byte

func (hc *mockHttpClientGetDataCatalog) Get(ctx context.Context, path string, query any) (int, []byte, error) {
	return 200, responseGetDataCatalog, nil
}

func (hc *mockHttpClientGetDataCatalog) Post(ctx context.Context, path string, data any) (int, []byte, error) {
	return 200, nil, nil
}

func (hc *mockHttpClientGetDataCatalog) PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error) {
	return 200, nil, nil
}

func TestGetDataCatalog(t *testing.T) {
	tests := []struct {
		name string
		arg  ParamsGetDataCatalog
		want ResponseGetDataCatalogRoot
	}{
		{
			name: "Unmarshal GetDataCatalog",
			arg: ParamsGetDataCatalog{
				SurveyYears: "202201",
				Limit:       2,
			},
			want: ResponseGetDataCatalogRoot{
				ResponseGetDataCatalog: ResponseGetDataCatalog{
					Result: ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 3, 23, 46, 29, 946000000, time.Local),
					},
					Parameter: ResponseGetDataCatalogParameter{
						CommonParams: CommonParams{
							Lang: "J",
						},
						ParamsGetDataCatalog: ParamsGetDataCatalog{
							DataType: "XLS",
							Limit:    1,
						},
					},
					DataCatalogList: DataCatalogListInf{
						Number: 42198,
						Result: ResultInf{FromNumber: 1,
							ToNumber: 1,
							NextKey:  2,
						},
						DataCatalog: []DataCatalogInf{
							{
								ID: "000001120179",
								Dataset: Dataset{
									StatName: StatName{
										Code: "00000002",
										Name: "一般職国家公務員在職状況統計表（人事統計報告）",
									},
									Organization: Organization{
										Code: "00000",
										Name: "内閣官房",
									},
									Title: DatasetTitle{
										Name: "一般職国家公務員在職状況統計表（平成２１年７月１日現在）_検察官在職状況統計表_2009年度",
										StatisticsNameSpec: StatisticsNameSpec{
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
								Resources: []Resources{
									{
										Resource: Resource{
											ID: "000006912081",
											Title: ResourceTitle{
												Name:        "5_省庁別、区分別検察官数（第５表）",
												TableNumber: 5,
												TitleSpec: TitleSpec{
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
	hc := mockHttpClientGetDataCatalog{}
	ac := NewApiClient(&hc, CommonParams{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ac.GetDataCatalog(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetDataCatalog() = %v, want %v", got, tt.want)
			}
		})
	}
}
