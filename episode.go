// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// EpisodeService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEpisodeService] method instead.
type EpisodeService struct {
	Options []option.RequestOption
}

// NewEpisodeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEpisodeService(opts ...option.RequestOption) (r EpisodeService) {
	r = EpisodeService{}
	r.Options = opts
	return
}

// Get Spotify catalog information for a single episode identified by its unique
// Spotify ID.
func (r *EpisodeService) Get(ctx context.Context, id string, query EpisodeGetParams, opts ...option.RequestOption) (res *shared.EpisodeObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("episodes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Spotify catalog information for several episodes based on their Spotify IDs.
func (r *EpisodeService) BulkGet(ctx context.Context, query EpisodeBulkGetParams, opts ...option.RequestOption) (res *EpisodeBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "episodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EpisodeBulkGetResponse struct {
	Episodes []shared.EpisodeObject `json:"episodes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Episodes    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EpisodeBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EpisodeBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EpisodeGetParams struct {
	// An
	// [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// If a country code is specified, only content that is available in that market
	// will be returned.<br/> If a valid user access token is specified in the request
	// header, the country associated with the user account will take priority over
	// this parameter.<br/> _**Note**: If neither market or user country are provided,
	// the content is considered unavailable for the client._<br/> Users can view the
	// country that is associated with their account in the
	// [account settings](https://www.spotify.com/account/overview/).
	Market param.Opt[string] `query:"market,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EpisodeGetParams]'s query parameters as `url.Values`.
func (r EpisodeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EpisodeBulkGetParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episodes. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	// An
	// [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// If a country code is specified, only content that is available in that market
	// will be returned.<br/> If a valid user access token is specified in the request
	// header, the country associated with the user account will take priority over
	// this parameter.<br/> _**Note**: If neither market or user country are provided,
	// the content is considered unavailable for the client._<br/> Users can view the
	// country that is associated with their account in the
	// [account settings](https://www.spotify.com/account/overview/).
	Market param.Opt[string] `query:"market,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EpisodeBulkGetParams]'s query parameters as `url.Values`.
func (r EpisodeBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
