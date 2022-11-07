package core

import (
	"context"
	"encoding/xml"
)

type ParamsGetStatsList struct {
	SurveyYears       string `url:"surveyYears,omitempty" xml:"SURVEY_YEARS,omitempty"`
	OpenYears         string `url:"openYears,omitempty" xml:"OPEN_YEARS,omitempty"`
	StatsField        int    `url:"statsField,omitempty" xml:"STATS_FIELD,omitempty"`
	StatsCode         int    `url:"statsCode,omitempty" xml:"STATS_CODE,omitempty"`
	SearchWord        string `url:"searchWord,omitempty" xml:"SEARCH_WORD,omitempty"`
	SearchKind        int    `url:"searchKind,omitempty" xml:"SEARCH_KIND,omitempty"`
	CollectArea       int    `url:"collectArea,omitempty" xml:"COLLECT_AREA,omitempty"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG,omitempty"`
	StatsNameList     string `url:"statsNameList,omitempty" xml:"STATS_NAME_LIST,omitempty"`
	StartPosition     int    `url:"startPosition,omitempty" xml:"START_POSITION,omitempty"`
	Limit             int    `url:"limit,omitempty" xml:"LIMIT,omitempty"`
	UpdatedDate       string `url:"updatedDate,omitempty" xml:"UPDATED_DATE,omitempty"`
}

type ParamsGetStatsListRoot struct {
	CommonParams
	ParamsGetStatsList
}

type ResponseGetStatsListParameter struct {
	CommonParams
	ParamsGetStatsList
}

type ResponseGetStatsListDataList struct {
	Number int        `xml:"NUMBER"`
	Result ResultInf  `xml:"RESULT_INF"`
	Table  []TableInf `xml:"TABLE_INF"`
}

type ResponseGetStatsList struct {
	Result    ResponseResult                `xml:"RESULT"`
	Parameter ResponseGetStatsListParameter `xml:"PARAMETER"`
	DataList  ResponseGetStatsListDataList `xml:"DATALIST_INF,omitempty"`
}

type ResponseGetStatsListRoot struct {
	ResponseGetStatsList `xml:"GET_STATS_LIST"`
}

// 統計表情報取得
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_1
func (c *ApiClient) GetStatsList(ctx context.Context, params ParamsGetStatsList) (*ResponseGetStatsListRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/getStatsList", ParamsGetStatsListRoot{
		CommonParams:       c.CommonParams,
		ParamsGetStatsList: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseGetStatsListRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
