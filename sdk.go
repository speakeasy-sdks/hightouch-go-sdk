package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
	"strings"
	"time"
)

var ServerList = []string{
	"https://api.hightouch.com/api/v1",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK Documentation: Hightouch Public Rest API to access syncs, models, sources and destinations
type SDK struct {

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*SDK)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		_language:   "go",
		_sdkVersion: "0.2.3",
		_genVersion: "1.8.6",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		sdk._securityClient = sdk._defaultClient

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	return sdk
}

// CreateDestination - Create Destination
// Create a new destination
func (s *SDK) CreateDestination(ctx context.Context, request operations.CreateDestinationRequest) (*operations.CreateDestinationResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/destinations"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateDestination200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// CreateModel - Create Model
// Create a new model
func (s *SDK) CreateModel(ctx context.Context, request operations.CreateModelRequest) (*operations.CreateModelResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/models"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateModel200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// CreateSource - Create Source
// Create a new source
func (s *SDK) CreateSource(ctx context.Context, request operations.CreateSourceRequest) (*operations.CreateSourceResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/sources"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateSource200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// CreateSync - Create Sync
// Create a new sync
func (s *SDK) CreateSync(ctx context.Context, request operations.CreateSyncRequest) (*operations.CreateSyncResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/syncs"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateSyncResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateSync200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 409:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// GetDestination - Get Destination
// Retrieve a destination based on its Hightouch ID
func (s *SDK) GetDestination(ctx context.Context, request operations.GetDestinationRequest) (*operations.GetDestinationResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/destinations/{destinationId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Destination
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Destination = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

// GetModel - Get Model
// Retrieve models from model ID
func (s *SDK) GetModel(ctx context.Context, request operations.GetModelRequest) (*operations.GetModelResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/models/{modelId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Model
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Model = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

// GetSource - Get Source
// Retrieve source from source ID
func (s *SDK) GetSource(ctx context.Context, request operations.GetSourceRequest) (*operations.GetSourceResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/sources/{sourceId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Source
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Source = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// GetSync - Get Sync
// Retrieve sync from sync ID
func (s *SDK) GetSync(ctx context.Context, request operations.GetSyncRequest) (*operations.GetSyncResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/syncs/{syncId}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSyncResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Sync
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Sync = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	}

	return res, nil
}

// ListDestination - List Destinations
// List the destinations in the user's workspace
func (s *SDK) ListDestination(ctx context.Context, request operations.ListDestinationRequest) (*operations.ListDestinationResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/destinations"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListDestination200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListDestination200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// ListModel - List Models
// List all the models in the current workspace
func (s *SDK) ListModel(ctx context.Context, request operations.ListModelRequest) (*operations.ListModelResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/models"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListModel200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListModel200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// ListSource - List Sources
// List all the sources in the current workspace
func (s *SDK) ListSource(ctx context.Context, request operations.ListSourceRequest) (*operations.ListSourceResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/sources"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSource200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListSource200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	}

	return res, nil
}

// ListSync - List Syncs
// List all the syncs in the current workspace
func (s *SDK) ListSync(ctx context.Context, request operations.ListSyncRequest) (*operations.ListSyncResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/syncs"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSyncResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSync200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListSync200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// ListSyncRuns - List Sync Runs
// List all sync runs under a sync
func (s *SDK) ListSyncRuns(ctx context.Context, request operations.ListSyncRunsRequest) (*operations.ListSyncRunsResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/syncs/{syncId}/runs", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSyncRunsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSyncRuns200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListSyncRuns200ApplicationJSONObject = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// TriggerRun - Trigger Sync
// Trigger a new run for the given sync.
//
// If a run is already in progress, this queues a sync run that will get
// executed immediately after the current run completes.
func (s *SDK) TriggerRun(ctx context.Context, request operations.TriggerRunRequest) (*operations.TriggerRunResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/syncs/{syncId}/trigger", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.TriggerRunResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.TriggerRunOutput
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.TriggerRunOutput = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// TriggerRunCustom - Trigger Sync From ID or Slug
// Trigger a new run globally based on sync id or sync slug
//
// If a run is already in progress, this queues a sync run that will get
// executed immediately after the current run completes.
func (s *SDK) TriggerRunCustom(ctx context.Context, request operations.TriggerRunCustomRequest) (*operations.TriggerRunCustomResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/syncs/trigger"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.TriggerRunCustomResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.TriggerRunCustom200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 400:
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	}

	return res, nil
}

// UpdateDestination - Update Destination
// Update an existing destination
//
// Patch a destination based on its Hightouch ID
func (s *SDK) UpdateDestination(ctx context.Context, request operations.UpdateDestinationRequest) (*operations.UpdateDestinationResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/destinations/{destinationId}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateDestination200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// UpdateModel - Update Model
// Update an existing model
//
// Patch a model based on its Hightouch ID
func (s *SDK) UpdateModel(ctx context.Context, request operations.UpdateModelRequest) (*operations.UpdateModelResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/models/{modelId}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateModel200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// UpdateSource - Update Source
// Update an existing source
//
// Patch a source based on its Hightouch ID
func (s *SDK) UpdateSource(ctx context.Context, request operations.UpdateSourceRequest) (*operations.UpdateSourceResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/sources/{sourceId}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateSourceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateSource200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}

// UpdateSync - Update Sync
// Update an existing sync
//
// Patch a sync based on its Hightouch ID
func (s *SDK) UpdateSync(ctx context.Context, request operations.UpdateSyncRequest) (*operations.UpdateSyncResponse, error) {
	baseURL := s._serverURL
	url := utils.GenerateURL(ctx, baseURL, "/syncs/{syncId}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := utils.ConfigureSecurityClient(s._defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateSyncResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateSync200ApplicationJSONAnyOf = out
		}
	case httpRes.StatusCode == 401:
	case httpRes.StatusCode == 404:
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ValidateErrorJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ValidateErrorJSON = out
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InternalServerErrorEnum
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.InternalServerError = out
		}
	}

	return res, nil
}
