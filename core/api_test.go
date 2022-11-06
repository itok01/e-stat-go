package core_test

import (
	"context"
	_ "embed"
	"net/http"
	"reflect"

	"github.com/itok01/e-stat-go/core"
)

type mockHttpClient struct{}

//go:embed testmock/get_stats_list.xml
var responseGetStatsList []byte

//go:embed testmock/get_meta_info.xml
var responseGetMetaInfoList []byte

//go:embed testmock/get_stat_data.xml
var responseGetStatsData []byte

//go:embed testmock/post_dataset.xml
var responsePostDataset []byte

//go:embed testmock/ref_dataset.xml
var responseRefDataset []byte

//go:embed testmock/get_dataset_list.xml
var responseGetDatasetList []byte

//go:embed testmock/get_data_catalog.xml
var responseGetDataCatalog []byte

func (hc *mockHttpClient) Get(ctx context.Context, path string, query any) (int, []byte, error) {
	resps := map[string]map[reflect.Type][]byte{
		"/getStatsList": {
			reflect.TypeOf(core.ParamsGetStatsListRoot{}): responseGetStatsList,
		},
		"/getMetaInfo": {
			reflect.TypeOf(core.ParamsGetMetaInfoListRoot{}): responseGetMetaInfoList,
		},
		"/getStatsData": {
			reflect.TypeOf(core.ParamsGetStatsDataRoot{}): responseGetStatsData,
		},
		"/refDataset": {
			reflect.TypeOf(core.ParamsRefDatasetRoot{}): responseRefDataset,
			reflect.TypeOf(core.ParamsGetDatasetListRoot{}): responseGetDatasetList,
		},
		"/getDataCatalog": {
			reflect.TypeOf(core.ParamsGetDataCatalogRoot{}): responseGetDataCatalog,
		},
	}

	if body, ok := resps[path][reflect.ValueOf(query).Type()]; ok {
		return http.StatusOK, body, nil
	}
	
	return http.StatusNotFound, nil, nil
}

func (hc *mockHttpClient) Post(ctx context.Context, path string, data any) (int, []byte, error) {
	resps := map[string]map[reflect.Type][]byte{
		"/postDataset": {
			reflect.TypeOf(core.ParamsPostDatasetRoot{}): responsePostDataset,
		},
	}

	if body, ok := resps[path][reflect.ValueOf(data).Type()]; ok {
		return http.StatusOK, body, nil
	}

	return http.StatusNotFound, nil, nil
}

func (hc *mockHttpClient) PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error) {
	return http.StatusNotFound, nil, nil
}
