package core_test

import (
	"context"
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/itok01/e-stat-go/core"
)

type mockHttpClientGetMetaInfoList struct{}

//go:embed testmock/get_meta_info.xml
var responseGetMetaInfoList []byte

func (hc *mockHttpClientGetMetaInfoList) Get(ctx context.Context, path string, query any) (int, []byte, error) {
	return 200, responseGetMetaInfoList, nil
}

func (hc *mockHttpClientGetMetaInfoList) Post(ctx context.Context, path string, data any) (int, []byte, error) {
	return 200, nil, nil
}

func (hc *mockHttpClientGetMetaInfoList) PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error) {
	return 200, nil, nil
}

func TestGetMetaInfoList(t *testing.T) {
	tests := []struct {
		name string
		arg  core.ParamsGetMetaInfoList
		want core.ResponseGetMetaInfoListRoot
	}{
		{
			name: "Unmarshal GetMetaInfoList",
			arg: core.ParamsGetMetaInfoList{
				StatsDataId: "0003109741",
			},
			want: core.ResponseGetMetaInfoListRoot{
				core.ResponseGetMetaInfoList{
					Result: core.ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 3, 1, 18, 39, 752000000, time.Local),
					},
					Parameter: core.ResponseGetMetaInfoListParameter{
						CommonParams: core.CommonParams{
							Lang: "J",
						},
						ParamsGetMetaInfoList: core.ParamsGetMetaInfoList{
							StatsDataId: "0003109741",
						},
					},
					DataList: core.ResponseGetMetaInfoMetaDataList{
						Number: 0,
						Table: core.TableInf{ID: "0003109741",
							StatName: core.StatName{
								Code: "00100409",
								Name: "国民経済計算",
							},
							GovOrg: core.GovOrg{
								Code: "00100",
								Name: "内閣府",
							},
							StatisticsName: "四半期別ＧＤＰ速報",
							Title: core.Title{
								Name: "国内総生産（支出側）及び各需要項目 名目原系列（1994年1Q～） 2015暦年基準",
							},
							Cycle:       "四半期",
							SurveyDate:  "202204-202206",
							OpenDate:    "2022-09-08",
							SmallArea:   0,
							CollectArea: "該当なし",
							MainCategory: core.MainCategory{
								Code: "07",
								Name: "企業・家計・経済",
							},
							SubCategory: core.SubCategory{
								Code: "05",
								Name: "国民経済計算",
							},
							OverallTotalNumber: 2508,
							UpdatedDate:        "2022-09-16",
							StatisticsNameSpec: core.StatisticsNameSpec{
								TabulationCategory: "四半期別ＧＤＰ速報",
							},
							TitleSpec: core.TitleSpec{
								TableCategory:     "国内総生産（支出側）及び各需要項目",
								TableName:         "名目原系列（1994年1Q～）",
								TableSubCategory1: "2015暦年基準",
							},
						},
						Class: core.ClassInf{
							ClassObj: []core.ClassObj{
								{
									ID:          "tab",
									Name:        "表章項目",
									Description: "Excelの書式設定で統計表の数値を\"-0.0\"としている場合、データベース上\"0.0\"として収録されているため、Excel統計表の数値とは必ずしも一致しない。",
									Class: []core.ClassObjClass{
										{Code: "11", Name: "金額", Unit: "10億円"},
									},
								},
								{
									ID:   "cat01",
									Name: "国内総生産_名目原系列",
									Class: []core.ClassObjClass{
										{Code: "11", Name: "国内総生産(支出側)", Level: "1"},
										{Code: "12", Name: "民間最終消費支出", Level: "1"},
										{Code: "13", Name: "民間最終消費支出_家計最終消費支出", Level: "2", ParentCode: "12"},
										{Code: "14", Name: "民間最終消費支出_家計最終消費支出_除く持ち家の帰属家賃", Level: "3", ParentCode: "13"},
										{Code: "15", Name: "民間住宅", Level: "1"},
										{Code: "16", Name: "民間企業設備", Level: "1"},
										{Code: "17", Name: "民間在庫変動", Level: "1"},
										{Code: "18", Name: "政府最終消費支出", Level: "1"},
										{Code: "19", Name: "公的固定資本形成", Level: "1"},
										{Code: "20", Name: "公的在庫変動", Level: "1"},
										{Code: "21", Name: "財貨・サービス_純輸出", Level: "1"},
										{Code: "22", Name: "財貨・サービス_輸出", Level: "1"},
										{Code: "23", Name: "財貨・サービス_輸入", Level: "1"},
										{Code: "24", Name: "<参考>海外からの所得_純受取", Level: "1"},
										{Code: "25", Name: "<参考>海外からの所得_受取", Level: "1"},
										{Code: "26", Name: "<参考>海外からの所得_支払", Level: "1"},
										{Code: "27", Name: "<参考>国民総所得", Level: "1"},
										{Code: "28", Name: "<参考>国内需要", Level: "1"},
										{Code: "29", Name: "<参考>民間需要", Level: "1"},
										{Code: "30", Name: "<参考>公的需要", Level: "1"},
										{Code: "31", Name: "<参考>総固定資本形成", Level: "1"},
										{Code: "32", Name: "<参考>最終需要", Level: "1"},
									},
								},
								{
									ID:   "time",
									Name: "時間軸（四半期）",
									Class: []core.ClassObjClass{
										{Code: "1994000103", Name: "1994年1～3月期", Level: "3", ParentCode: "1994010000"},
										{Code: "1994000406", Name: "1994年4～6月期", Level: "3", ParentCode: "1994010000"},
										{Code: "1994000709", Name: "1994年7～9月期", Level: "3", ParentCode: "1994020000"},
										{Code: "1994001012", Name: "1994年10～12月期", Level: "3", ParentCode: "1994020000"},
										{Code: "1995000103", Name: "1995年1～3月期", Level: "3", ParentCode: "1995010000"},
										{Code: "1995000406", Name: "1995年4～6月期", Level: "3", ParentCode: "1995010000"},
										{Code: "1995000709", Name: "1995年7～9月期", Level: "3", ParentCode: "1995020000"},
										{Code: "1995001012", Name: "1995年10～12月期", Level: "3", ParentCode: "1995020000"},
										{Code: "1996000103", Name: "1996年1～3月期", Level: "3", ParentCode: "1996010000"},
										{Code: "1996000406", Name: "1996年4～6月期", Level: "3", ParentCode: "1996010000"},
										{Code: "1996000709", Name: "1996年7～9月期", Level: "3", ParentCode: "1996020000"},
										{Code: "1996001012", Name: "1996年10～12月期", Level: "3", ParentCode: "1996020000"},
										{Code: "1997000103", Name: "1997年1～3月期", Level: "3", ParentCode: "1997010000"},
										{Code: "1997000406", Name: "1997年4～6月期", Level: "3", ParentCode: "1997010000"},
										{Code: "1997000709", Name: "1997年7～9月期", Level: "3", ParentCode: "1997020000"},
										{Code: "1997001012", Name: "1997年10～12月期", Level: "3", ParentCode: "1997020000"},
										{Code: "1998000103", Name: "1998年1～3月期", Level: "3", ParentCode: "1998010000"},
										{Code: "1998000406", Name: "1998年4～6月期", Level: "3", ParentCode: "1998010000"},
										{Code: "1998000709", Name: "1998年7～9月期", Level: "3", ParentCode: "1998020000"},
										{Code: "1998001012", Name: "1998年10～12月期", Level: "3", ParentCode: "1998020000"},
										{Code: "1999000103", Name: "1999年1～3月期", Level: "3", ParentCode: "1999010000"},
										{Code: "1999000406", Name: "1999年4～6月期", Level: "3", ParentCode: "1999010000"},
										{Code: "1999000709", Name: "1999年7～9月期", Level: "3", ParentCode: "1999020000"},
										{Code: "1999001012", Name: "1999年10～12月期", Level: "3", ParentCode: "1999020000"},
										{Code: "2000000103", Name: "2000年1～3月期", Level: "3", ParentCode: "2000010000"},
										{Code: "2000000406", Name: "2000年4～6月期", Level: "3", ParentCode: "2000010000"},
										{Code: "2000000709", Name: "2000年7～9月期", Level: "3", ParentCode: "2000020000"},
										{Code: "2000001012", Name: "2000年10～12月期", Level: "3", ParentCode: "2000020000"},
										{Code: "2001000103", Name: "2001年1～3月期", Level: "3", ParentCode: "2001010000"},
										{Code: "2001000406", Name: "2001年4～6月期", Level: "3", ParentCode: "2001010000"},
										{Code: "2001000709", Name: "2001年7～9月期", Level: "3", ParentCode: "2001020000"},
										{Code: "2001001012", Name: "2001年10～12月期", Level: "3", ParentCode: "2001020000"},
										{Code: "2002000103", Name: "2002年1～3月期", Level: "3", ParentCode: "2002010000"},
										{Code: "2002000406", Name: "2002年4～6月期", Level: "3", ParentCode: "2002010000"},
										{Code: "2002000709", Name: "2002年7～9月期", Level: "3", ParentCode: "2002020000"},
										{Code: "2002001012", Name: "2002年10～12月期", Level: "3", ParentCode: "2002020000"},
										{Code: "2003000103", Name: "2003年1～3月期", Level: "3", ParentCode: "2003010000"},
										{Code: "2003000406", Name: "2003年4～6月期", Level: "3", ParentCode: "2003010000"},
										{Code: "2003000709", Name: "2003年7～9月期", Level: "3", ParentCode: "2003020000"},
										{Code: "2003001012", Name: "2003年10～12月期", Level: "3", ParentCode: "2003020000"},
										{Code: "2004000103", Name: "2004年1～3月期", Level: "3", ParentCode: "2004010000"},
										{Code: "2004000406", Name: "2004年4～6月期", Level: "3", ParentCode: "2004010000"},
										{Code: "2004000709", Name: "2004年7～9月期", Level: "3", ParentCode: "2004020000"},
										{Code: "2004001012", Name: "2004年10～12月期", Level: "3", ParentCode: "2004020000"},
										{Code: "2005000103", Name: "2005年1～3月期", Level: "3", ParentCode: "2005010000"},
										{Code: "2005000406", Name: "2005年4～6月期", Level: "3", ParentCode: "2005010000"},
										{Code: "2005000709", Name: "2005年7～9月期", Level: "3", ParentCode: "2005020000"},
										{Code: "2005001012", Name: "2005年10～12月期", Level: "3", ParentCode: "2005020000"},
										{Code: "2006000103", Name: "2006年1～3月期", Level: "3", ParentCode: "2006010000"},
										{Code: "2006000406", Name: "2006年4～6月期", Level: "3", ParentCode: "2006010000"},
										{Code: "2006000709", Name: "2006年7～9月期", Level: "3", ParentCode: "2006020000"},
										{Code: "2006001012", Name: "2006年10～12月期", Level: "3", ParentCode: "2006020000"},
										{Code: "2007000103", Name: "2007年1～3月期", Level: "3", ParentCode: "2007010000"},
										{Code: "2007000406", Name: "2007年4～6月期", Level: "3", ParentCode: "2007010000"},
										{Code: "2007000709", Name: "2007年7～9月期", Level: "3", ParentCode: "2007020000"},
										{Code: "2007001012", Name: "2007年10～12月期", Level: "3", ParentCode: "2007020000"},
										{Code: "2008000103", Name: "2008年1～3月期", Level: "3", ParentCode: "2008010000"},
										{Code: "2008000406", Name: "2008年4～6月期", Level: "3", ParentCode: "2008010000"},
										{Code: "2008000709", Name: "2008年7～9月期", Level: "3", ParentCode: "2008020000"},
										{Code: "2008001012", Name: "2008年10～12月期", Level: "3", ParentCode: "2008020000"},
										{Code: "2009000103", Name: "2009年1～3月期", Level: "3", ParentCode: "2009010000"},
										{Code: "2009000406", Name: "2009年4～6月期", Level: "3", ParentCode: "2009010000"},
										{Code: "2009000709", Name: "2009年7～9月期", Level: "3", ParentCode: "2009020000"},
										{Code: "2009001012", Name: "2009年10～12月期", Level: "3", ParentCode: "2009020000"},
										{Code: "2010000103", Name: "2010年1～3月期", Level: "3", ParentCode: "2010010000"},
										{Code: "2010000406", Name: "2010年4～6月期", Level: "3", ParentCode: "2010010000"},
										{Code: "2010000709", Name: "2010年7～9月期", Level: "3", ParentCode: "2010020000"},
										{Code: "2010001012", Name: "2010年10～12月期", Level: "3", ParentCode: "2010020000"},
										{Code: "2011000103", Name: "2011年1～3月期", Level: "3", ParentCode: "2011010000"},
										{Code: "2011000406", Name: "2011年4～6月期", Level: "3", ParentCode: "2011010000"},
										{Code: "2011000709", Name: "2011年7～9月期", Level: "3", ParentCode: "2011020000"},
										{Code: "2011001012", Name: "2011年10～12月期", Level: "3", ParentCode: "2011020000"},
										{Code: "2012000103", Name: "2012年1～3月期", Level: "3", ParentCode: "2012010000"},
										{Code: "2012000406", Name: "2012年4～6月期", Level: "3", ParentCode: "2012010000"},
										{Code: "2012000709", Name: "2012年7～9月期", Level: "3", ParentCode: "2012020000"},
										{Code: "2012001012", Name: "2012年10～12月期", Level: "3", ParentCode: "2012020000"},
										{Code: "2013000103", Name: "2013年1～3月期", Level: "3", ParentCode: "2013010000"},
										{Code: "2013000406", Name: "2013年4～6月期", Level: "3", ParentCode: "2013010000"},
										{Code: "2013000709", Name: "2013年7～9月期", Level: "3", ParentCode: "2013020000"},
										{Code: "2013001012", Name: "2013年10～12月期", Level: "3", ParentCode: "2013020000"},
										{Code: "2014000103", Name: "2014年1～3月期", Level: "3", ParentCode: "2014010000"},
										{Code: "2014000406", Name: "2014年4～6月期", Level: "3", ParentCode: "2014010000"},
										{Code: "2014000709", Name: "2014年7～9月期", Level: "3", ParentCode: "2014020000"},
										{Code: "2014001012", Name: "2014年10～12月期", Level: "3", ParentCode: "2014020000"},
										{Code: "2015000103", Name: "2015年1～3月期", Level: "3", ParentCode: "2015010000"},
										{Code: "2015000406", Name: "2015年4～6月期", Level: "3", ParentCode: "2015010000"},
										{Code: "2015000709", Name: "2015年7～9月期", Level: "3", ParentCode: "2015020000"},
										{Code: "2015001012", Name: "2015年10～12月期", Level: "3", ParentCode: "2015020000"},
										{Code: "2016000103", Name: "2016年1～3月期", Level: "3", ParentCode: "2016010000"},
										{Code: "2016000406", Name: "2016年4～6月期", Level: "3", ParentCode: "2016010000"},
										{Code: "2016000709", Name: "2016年7～9月期", Level: "3", ParentCode: "2016020000"},
										{Code: "2016001012", Name: "2016年10～12月期", Level: "3", ParentCode: "2016020000"},
										{Code: "2017000103", Name: "2017年1～3月期", Level: "3", ParentCode: "2017010000"},
										{Code: "2017000406", Name: "2017年4～6月期", Level: "3", ParentCode: "2017010000"},
										{Code: "2017000709", Name: "2017年7～9月期", Level: "3", ParentCode: "2017020000"},
										{Code: "2017001012", Name: "2017年10～12月期", Level: "3", ParentCode: "2017020000"},
										{Code: "2018000103", Name: "2018年1～3月期", Level: "3", ParentCode: "2018010000"},
										{Code: "2018000406", Name: "2018年4～6月期", Level: "3", ParentCode: "2018010000"},
										{Code: "2018000709", Name: "2018年7～9月期", Level: "3", ParentCode: "2018020000"},
										{Code: "2018001012", Name: "2018年10～12月期", Level: "3", ParentCode: "2018020000"},
										{Code: "2019000103", Name: "2019年1～3月期", Level: "3", ParentCode: "2019010000"},
										{Code: "2019000406", Name: "2019年4～6月期", Level: "3", ParentCode: "2019010000"},
										{Code: "2019000709", Name: "2019年7～9月期", Level: "3", ParentCode: "2019020000"},
										{Code: "2019001012", Name: "2019年10～12月期", Level: "3", ParentCode: "2019020000"},
										{Code: "2020000103", Name: "2020年1～3月期", Level: "3", ParentCode: "2020010000"},
										{Code: "2020000406", Name: "2020年4～6月期", Level: "3", ParentCode: "2020010000"},
										{Code: "2020000709", Name: "2020年7～9月期", Level: "3", ParentCode: "2020020000"},
										{Code: "2020001012", Name: "2020年10～12月期", Level: "3", ParentCode: "2020020000"},
										{Code: "2021000103", Name: "2021年1～3月期", Level: "3", ParentCode: "2021010000"},
										{Code: "2021000406", Name: "2021年4～6月期", Level: "3", ParentCode: "2021010000"},
										{Code: "2021000709", Name: "2021年7～9月期", Level: "3", ParentCode: "2021020000"},
										{Code: "2021001012", Name: "2021年10～12月期", Level: "3", ParentCode: "2021020000"},
										{Code: "2022000103", Name: "2022年1～3月期", Level: "3", ParentCode: "2022010000"},
										{Code: "2022000406", Name: "2022年4～6月期", Level: "3", ParentCode: "2022010000"},
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
	hc := mockHttpClientGetMetaInfoList{}
	ac := core.NewApiClient(&hc, core.CommonParams{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ac.GetMetaInfoList(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetMetaInfoList() = %v, want %v", got, tt.want)
			}
		})
	}
}
