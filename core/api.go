package core

import "context"

type ApiClient struct {
	HttpClient   IHttpClient
	CommonParams CommonParams
}

type IApiClient interface {
	GetStatsList(ctx context.Context, params ParamsGetStatsList) (*ResponseGetStatsListRoot, error)
	GetMetaInfoList(ctx context.Context, params ParamsGetMetaInfoList) (*ResponseGetMetaInfoListRoot, error)
	GetStatsData(ctx context.Context, params ParamsGetStatsData) (*ResponseGetStatsDataRoot, error)
	PostDataset(ctx context.Context, params ParamsPostDataset) (*ResponsePostDatasetRoot, error)
	RefDataset(ctx context.Context, params ParamsRefDataset) (*ResponseRefDatasetRoot, error)
	GetDatasetList(ctx context.Context, params ParamsGetDatasetList) (*ResponseGetDatasetListRoot, error)
	GetDataCatalog(ctx context.Context, params ParamsGetDataCatalog) (*ResponseGetDataCatalogRoot, error)
	GetStatsDatas(ctx context.Context, params ParamsGetStatsDatas, statsDatasSpec []StatsDatasSpec) (*ResponseGetStatsData, error)
}

func NewApiClient(
	httpClient IHttpClient,
	commonParams CommonParams,
) IApiClient {
	return &ApiClient{
		HttpClient:   httpClient,
		CommonParams: commonParams,
	}
}
