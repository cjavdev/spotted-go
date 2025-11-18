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

// MeEpisodeService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeEpisodeService] method instead.
type MeEpisodeService struct {
	Options []option.RequestOption
}

// NewMeEpisodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeEpisodeService(opts ...option.RequestOption) (r MeEpisodeService) {
	r = MeEpisodeService{}
	r.Options = opts
	return
}

// Get a list of the episodes saved in the current Spotify user's library.<br/>
// This API endpoint is in **beta** and could change without warning. Please share
// any feedback that you have, or issues that you discover, in our
// [developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer).
func (r *MeEpisodeService) List(ctx context.Context, query MeEpisodeListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MeEpisodeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/episodes"
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

// Get a list of the episodes saved in the current Spotify user's library.<br/>
// This API endpoint is in **beta** and could change without warning. Please share
// any feedback that you have, or issues that you discover, in our
// [developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer).
func (r *MeEpisodeService) ListAutoPaging(ctx context.Context, query MeEpisodeListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MeEpisodeListResponse] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

// Check if one or more episodes is already saved in the current Spotify user's
// 'Your Episodes' library.<br/> This API endpoint is in **beta** and could change
// without warning. Please share any feedback that you have, or issues that you
// discover, in our
// [developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)..
func (r *MeEpisodeService) Check(ctx context.Context, query MeEpisodeCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/episodes/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove one or more episodes from the current user's library.<br/> This API
// endpoint is in **beta** and could change without warning. Please share any
// feedback that you have, or issues that you discover, in our
// [developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer).
func (r *MeEpisodeService) Remove(ctx context.Context, body MeEpisodeRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/episodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Save one or more episodes to the current user's library.<br/> This API endpoint
// is in **beta** and could change without warning. Please share any feedback that
// you have, or issues that you discover, in our
// [developer community forum](https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer).
func (r *MeEpisodeService) Save(ctx context.Context, body MeEpisodeSaveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/episodes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type MeEpisodeListResponse struct {
	// The date and time the episode was saved. Timestamps are returned in ISO 8601
	// format as Coordinated Universal Time (UTC) with a zero offset:
	// YYYY-MM-DDTHH:MM:SSZ.
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Information about the episode.
	Episode shared.EpisodeObject `json:"episode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		Episode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeEpisodeListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeEpisodeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeEpisodeListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
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
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeEpisodeListParams]'s query parameters as `url.Values`.
func (r MeEpisodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeEpisodeCheckParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episodes. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeEpisodeCheckParams]'s query parameters as `url.Values`.
func (r MeEpisodeCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeEpisodeRemoveParams struct {
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). <br/>A maximum
	// of 50 items can be specified in one request. _**Note**: if the `ids` parameter
	// is present in the query string, any IDs listed here in the body will be
	// ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeEpisodeRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeEpisodeRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeEpisodeRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeEpisodeSaveParams struct {
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). <br/>A maximum
	// of 50 items can be specified in one request. _**Note**: if the `ids` parameter
	// is present in the query string, any IDs listed here in the body will be
	// ignored._
	IDs []string `json:"ids,omitzero,required"`
	paramObj
}

func (r MeEpisodeSaveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeEpisodeSaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeEpisodeSaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
