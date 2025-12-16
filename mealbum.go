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
	"github.com/cjavdev/spotted-go/shared/constant"
)

// MeAlbumService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeAlbumService] method instead.
type MeAlbumService struct {
	Options []option.RequestOption
}

// NewMeAlbumService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeAlbumService(opts ...option.RequestOption) (r MeAlbumService) {
	r = MeAlbumService{}
	r.Options = opts
	return
}

// Get a list of the albums saved in the current Spotify user's 'Your Music'
// library.
func (r *MeAlbumService) List(ctx context.Context, query MeAlbumListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MeAlbumListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/albums"
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

// Get a list of the albums saved in the current Spotify user's 'Your Music'
// library.
func (r *MeAlbumService) ListAutoPaging(ctx context.Context, query MeAlbumListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MeAlbumListResponse] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

// Check if one or more albums is already saved in the current Spotify user's 'Your
// Music' library.
func (r *MeAlbumService) Check(ctx context.Context, query MeAlbumCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/albums/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove one or more albums from the current user's 'Your Music' library.
func (r *MeAlbumService) Remove(ctx context.Context, body MeAlbumRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/albums"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Save one or more albums to the current user's 'Your Music' library.
func (r *MeAlbumService) Save(ctx context.Context, body MeAlbumSaveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/albums"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type MeAlbumListResponse struct {
	// The date and time the album was saved Timestamps are returned in ISO 8601 format
	// as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If
	// the time is imprecise (for example, the date/time of an album release), an
	// additional field indicates the precision; see for example, release_date in an
	// album object.
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Information about the album.
	Album MeAlbumListResponseAlbum `json:"album"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		Album       respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeAlbumListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeAlbumListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the album.
type MeAlbumListResponseAlbum struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	ID string `json:"id,required"`
	// The type of the album.
	//
	// Any of "album", "single", "compilation".
	AlbumType string `json:"album_type,required"`
	// The markets in which the album is available:
	// [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// _**NOTE**: an album is considered available in a market when at least 1 of its
	// tracks is available in that market._
	AvailableMarkets []string `json:"available_markets,required"`
	// Known external URLs for this album.
	ExternalURLs shared.ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the album.
	Href string `json:"href,required"`
	// The cover art for the album in various sizes, widest first.
	Images []shared.ImageObject `json:"images,required"`
	// The name of the album. In case of an album takedown, the value may be an empty
	// string.
	Name string `json:"name,required"`
	// The date the album was first released.
	ReleaseDate string `json:"release_date,required"`
	// The precision with which `release_date` value is known.
	//
	// Any of "year", "month", "day".
	ReleaseDatePrecision string `json:"release_date_precision,required"`
	// The number of tracks in the album.
	TotalTracks int64 `json:"total_tracks,required"`
	// The object type.
	Type constant.Album `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	Uri string `json:"uri,required"`
	// The artists of the album. Each artist object includes a link in `href` to more
	// detailed information about the artist.
	Artists []shared.SimplifiedArtistObject `json:"artists"`
	// The copyright statements of the album.
	Copyrights []shared.CopyrightObject `json:"copyrights"`
	// Known external IDs for the album.
	ExternalIDs shared.ExternalIDObject `json:"external_ids"`
	// **Deprecated** The array is always empty.
	//
	// Deprecated: deprecated
	Genres []string `json:"genres"`
	// The label associated with the album.
	Label string `json:"label"`
	// The popularity of the album. The value will be between 0 and 100, with 100 being
	// the most popular.
	Popularity int64 `json:"popularity"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions shared.AlbumRestrictionObject `json:"restrictions"`
	// The tracks of the album.
	Tracks MeAlbumListResponseAlbumTracks `json:"tracks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlbumType            respjson.Field
		AvailableMarkets     respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		Images               respjson.Field
		Name                 respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		TotalTracks          respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		Artists              respjson.Field
		Copyrights           respjson.Field
		ExternalIDs          respjson.Field
		Genres               respjson.Field
		Label                respjson.Field
		Popularity           respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		Tracks               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeAlbumListResponseAlbum) RawJSON() string { return r.JSON.raw }
func (r *MeAlbumListResponseAlbum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The tracks of the album.
type MeAlbumListResponseAlbumTracks struct {
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
	Total int64                          `json:"total,required"`
	Items []shared.SimplifiedTrackObject `json:"items"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Limit       respjson.Field
		Next        respjson.Field
		Offset      respjson.Field
		Previous    respjson.Field
		Total       respjson.Field
		Items       respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeAlbumListResponseAlbumTracks) RawJSON() string { return r.JSON.raw }
func (r *MeAlbumListResponseAlbumTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeAlbumListParams struct {
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

// URLQuery serializes [MeAlbumListParams]'s query parameters as `url.Values`.
func (r MeAlbumListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAlbumCheckParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the albums.
	// Maximum: 20 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeAlbumCheckParams]'s query parameters as `url.Values`.
func (r MeAlbumCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAlbumRemoveParams struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published param.Opt[bool] `json:"published,omitzero"`
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `["4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M"]`<br/>A maximum of 50 items
	// can be specified in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeAlbumRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeAlbumRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeAlbumRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeAlbumSaveParams struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published param.Opt[bool] `json:"published,omitzero"`
	// A JSON array of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `["4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M"]`<br/>A maximum of 50 items
	// can be specified in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeAlbumSaveParams) MarshalJSON() (data []byte, err error) {
	type shadow MeAlbumSaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeAlbumSaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
