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

// BrowseCategoryService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowseCategoryService] method instead.
type BrowseCategoryService struct {
	Options []option.RequestOption
}

// NewBrowseCategoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrowseCategoryService(opts ...option.RequestOption) (r BrowseCategoryService) {
	r = BrowseCategoryService{}
	r.Options = opts
	return
}

// Get a single category used to tag items in Spotify (on, for example, the Spotify
// player’s “Browse” tab).
func (r *BrowseCategoryService) Get(ctx context.Context, categoryID string, query BrowseCategoryGetParams, opts ...option.RequestOption) (res *BrowseCategoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("browse/categories/%s", categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of categories used to tag items in Spotify (on, for example, the
// Spotify player’s “Browse” tab).
func (r *BrowseCategoryService) List(ctx context.Context, query BrowseCategoryListParams, opts ...option.RequestOption) (res *BrowseCategoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browse/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of Spotify playlists tagged with a particular category.
//
// Deprecated: deprecated
func (r *BrowseCategoryService) GetPlaylists(ctx context.Context, categoryID string, query BrowseCategoryGetPlaylistsParams, opts ...option.RequestOption) (res *BrowseCategoryGetPlaylistsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return
	}
	path := fmt.Sprintf("browse/categories/%s/playlists", categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BrowseCategoryGetResponse struct {
	// The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of
	// the category.
	ID string `json:"id,required"`
	// A link to the Web API endpoint returning full details of the category.
	Href string `json:"href,required"`
	// The category icon, in various sizes.
	Icons []shared.ImageObject `json:"icons,required"`
	// The name of the category.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Href        respjson.Field
		Icons       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseCategoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowseCategoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseCategoryListResponse struct {
	Categories BrowseCategoryListResponseCategories `json:"categories,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseCategoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowseCategoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseCategoryListResponseCategories struct {
	// A link to the Web API endpoint returning the full result of the request
	Href string `json:"href,required"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int64 `json:"limit,required"`
	// URL to the next page of items. ( `null` if none)
	Next string `json:"next,required"`
	// The offset of the items returned (as set in the query or by default)
	Offset int64 `json:"offset,required"`
	// URL to the previous page of items. ( `null` if none)
	Previous string `json:"previous,required"`
	// The total number of items available to return.
	Total int64                                      `json:"total,required"`
	Items []BrowseCategoryListResponseCategoriesItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Limit       respjson.Field
		Next        respjson.Field
		Offset      respjson.Field
		Previous    respjson.Field
		Total       respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseCategoryListResponseCategories) RawJSON() string { return r.JSON.raw }
func (r *BrowseCategoryListResponseCategories) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseCategoryListResponseCategoriesItem struct {
	// The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of
	// the category.
	ID string `json:"id,required"`
	// A link to the Web API endpoint returning full details of the category.
	Href string `json:"href,required"`
	// The category icon, in various sizes.
	Icons []shared.ImageObject `json:"icons,required"`
	// The name of the category.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Href        respjson.Field
		Icons       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseCategoryListResponseCategoriesItem) RawJSON() string { return r.JSON.raw }
func (r *BrowseCategoryListResponseCategoriesItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseCategoryGetPlaylistsResponse struct {
	// The localized message of a playlist.
	Message   string                      `json:"message"`
	Playlists shared.PagingPlaylistObject `json:"playlists"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Playlists   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseCategoryGetPlaylistsResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowseCategoryGetPlaylistsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseCategoryGetParams struct {
	// The desired language, consisting of an
	// [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an
	// [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2),
	// joined by an underscore. For example: `es_MX`, meaning &quot;Spanish
	// (Mexico)&quot;. Provide this parameter if you want the category strings returned
	// in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the
	// specified language is not available, the category strings returned will be in
	// the Spotify default language (American English)._
	Locale param.Opt[string] `query:"locale,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowseCategoryGetParams]'s query parameters as
// `url.Values`.
func (r BrowseCategoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrowseCategoryListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The desired language, consisting of an
	// [ISO 639-1](http://en.wikipedia.org/wiki/ISO_639-1) language code and an
	// [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2),
	// joined by an underscore. For example: `es_MX`, meaning &quot;Spanish
	// (Mexico)&quot;. Provide this parameter if you want the category strings returned
	// in a particular language.<br/> _**Note**: if `locale` is not supplied, or if the
	// specified language is not available, the category strings returned will be in
	// the Spotify default language (American English)._
	Locale param.Opt[string] `query:"locale,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowseCategoryListParams]'s query parameters as
// `url.Values`.
func (r BrowseCategoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrowseCategoryGetPlaylistsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowseCategoryGetPlaylistsParams]'s query parameters as
// `url.Values`.
func (r BrowseCategoryGetPlaylistsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
