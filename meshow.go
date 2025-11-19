// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/pagination"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// MeShowService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeShowService] method instead.
type MeShowService struct {
	Options []option.RequestOption
}

// NewMeShowService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeShowService(opts ...option.RequestOption) (r MeShowService) {
	r = MeShowService{}
	r.Options = opts
	return
}

// Get a list of shows saved in the current Spotify user's library. Optional
// parameters can be used to limit the number of shows returned.
func (r *MeShowService) List(ctx context.Context, query MeShowListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MeShowListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/shows"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of shows saved in the current Spotify user's library. Optional
// parameters can be used to limit the number of shows returned.
func (r *MeShowService) ListAutoPaging(ctx context.Context, query MeShowListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MeShowListResponse] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

// Check if one or more shows is already saved in the current Spotify user's
// library.
func (r *MeShowService) Check(ctx context.Context, query MeShowCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/shows/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete one or more shows from current Spotify user's library.
func (r *MeShowService) Remove(ctx context.Context, body MeShowRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/shows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Save one or more shows to current Spotify user's library.
func (r *MeShowService) Save(ctx context.Context, body MeShowSaveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/shows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type MeShowListResponse struct {
	// The date and time the show was saved. Timestamps are returned in ISO 8601 format
	// as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If
	// the time is imprecise (for example, the date/time of an album release), an
	// additional field indicates the precision; see for example, release_date in an
	// album object.
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Information about the show.
	Show shared.ShowBase `json:"show"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		Show        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeShowListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeShowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeShowListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeShowListParams]'s query parameters as `url.Values`.
func (r MeShowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeShowCheckParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the shows.
	// Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeShowCheckParams]'s query parameters as `url.Values`.
func (r MeShowCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeShowRemoveParams struct {
	// A JSON array of the
	// [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).
	// A maximum of 50 items can be specified in one request. _Note: if the `ids`
	// parameter is present in the query string, any IDs listed here in the body will
	// be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeShowRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeShowRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeShowRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeShowSaveParams struct {
	// A JSON array of the
	// [Spotify IDs](https://developer.spotify.com/documentation/web-api/#spotify-uris-and-ids).
	// A maximum of 50 items can be specified in one request. _Note: if the `ids`
	// parameter is present in the query string, any IDs listed here in the body will
	// be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeShowSaveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeShowSaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeShowSaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
