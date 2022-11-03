package core

import (
	"context"
	_ "embed"
	"reflect"
	"testing"
	"time"
)

type mockHttpClientGetStatsList struct{}

//go:embed testmock/get_stats_list.xml
var responseGetStatsList []byte

func (hc *mockHttpClientGetStatsList) Get(ctx context.Context, path string, query any) (int, []byte, error) {
	return 200, responseGetStatsList, nil
}

func (hc *mockHttpClientGetStatsList) Post(ctx context.Context, path string, data any) (int, []byte, error) {
	return 200, nil, nil
}

func (hc *mockHttpClientGetStatsList) PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error) {
	return 200, nil, nil
}

func TestGetStatsList(t *testing.T) {
	tests := []struct {
		name string
		arg  ParamsGetStatsList
		want ResponseGetStatsListRoot
	}{
		{
			name: "Unmarshal GetStatsList",
			arg: ParamsGetStatsList{
				SurveyYears: "202201",
				Limit:       2,
			},
			want: ResponseGetStatsListRoot{
				ResponseGetStatsList: ResponseGetStatsList{
					Result: ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.October, 29, 18, 20, 5, 880000000, time.Local),
					},
					Parameter: ResponseGetStatsListParameter{
						CommonParams: CommonParams{
							Lang: "J",
						},
						ParamsGetStatsList: ParamsGetStatsList{
							SurveyYears: "202201",
							Limit:       2,
						},
					},
					DataList: &ResponseGetStatsListDataList{
						Number: 101,
						Result: ResultInf{
							FromNumber: 1,
							ToNumber:   2,
							NextKey:    3,
						},
						Table: []TableInf{
							{
								ID: "0003355224",
								StatName: StatName{
									Code: "00100401",
									Name: "機械受注統計調査",
								},
								GovOrg: GovOrg{
									Code: "00100",
									Name: "内閣府",
								},
								StatisticsName: "機械受注統計調査 主要長期時系列統計表（2005/4～）",
								Title: Title{
									Number: "",
									Name:   "長期系列\u3000主要需要者別機械受注額(原系列・年度)",
								},
								Cycle:       "年度次",
								SurveyDate:  "202104-202203",
								OpenDate:    "2022-05-19",
								SmallArea:   0,
								CollectArea: "該当なし",
								MainCategory: MainCategory{
									Code: "05",
									Name: "鉱工業",
								},
								SubCategory: SubCategory{
									Code: "02",
									Name: "製造業",
								},
								OverallTotalNumber: 289,
								UpdatedDate:        "2022-10-19",
								StatisticsNameSpec: StatisticsNameSpec{
									TabulationCategory:     "機械受注統計調査",
									TabulationSubCategory1: "主要長期時系列統計表（2005/4～）",
								},
								TitleSpec: TitleSpec{
									TableName:        "長期系列\u3000主要需要者別機械受注額(原系列・年度)",
									TableExplanation: "Volatile orders:Orders for ships and those from electric power companies.",
								},
							},
							{
								ID: "0003355228",
								StatName: StatName{
									Code: "00100401",
									Name: "機械受注統計調査",
								},
								GovOrg: GovOrg{
									Code: "00100",
									Name: "内閣府",
								},
								StatisticsName: "機械受注統計調査 主要長期時系列統計表（2005/4～）",
								Title: Title{
									Number: "",
									Name:   "長期系列\u3000機種別機械受注額(大分類・中分類、原系列・年度)",
								},
								Cycle:       "年度次",
								SurveyDate:  "202104-202203",
								OpenDate:    "2022-05-19",
								SmallArea:   0,
								CollectArea: "該当なし",
								MainCategory: MainCategory{
									Code: "05",
									Name: "鉱工業",
								},
								SubCategory: SubCategory{
									Code: "02",
									Name: "製造業",
								},
								OverallTotalNumber: 578,
								UpdatedDate:        "2022-10-19",
								StatisticsNameSpec: StatisticsNameSpec{
									TabulationCategory:     "機械受注統計調査",
									TabulationSubCategory1: "主要長期時系列統計表（2005/4～）",
								},
								TitleSpec: TitleSpec{
									TableName:        "長期系列\u3000機種別機械受注額(大分類・中分類、原系列・年度)",
									TableExplanation: "（備考）「電子計算機等」は、「電子計算機」と「半導体製造装置」の合計。",
								},
							},
						},
					},
				},
			},
		},
	}

	ctx := context.Background()
	hc := mockHttpClientGetStatsList{}
	ac := NewApiClient(&hc, CommonParams{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ac.GetStatsList(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetStatsList() = %v, want %v", got, tt.want)
			}
		})
	}
}
