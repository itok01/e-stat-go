package core

import (
	"context"
	"encoding/xml"
)

type ParamsPostDataset struct {
	DataSetID   string `url:"dataSetId,omitempty" xml:"DATA_SET_ID"`
	StatsDataId string `url:"statsDataId,omitempty" xml:"STATS_DATA_ID"`
	NarrowingConditon
	OpenSpecified string `url:"openSpecified,omitempty" xml:"OPEN_SPECIFIED"`
	ProcessMode   string `url:"processMode,omitempty" xml:"PROCESS_MODE"`
	DataSetName   string `url:"dataSetName,omitempty" xml:"DATASET_NAME"`
}

type ParamsPostDatasetRoot struct {
	CommonParams
	ParamsPostDataset
}

type ResponsePostDatasetParameter struct {
	CommonParams
	DataSetID         string            `url:"dataSetId,omitempty" xml:"DATA_SET_ID"`
	StatsDataId       string            `url:"statsDataId,omitempty" xml:"STATS_DATA_ID"`
	NarrowingConditon NarrowingConditon `xml:"NARROWING_COND"`
	OpenSpecified     string            `url:"openSpecified,omitempty" xml:"OPEN_SPECIFIED"`
	ProcessMode       string            `url:"processMode,omitempty" xml:"PROCESS_MODE"`
	DataSetName       string            `url:"dataSetName,omitempty" xml:"DATASET_NAME"`
}

type ResponsePostDatasetRegistInf struct {
	Mode        string `xml:"mode,attr"`
	DatasetId   string `xml:"DATASET_ID"`
	StatsDataId string `xml:"STATS_DATA_ID"`
	PublicState string `xml:"PUBLIC_STATE"`
	TotalNumber int    `xml:"TOTAL_NUMBER"`
}

type ResponsePostDataset struct {
	Result    ResponseResult               `xml:"RESULT"`
	Parameter ResponsePostDatasetParameter `xml:"PARAMETER"`
	RefistInf ResponsePostDatasetRegistInf `xml:"REGIST_INF,omitempty"`
}

type ResponsePostDatasetRoot struct {
	ResponsePostDataset `xml:"POST_DATASET"`
}

// データセット登録
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_4
func (c *ApiClient) PostDataset(ctx context.Context, params ParamsPostDataset) (*ResponsePostDatasetRoot, error) {
	_, body, err := c.HttpClient.Post(ctx, "/postDataset", ParamsPostDatasetRoot{
		CommonParams:      c.CommonParams,
		ParamsPostDataset: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponsePostDatasetRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type ParamsRefDataset struct {
	DataSetID         string `url:"dataSetId,omitempty" xml:"DATA_SET_ID"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG"`
}

type ParamsRefDatasetRoot struct {
	CommonParams
	ParamsRefDataset
}

type ResponseRefDatasetParameter struct {
	CommonParams
	ParamsRefDataset
}

type ResponseRefDatasetResultInf struct {
	TotalNumber int `xml:"TOTAL_NUMBER"`
}

type ResponseRefDatasetInf struct {
	ID          string                      `xml:"id,attr"`
	DataSetName string                      `xml:"DATASET_NAME"`
	PublicState string                      `xml:"PUBLIC_STATE"`
	Result      ResponseRefDatasetResultInf `xml:"RESULT_INF"`
	TableInf    TableInf                    `xml:"TABLE_INF"`
}

type ResponseRefDatasetListInf struct {
	Number  int                     `xml:"NUMBER"`
	Dataset []ResponseRefDatasetInf `xml:"DATASET_INF"`
}

type ResponseRefDataset struct {
	Result    ResponseResult              `xml:"RESULT"`
	Parameter ResponseRefDatasetParameter `xml:"PARAMETER"`
	Dataset   ResponseRefDatasetInf       `xml:"DATASET_INF"`
}

type ResponseRefDatasetRoot struct {
	ResponseRefDataset `xml:"REF_DATASET,omitempty"`
}

// データセット参照
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_5
func (c *ApiClient) RefDataset(ctx context.Context, params ParamsRefDataset) (*ResponseRefDatasetRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/refDataset", ParamsRefDatasetRoot{
		CommonParams:     c.CommonParams,
		ParamsRefDataset: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseRefDatasetRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

type ParamsGetDatasetList struct {
	CollectArea       string `url:"collectArea,omitempty" xml:"COLLECT_AREA"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG"`
}

type ParamsGetDatasetListRoot struct {
	CommonParams
	ParamsGetDatasetList
}

type ResponseGetDatasetListParameter struct {
	CommonParams
	ParamsGetDatasetList
}

type ResponseGetDatasetList struct {
	Result      ResponseResult              `xml:"RESULT"`
	Parameter   ResponseRefDatasetParameter `xml:"PARAMETER"`
	DatasetList ResponseRefDatasetListInf   `xml:"DATASET_LIST_INF,omitempty"`
}

type ResponseGetDatasetListRoot struct {
	ResponseGetDatasetList `xml:"GET_DATASET_LIST,omitempty"`
}

// データセット参照
//
// パラメータのデータセットIDを指定しない場合、以下のように複数のデータセット情報が一覧として出力されます。
// ルートタグは<GET_DATASET_LIST>になります。
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_5
func (c *ApiClient) GetDatasetList(ctx context.Context, params ParamsGetDatasetList) (*ResponseGetDatasetListRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/refDataset", ParamsGetDatasetListRoot{
		CommonParams:         c.CommonParams,
		ParamsGetDatasetList: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseGetDatasetListRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
