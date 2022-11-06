package core

import (
	"context"
	"encoding/xml"
)

type ParamsGetMetaInfoList struct {
	StatsDataId       string `url:"statsDataId" xml:"STATS_DATA_ID,omitempty"`
	ExplanationGetFlg string `url:"explanationGetFlg,omitempty" xml:"EXPLANATION_GET_FLG,omitempty"`
}

type ParamsGetMetaInfoListRoot struct {
	CommonParams
	ParamsGetMetaInfoList
}

type ResponseGetMetaInfoListParameter struct {
	CommonParams
	ParamsGetMetaInfoList
}

type ResponseGetMetaInfoMetaDataList struct {
	Number int      `xml:"NUMBER"`
	Table  TableInf `xml:"TABLE_INF"`
	Class  ClassInf `xml:"CLASS_INF"`
}

type ResponseGetMetaInfoList struct {
	Result    ResponseResult                   `xml:"RESULT"`
	Parameter ResponseGetMetaInfoListParameter `xml:"PARAMETER"`
	DataList  ResponseGetMetaInfoMetaDataList  `xml:"METADATA_INF,omitempty"`
}

type ResponseGetMetaInfoListRoot struct {
	ResponseGetMetaInfoList `xml:"GET_META_INFO"`
}

// メタ情報取得
//
// https://www.e-stat.go.jp/api/api-info/e-stat-manual3-0#api_2_2
func (c *ApiClient) GetMetaInfoList(ctx context.Context, params ParamsGetMetaInfoList) (*ResponseGetMetaInfoListRoot, error) {
	_, body, err := c.HttpClient.Get(ctx, "/getMetaInfo", ParamsGetMetaInfoListRoot{
		CommonParams:          c.CommonParams,
		ParamsGetMetaInfoList: params,
	})
	if err != nil {
		return nil, err
	}

	var data *ResponseGetMetaInfoListRoot
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
