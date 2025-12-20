package v2

import (
	"context"
	"net/url"
)

const (
	// Market & Token Data

	// TokenSecurity https://docs.mobula.io/rest-api-reference/endpoint/token-security-get
	TokenSecurity = "/api/2/token/security"
	// TokenDetails https://docs.mobula.io/rest-api-reference/endpoint/token-details
	TokenDetails = "/api/2/token/details"
	// AssetDetails https://docs.mobula.io/rest-api-reference/endpoint/asset-details
	AssetDetails = "/api/2/asset/details"
	// MarketDetails https://docs.mobula.io/rest-api-reference/endpoint/market-details
	MarketDetails = "/api/2/market/details"
	// TokenMarkets https://docs.mobula.io/rest-api-reference/endpoint/token-markets
	TokenMarkets = "/api/2/token/markets"
)

type HTTPClient interface {
	Get(ctx context.Context, path string, queryParams url.Values, result interface{}) error
}

func GetTokenSecurity(ctx context.Context, client HTTPClient, req *TokenSecurityRequest) (*TokenSecurityResponse, error) {
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("blockchain", req.Blockchain)

	var resp TokenSecurityResponse
	if err := client.Get(ctx, TokenSecurity, params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetTokenDetails(ctx context.Context, client HTTPClient, req *TokenDetailsRequest) (*TokenDetailsResponse, error) {
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("blockchain", req.Blockchain)

	var resp TokenDetailsResponse
	if err := client.Get(ctx, TokenDetails, params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetAssetDetails(ctx context.Context, client HTTPClient, req *AssetDetailsRequest) (*AssetDetailsResponse, error) {
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("blockchain", req.Blockchain)

	var resp AssetDetailsResponse
	if err := client.Get(ctx, AssetDetails, params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetMarketDetails(ctx context.Context, client HTTPClient, req *MarketDetailsRequest) (*MarketDetailsResponse, error) {
	params := url.Values{}
	params.Set("blockchain", req.Blockchain)
	params.Set("address", req.Address)

	var resp MarketDetailsResponse
	if err := client.Get(ctx, MarketDetails, params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetTokenMarkets(ctx context.Context, client HTTPClient, req *TokenMarketsRequest) (*TokenMarketsResponse, error) {
	params := url.Values{}
	params.Set("address", req.Address)
	params.Set("blockchain", req.Blockchain)
	params.Set("limit", req.Limit)

	var resp TokenMarketsResponse
	if err := client.Get(ctx, TokenMarkets, params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
