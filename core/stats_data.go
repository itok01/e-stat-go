package core

import (
	"context"
	"encoding/xml"
)

type ParamsGetStatsData struct {
	DataSetID         string `json:"dataSetId,omitempty" xml:"DATA_SET_ID"`
	StatsDataId       string `url:"statsDataId,omitempty" xml:"STATS_DATA_ID"`
	NarrowingConditon `xml:"NARROWING_COND"`
	StartPosition     int    `url:"startPosition,omitempty" xml:"START_POSITION"`
	Limit             int    `url:"limit,omitempty" xml:"LIMIT"`
	MetaGetFlg        string `url:"metaGetFlg,omitempty" xml:"METAGET_FLG"`
	CntGetFlg         string `url:"cntGetFlg,omitempty" xml:"CNT_GET_FLG"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG"`
	AnnotationGetFlg  string `url:"annotationGetFlg,omitempty" xml:"ANNOTATION_GET_FLG"`
	ReplaceSpChar     int    `url:"replaceSpChar,omitempty" xml:"REPLACE_SP_CHAR"`
}

type paramsGetStatsData struct {
	CommonParams
	ParamsGetStatsData
}

type ResponseGetStatsDataParameter struct {
	CommonParams
	ParamsGetStatsData
}

type ResponseGetStatsDataStatisticalData struct {
	Number int       `xml:"NUMBER"`
	Result ResultInf `xml:"RESULT_INF"`
	Table  TableInf  `xml:"TABLE_INF"`
	Class  ClassInf  `xml:"CLASS_INF"`
	Data   DataInf   `xml:"DATA_INF"`
}

type ResponseGetStatsData struct {
	Result    ResponseResult                      `xml:"RESULT"`
	Parameter ResponseGetStatsDataParameter       `xml:"PARAMETER"`
	DataList  ResponseGetStatsDataStatisticalData `xml:"STATISTICAL_DATA,omitempty"`
}

type ResponseGetStatsDataRoot struct {
	ResponseGetStatsData `xml:"GET_STATS_DATA"`
}

// 統計データ取得
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_3
func (c *ApiClient) GetStatsData(ctx context.Context, params ParamsGetStatsData) (*ResponseGetStatsDataRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/getStatsData", paramsGetStatsData{
		CommonParams:       c.CommonParams,
		ParamsGetStatsData: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseGetStatsDataRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type StatsDatasSpec struct {
	StatsDataId       string `url:"statsDataId,omitempty" json:"statsDataId,omitempty" xml:"STATS_DATA_ID"`
	NarrowingConditon `xml:"NARROWING_COND"`
	StartPosition     int `url:"startPosition,omitempty" json:"startPosition,omitempty" xml:"START_POSITION"`
	Limit             int `url:"limit,omitempty" json:"limit,omitempty" xml:"LIMIT"`
}

type ParamsGetStatsDatas struct {
	MetaGetFlg        string `url:"metaGetFlg,omitempty" json:"metaGetFlg,omitempty" xml:"META_GET_FLG"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" json:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG"`
	AnnotationGetFlg  string `url:"annotationGetFlg,omitempty" json:"annotationGetFlg,omitempty" xml:"ANNOTATION_GET_FLG"`
	ReplaceSpChar     int    `url:"replaceSpChar,omitempty" json:"replaceSpChar,omitempty" xml:"REPLACE_SP_CHAR"`
	SectionHeaderFlg  string `url:"sectionHeaderFlg,omitempty" json:"sectionHeaderFlg,omitempty" xml:"SECTION_HEADER_FLG"`
	DataSetID         string `url:"dataSetId,omitempty" json:"dataSetId,omitempty" xml:"DATA_SET_ID"`
	// StatsDatasSpec    []StatsDatasSpec `json:"statsDatasSpec" xml:"STATS_DATAS_SPEC"`
}

type paramsGetStatsDatas struct {
	CommonParams
	ParamsGetStatsDatas
}

type ResponseGetStatsParameter struct {
	RequestNumber string `xml:"requestNo,attr"`
	ParamsGetStatsDatas
}

type ResponseGetStatsParameterList struct {
	CommonParams
	Parameter []ResponseGetStatsParameter `xml:"PARAMETER"`
}

type ResponseGetStatsDataStatisticalDataListTableList struct {
	RequestNumber string     `xml:"requestNo,attr"`
	TableInf      []TableInf `xml:"TABLE_INF"`
}

type ResponseGetStatsDataStatisticalDataListClassList struct {
	RequestNumber string     `xml:"requestNo,attr"`
	ClassInf      []ClassInf `xml:"CLASS_INF"`
}

type ResponseGetStatsDataStatisticalDataListDataList struct {
	RequestNumber string    `xml:"requestNo,attr"`
	DataInf       []DataInf `xml:"DATA_INF"`
}

type ResponseGetStatsDataStatisticalDataList struct {
	ResultInf    ResultInf                                        `xml:"RESULT_INF"`
	TableInfList ResponseGetStatsDataStatisticalDataListTableList `xml:"TABLE_INF_LIST"`
	ClassInfList ResponseGetStatsDataStatisticalDataListClassList `xml:"CLASS_INF_LIST"`
	DataInfList  ResponseGetStatsDataStatisticalDataListDataList  `xml:"DATA_INF_LIST"`
}

type ResponseGetStatsDatas struct {
	Result              ResponseResult                          `xml:"RESULT"`
	ParameterList       ResponseGetStatsParameterList           `xml:"PARAMETER_LIST"`
	StatisticalDataList ResponseGetStatsDataStatisticalDataList `xml:"STATISTICAL_DATA_LIST,omitempty"`
}

type ResponseGetStatsDatasRoot struct {
	ResponseGetStatsData `xml:"GET_STATS_DATAS"`
}

// 統計データ一括取得
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_7
func (c *ApiClient) GetStatsDatas(ctx context.Context, params ParamsGetStatsDatas, statsDatasSpec []StatsDatasSpec) (*ResponseGetStatsData, error) {
	_, body, err := c.HttpClient.PostJsonWithQuery(ctx, "/getStatsDatas", &paramsGetStatsDatas{
		CommonParams:        c.CommonParams,
		ParamsGetStatsDatas: params,
	}, statsDatasSpec)
	if err != nil {
		return nil, err
	}

	var data *ResponseGetStatsData
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
