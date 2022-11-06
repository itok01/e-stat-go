package core

import (
	"context"
	"encoding/xml"
)

type ParamsGetDataCatalog struct {
	SurveyYears       string `url:"surveyYears,omitempty" xml:"SURVEY_YEARS,omitempty"`
	OpenYears         string `url:"openYears,omitempty" xml:"OPEN_YEARS,omitempty"`
	StatsField        int    `url:"statsField,omitempty" xml:"STATS_FIELD,omitempty"`
	StatsCode         int    `url:"statsCode,omitempty" xml:"STATS_CODE,omitempty"`
	SearchWord        string `url:"searchWord,omitempty" xml:"SEARCH_WORD,omitempty"`
	CollectArea       int    `url:"collectArea,omitempty" xml:"COLLECT_AREA,omitempty"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG,omitempty"`
	DataType          string `url:"dataType,omitempty" xml:"DATA_TYPE"`
	StartPosition     int    `url:"startPosition,omitempty" xml:"START_POSITION,omitempty"`
	CatalogId         int    `url:"catalogId,omitempty" xml:"CATALOG_ID"`
	ResourceId        int    `url:"resourceId,omitempty" xml:"RESOURCE_ID"`
	Limit             int    `url:"limit,omitempty" xml:"LIMIT,omitempty"`
	UpdatedDate       string `url:"updatedDate,omitempty" xml:"UPDATED_DATE,omitempty"`
}

type ParamsGetDataCatalogRoot struct {
	CommonParams
	ParamsGetDataCatalog
}

type ResponseGetDataCatalogParameter struct {
	CommonParams
	ParamsGetDataCatalog
}

type DatasetTitle struct {
	Name string `xml:"NAME"`
	StatisticsNameSpec
	Cycle       string `xml:"CYCLE"`
	SurveyDate  string `xml:"SURVEY_DATE"`
	CollectArea string `xml:"COLLECT_AREA"`
}

type Dataset struct {
	StatName          StatName     `xml:"STAT_NAME"`
	Organization      Organization `xml:"ORGANIZATION"`
	Title             DatasetTitle `xml:"TITLE"`
	Description       Description  `xml:"DESCRIPTION"`
	Publisher         string       `xml:"PUBLISHER"`
	ContactPoint      string       `xml:"CONTACT_POINT"`
	Creator           string       `xml:"CREATOR"`
	ReleaseDate       string       `xml:"RELEASE_DATE"`
	LastModifiedDate  string       `xml:"LAST_MODIFIED_DATE"`
	FrequencyOfUpdate string       `xml:"FREQUENCY_OF_UPDATE"`
	LandingPage       string       `xml:"LANDING_PAGE"`
}

type ResourceTitle struct {
	Name        string `xml:"NAME"`
	TableNumber int    `xml:"TABLE_NO"`
	TitleSpec
}

type Resource struct {
	ID    string        `xml:"id,attr"`
	Title ResourceTitle `xml:"TITLE"`
	URL   string        `xml:"URL"`
	Description
	Format            string `xml:"FORMAT"`
	ReleaseDate       string `xml:"RELEASE_DATE"`
	LastModifiedDate  string `xml:"LAST_MODIFIED_DATE"`
	ResourceLicenceID string `xml:"RESOURCE_LICENCE_ID"`
	Language          string `xml:"LANGUAGE"`
}

type Resources struct {
	Resource Resource `xml:"RESOURCE"`
}

type DataCatalogInf struct {
	ID        string      `xml:"id,attr"`
	Dataset   Dataset     `xml:"DATASET"`
	Resources []Resources `xml:"RESOURCES"`
}

type DataCatalogListInf struct {
	Number      int              `xml:"NUMBER"`
	Result      ResultInf        `xml:"RESULT_INF"`
	DataCatalog []DataCatalogInf `xml:"DATA_CATALOG_INF"`
}

type ResponseGetDataCatalog struct {
	Result          ResponseResult                  `xml:"RESULT"`
	Parameter       ResponseGetDataCatalogParameter `xml:"PARAMETER"`
	DataCatalogList DataCatalogListInf              `xml:"DATA_CATALOG_LIST_INF,omitempty"`
}

type ResponseGetDataCatalogRoot struct {
	ResponseGetDataCatalog `xml:"GET_DATA_CATALOG"`
}

// 統計表情報取得
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_1
func (c *ApiClient) GetDataCatalog(ctx context.Context, params ParamsGetDataCatalog) (*ResponseGetDataCatalogRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/getDataCatalog", ParamsGetDataCatalogRoot{
		CommonParams:         c.CommonParams,
		ParamsGetDataCatalog: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseGetDataCatalogRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
