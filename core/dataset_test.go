package core_test

import (
	"context"
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/itok01/e-stat-go/core"
)

func TestPostDataset(t *testing.T) {
	tests := []struct {
		name string
		arg  core.ParamsPostDataset
		want core.ResponsePostDatasetRoot
	}{
		{
			name: "Unmarshal",
			arg: core.ParamsPostDataset{
				StatsDataId: "0003010900",
				DataSetName: "住宅・土地統計調査　データセット１",
				NarrowingConditon: core.NarrowingConditon{
					CategoryCondition: core.CategoryCondition{
						CodeCat01From: "01",
						CodeCat01To:   "02",
						CodeCat02:     "03",
					},
					AreaCondition: core.AreaCondition{
						LevelArea: "1",
					},
				},
				OpenSpecified: "1",
				ProcessMode:   "E",
			},
			want: core.ResponsePostDatasetRoot{
				ResponsePostDataset: core.ResponsePostDataset{
					Result: core.ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 3, 21, 43, 10, 306000000, time.Local),
					},
					Parameter: core.ResponsePostDatasetParameter{
						CommonParams: core.CommonParams{
							Lang: "J",
						},
						StatsDataId: "0003010900",
						NarrowingConditon: core.NarrowingConditon{
							AreaCondition: core.AreaCondition{
								LevelArea: "1",
							},
							CategoryCondition: core.CategoryCondition{
								CodeCat01From: "01",
								CodeCat01To:   "02",
								CodeCat02:     "03",
							},
						},
						OpenSpecified: "1",
						ProcessMode:   "E",
						DataSetName:   "住宅・土地統計調査\u3000データセット１",
					},
					RefistInf: core.ResponsePostDatasetRegistInf{
						Mode:        "add",
						DatasetId:   "00200522-20221103214310",
						StatsDataId: "0003010900",
						PublicState: "yes",
						TotalNumber: 2,
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
			got, _ := ac.PostDataset(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("PostDataset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefDataset(t *testing.T) {
	tests := []struct {
		name string
		arg  core.ParamsRefDataset
		want core.ResponseRefDatasetRoot
	}{
		{
			name: "Unmarshal",
			arg: core.ParamsRefDataset{
				DataSetID: "CTCdemo-kokusei1",
			},
			want: core.ResponseRefDatasetRoot{
				ResponseRefDataset: core.ResponseRefDataset{
					Result: core.ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 4, 1, 51, 47, 507000000, time.Local),
					},
					Parameter: core.ResponseRefDatasetParameter{
						CommonParams: core.CommonParams{
							Lang: "J",
						},
					},
					Dataset: core.ResponseRefDatasetInf{
						ID:          "CTCdemo-kokusei1",
						DataSetName: "",
						PublicState: "YES",
						Result: core.ResponseRefDatasetResultInf{
							TotalNumber: 14382,
						},
						TableInf: core.TableInf{ID: "0003038586",
							StatName: core.StatName{
								Code: "00200521",
								Name: "国勢調査",
							},
							GovOrg: core.GovOrg{
								Code: "00200",
								Name: "総務省",
							},
							StatisticsName: "平成22年国勢調査 人口等基本集計（男女・年齢・配偶関係，世帯の構成，住居の状態など）",
							Title: core.Title{
								Number: "00100",
								Name:   "人口，人口増減，面積及び人口密度 全国，市部・郡部，都道府県，市部・郡部，支庁，郡計，市区町村・旧市町村，全域・人口集中地区",
							},
							Cycle:       "-",
							SurveyDate:  "201010",
							OpenDate:    "2011-10-26",
							SmallArea:   0,
							CollectArea: "該当なし",
							MainCategory: core.MainCategory{
								Code: "02",
								Name: "人口・世帯",
							},
							SubCategory: core.SubCategory{
								Code: "01",
								Name: "人口",
							},
							OverallTotalNumber: 0,
							UpdatedDate:        "2022-08-31",
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
			got, _ := ac.RefDataset(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("RefDataset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDatasetList(t *testing.T) {
	tests := []struct {
		name string
		arg  core.ParamsGetDatasetList
		want core.ResponseGetDatasetListRoot
	}{
		{
			name: "Unmarshal",
			arg:  core.ParamsGetDatasetList{},
			want: core.ResponseGetDatasetListRoot{
				ResponseGetDatasetList: core.ResponseGetDatasetList{
					Result: core.ResponseResult{
						Status:   0,
						ErrorMsg: "正常に終了しました。",
						Date:     time.Date(2022, time.November, 3, 23, 11, 22, 488000000, time.Local),
					},
					Parameter: core.ResponseRefDatasetParameter{
						CommonParams: core.CommonParams{
							Lang: "J",
						},
					},
					DatasetList: core.ResponseRefDatasetListInf{
						Number: 17,
						Dataset: []core.ResponseRefDatasetInf{
							{
								ID:          "00200502-20181005113922",
								DataSetName: "test01010101",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 40002,
								},
								TableInf: core.TableInf{

									ID: "0000020201",
									StatName: core.StatName{
										Code: "00200502",
										Name: "社会・人口統計体系",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "市区町村データ 基礎データ（廃置分合処理済）",
									Title: core.Title{
										Number: "0000020201",
										Name:   "Ａ\u3000人口・世帯",
									},
									Cycle:       "年度次",
									SurveyDate:  "0",
									OpenDate:    "2022-06-21",
									SmallArea:   0,
									CollectArea: "市区町村",
									MainCategory: core.MainCategory{
										Code: "99",
										Name: "その他",
									},
									SubCategory: core.SubCategory{
										Code: "99",
										Name: "その他",
									},
									OverallTotalNumber: 1567404,
									UpdatedDate:        "2022-06-21",
								},
							},
							{
								ID:          "CTCdemo-kokusei1",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 14382,
								},
								TableInf: core.TableInf{
									ID: "0003038586",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成22年国勢調査 人口等基本集計（男女・年齢・配偶関係，世帯の構成，住居の状態など）",
									Title: core.Title{
										Number: "00100",
										Name:   "人口，人口増減，面積及び人口密度 全国，市部・郡部，都道府県，市部・郡部，支庁，郡計，市区町村・旧市町村，全域・人口集中地区",
									},
									Cycle:       "-",
									SurveyDate:  "201010",
									OpenDate:    "2011-10-26",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 38712,
									UpdatedDate:        "2022-08-31",
								},
							},
							{
								ID:          "00200521-20181022093938",
								DataSetName: "testCCC",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 159,
								},
								TableInf: core.TableInf{
									ID: "0000030001",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "昭和55年国勢調査 第1次基本集計 全国編",
									Title: core.Title{
										Number: "00101",
										Name:   "男女の別（性別）（３），年齢５歳階級（２３），人口 全国・市部・郡部・都道府県（４７），全域・人口集中地区の別",
									},
									Cycle:       "-",
									SurveyDate:  "198010",
									OpenDate:    "2007-10-05",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 3651,
									UpdatedDate:        "2021-06-25",
								},
							},
							{
								ID:          "CTCdemo-kokusei2",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 16779,
								},
								TableInf: core.TableInf{
									ID: "0003038587",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成22年国勢調査 人口等基本集計（男女・年齢・配偶関係，世帯の構成，住居の状態など）",
									Title: core.Title{
										Number: "00200",
										Name:   "男女別人口及び世帯の種類(２区分)別世帯数 全国，市部・郡部，都道府県，市部・郡部，支庁，郡計，市区町村・旧市町村，全域・人口集中地区",
									},
									Cycle:       "-",
									SurveyDate:  "201010",
									OpenDate:    "2011-10-26",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 39592,
									UpdatedDate:        "2021-06-25",
								},
							},
							{
								ID:          "okayama-jinse",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 4,
								},
								TableInf: core.TableInf{
									ID: "0003038587",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成22年国勢調査 人口等基本集計（男女・年齢・配偶関係，世帯の構成，住居の状態など）",
									Title: core.Title{
										Number: "00200",
										Name:   "男女別人口及び世帯の種類(２区分)別世帯数 全国，市部・郡部，都道府県，市部・郡部，支庁，郡計，市区町村・旧市町村，全域・人口集中地区",
									},
									Cycle:       "-",
									SurveyDate:  "201010",
									OpenDate:    "2011-10-26",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 39592,
									UpdatedDate:        "2021-06-25",
								},
							},
							{
								ID:          "yokohama-jinse",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 4,
								},
								TableInf: core.TableInf{
									ID: "0003038587",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成22年国勢調査 人口等基本集計（男女・年齢・配偶関係，世帯の構成，住居の状態など）",
									Title: core.Title{
										Number: "00200",
										Name:   "男女別人口及び世帯の種類(２区分)別世帯数 全国，市部・郡部，都道府県，市部・郡部，支庁，郡計，市区町村・旧市町村，全域・人口集中地区",
									},
									Cycle:       "-",
									SurveyDate:  "201010",
									OpenDate:    "2011-10-26",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 39592,
									UpdatedDate:        "2021-06-25",
								},
							},
							{
								ID:          "00200521-20141212155112",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 114507,
								},
								TableInf: core.TableInf{
									ID: "0003064691",
									StatName: core.StatName{
										Code: "00200521",
										Name: "国勢調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成22年国勢調査 移動人口の産業等集計(移動人口の労働力状態，産業(大分類)，教育)",
									Title: core.Title{
										Number: "00402",
										Name:   "現住都道府県による5年前の常住地，在学か否かの別・最終卒業学校の種類(6区分)，男女別人口(転入) 都道府県",
									},
									Cycle:       "-",
									SurveyDate:  "201010",
									OpenDate:    "2012-07-31",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "02",
										Name: "人口・世帯",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "人口",
									},
									OverallTotalNumber: 229014,
									UpdatedDate:        "2022-08-31",
								},
							},
							{
								ID:          "00200522-20150107180406",
								DataSetName: "住宅・土地統計調査\u3000データセット１",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 2,
								},
								TableInf: core.TableInf{
									ID: "0003010900",
									StatName: core.StatName{
										Code: "00200522",
										Name: "住宅・土地統計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成20年住宅・土地統計調査 全国編",
									Title: core.Title{
										Number: "001-2",
										Name:   "居住世帯の有無(9区分)別住宅数及び建物の種類(4区分)別住宅以外で人が居住する建物数―全国，人口集中地区",
									},
									Cycle:       "-",
									SurveyDate:  "200810",
									OpenDate:    "2010-03-30",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "08",
										Name: "住宅・土地・建設",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "住宅・土地",
									},
									OverallTotalNumber: 250,
									UpdatedDate:        "2020-01-31",
								},
							},
							{
								ID:          "CTCdemo-jutaku1",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 31632,
								},
								TableInf: core.TableInf{
									ID: "0003009791",
									StatName: core.StatName{
										Code: "00200522",
										Name: "住宅・土地統計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成20年住宅・土地統計調査 都道府県編",
									Title: core.Title{
										Number: "015-1",
										Name:   "住宅の種類(2区分)，住宅の所有の関係(5区分)，建て方(4区分)・建築の時期(6区分)別住宅数，世帯数，世帯人員，１住宅当たり居住室数，１住宅当たり居住室の畳数，１住宅当たり延べ面積，１人当たり居住室の畳数及び１室当たり人員―市区町村",
									},
									Cycle:       "-",
									SurveyDate:  "200810",
									OpenDate:    "2011-01-06",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "08",
										Name: "住宅・土地・建設",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "住宅・土地",
									},
									OverallTotalNumber: 769712,
									UpdatedDate:        "2020-01-31",
								},
							},
							{
								ID:          "00200543-20160708172751",
								DataSetName: "平成27年[大学] 組織，大学等の種類，専門別研究本務者数（大学等）",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 640,
								},
								TableInf: core.TableInf{
									ID: "0003117584",
									StatName: core.StatName{
										Code: "00200543",
										Name: "科学技術研究調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成27年科学技術研究調査",
									Title: core.Title{
										Number: "40201",
										Name:   "[大学] 組織，大学等の種類，専門別研究本務者数（大学等）",
									},
									Cycle:       "",
									SurveyDate:  "-",
									OpenDate:    "",
									SmallArea:   0,
									CollectArea: "",
									MainCategory: core.MainCategory{
										Code: "",
										Name: "",
									},
									SubCategory: core.SubCategory{
										Code: "",
										Name: "",
									},
									OverallTotalNumber: 0,
									UpdatedDate:        "",
								},
							},
							{
								ID:          "CTCdemo-keizai1",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 8058,
								},
								TableInf: core.TableInf{
									ID: "0003032616",
									StatName: core.StatName{
										Code: "00200552",
										Name: "経済センサス－基礎調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "平成21年経済センサス-基礎調査 事業所に関する集計",
									Title: core.Title{
										Number: "04700",
										Name:   "産業（大分類），資本金階級（10区分），単独・本所（２区分），存続・新設・廃業別民営事業所数及び男女別従業者数（外国の会社を除く会社の単独及び本所事業所）－全国，都道府県，市区町村",
									},
									Cycle:       "-",
									SurveyDate:  "200907",
									OpenDate:    "2011-06-03",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "07",
										Name: "企業・家計・経済",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "企業活動",
									},
									OverallTotalNumber: 8434136,
									UpdatedDate:        "2021-03-31",
								},
							},
							{
								ID:          "CTCdemo-kakei1",
								DataSetName: "家計調査デモ用",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 454080,
								},
								TableInf: core.TableInf{
									ID: "0003013276",
									StatName: core.StatName{
										Code: "00200561",
										Name: "家計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "家計調査 家計収支編 二人以上の世帯",
									Title: core.Title{
										Number: "010",
										Name:   "品目分類 品目分類（平成22年改定）（総数：金額）",
									},
									Cycle:       "月次",
									SurveyDate:  "0",
									OpenDate:    "2015-02-06",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "07",
										Name: "企業・家計・経済",
									},
									SubCategory: core.SubCategory{
										Code: "04",
										Name: "家計",
									},
									OverallTotalNumber: 3819928,
									UpdatedDate:        "2019-10-08",
								},
							},
							{
								ID:          "CTCdemo-kakei2011",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 423072,
								},
								TableInf: core.TableInf{
									ID: "0003013276",
									StatName: core.StatName{
										Code: "00200561",
										Name: "家計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "家計調査 家計収支編 二人以上の世帯",
									Title: core.Title{
										Number: "010",
										Name:   "品目分類 品目分類（平成22年改定）（総数：金額）",
									},
									Cycle:       "月次",
									SurveyDate:  "0",
									OpenDate:    "2015-02-06",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "07",
										Name: "企業・家計・経済",
									},
									SubCategory: core.SubCategory{
										Code: "04",
										Name: "家計",
									},
									OverallTotalNumber: 3819928,
									UpdatedDate:        "2019-10-08",
								},
							},
							{
								ID:          "CTCdemo-kouri1",
								DataSetName: "小売調査デモ用",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 601506,
								},
								TableInf: core.TableInf{
									ID: "0003013703",
									StatName: core.StatName{
										Code: "00200571",
										Name: "小売物価統計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "小売物価統計調査",
									Title: core.Title{
										Number: "",
										Name:   "主要品目の都市別小売価格－県庁所在市及び人口15万以上の市",
									},
									Cycle:       "",
									SurveyDate:  "-",
									OpenDate:    "",
									SmallArea:   0,
									CollectArea: "",
									MainCategory: core.MainCategory{
										Code: "",
										Name: "",
									},
									SubCategory: core.SubCategory{
										Code: "",
										Name: "",
									},
									OverallTotalNumber: 0,
									UpdatedDate:        "",
								},
							},
							{
								ID:          "okayama-gas",
								DataSetName: "",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 1,
								},
								TableInf: core.TableInf{
									ID: "0003013703",
									StatName: core.StatName{
										Code: "00200571",
										Name: "小売物価統計調査",
									},
									GovOrg: core.GovOrg{
										Code: "00200",
										Name: "総務省",
									},
									StatisticsName: "小売物価統計調査",
									Title: core.Title{
										Number: "",
										Name:   "主要品目の都市別小売価格－県庁所在市及び人口15万以上の市",
									},
									Cycle:       "",
									SurveyDate:  "-",
									OpenDate:    "",
									SmallArea:   0,
									CollectArea: "",
									MainCategory: core.MainCategory{
										Code: "",
										Name: "",
									},
									SubCategory: core.SubCategory{
										Code: "",
										Name: "",
									},
									OverallTotalNumber: 0,
									UpdatedDate:        "",
								},
							},
							{
								ID:          "00350300-20220625203025",
								DataSetName: "beefpork",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 999,
								},
								TableInf: core.TableInf{
									ID: "0003425296",
									StatName: core.StatName{
										Code: "00350300",
										Name: "普通貿易統計",
									},
									GovOrg: core.GovOrg{
										Code: "00350",
										Name: "財務省",
									},
									StatisticsName: "貿易統計_全国分 概況品別国別表 輸入",
									Title: core.Title{
										Number: "",
										Name:   "確速 概況品別国別表 (輸入 1-12月：確々報)2021年 ,  (輸入 1-3月：確報 , 4月：輸入9桁速報)2022年",
									},
									Cycle:       "年次",
									SurveyDate:  "202101-202112",
									OpenDate:    "2022-10-28",
									SmallArea:   0,
									CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "16",
										Name: "国際",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "貿易・国際収支",
									},
									OverallTotalNumber: 949364,
									UpdatedDate:        "2022-10-28",
								},
							},
							{
								ID:          "00350300-20220625205200",
								DataSetName: "beefandpork",
								PublicState: "YES",
								Result: core.ResponseRefDatasetResultInf{
									TotalNumber: 1917,
								},
								TableInf: core.TableInf{
									ID: "0003425296",
									StatName: core.StatName{
										Code: "00350300",
										Name: "普通貿易統計",
									},
									GovOrg: core.GovOrg{
										Code: "00350",
										Name: "財務省",
									}, StatisticsName: "貿易統計_全国分 概況品別国別表 輸入",
									Title: core.Title{
										Number: "",
										Name:   "確速 概況品別国別表 (輸入 1-12月：確々報)2021年 ,  (輸入 1-3月：確報 , 4月：輸入9桁速報)2022年",
									},
									Cycle:      "年次",
									SurveyDate: "202101-202112",
									OpenDate:   "2022-10-28",
									SmallArea:  0, CollectArea: "該当なし",
									MainCategory: core.MainCategory{
										Code: "16",
										Name: "国際",
									},
									SubCategory: core.SubCategory{
										Code: "01",
										Name: "貿易・国際収支",
									},
									OverallTotalNumber: 949364, UpdatedDate: "2022-10-28",
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
			got, _ := ac.GetDatasetList(ctx, tt.arg)
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetDatasetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
