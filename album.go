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
	"github.com/cjavdev/spotted-go/packages/pagination"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
	"github.com/cjavdev/spotted-go/shared/constant"
)

// AlbumService contains methods and other services that help with interacting with
// the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlbumService] method instead.
type AlbumService struct {
	Options []option.RequestOption
}

// NewAlbumService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlbumService(opts ...option.RequestOption) (r AlbumService) {
	r = AlbumService{}
	r.Options = opts
	return
}

// Get Spotify catalog information for a single album.
func (r *AlbumService) Get(ctx context.Context, id string, query AlbumGetParams, opts ...option.RequestOption) (res *AlbumGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("albums/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Spotify catalog information for multiple albums identified by their Spotify
// IDs.
func (r *AlbumService) BulkGet(ctx context.Context, query AlbumBulkGetParams, opts ...option.RequestOption) (res *AlbumBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "albums"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Spotify catalog information about an album’s tracks. Optional parameters can
// be used to limit the number of tracks returned.
func (r *AlbumService) ListTracks(ctx context.Context, id string, query AlbumListTracksParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.SimplifiedTrackObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("albums/%s/tracks", id)
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

// Get Spotify catalog information about an album’s tracks. Optional parameters can
// be used to limit the number of tracks returned.
func (r *AlbumService) ListTracksAutoPaging(ctx context.Context, id string, query AlbumListTracksParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.SimplifiedTrackObject] {
	return pagination.NewCursorURLPageAutoPager(r.ListTracks(ctx, id, query, opts...))
}

type AlbumGetResponse struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	ID string `json:"id,required"`
	// The type of the album.
	//
	// Any of "album", "single", "compilation".
	AlbumType AlbumGetResponseAlbumType `json:"album_type,required"`
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
	ReleaseDatePrecision AlbumGetResponseReleaseDatePrecision `json:"release_date_precision,required"`
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
	Tracks AlbumGetResponseTracks `json:"tracks"`
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
func (r AlbumGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the album.
type AlbumGetResponseAlbumType string

const (
	AlbumGetResponseAlbumTypeAlbum       AlbumGetResponseAlbumType = "album"
	AlbumGetResponseAlbumTypeSingle      AlbumGetResponseAlbumType = "single"
	AlbumGetResponseAlbumTypeCompilation AlbumGetResponseAlbumType = "compilation"
)

// The precision with which `release_date` value is known.
type AlbumGetResponseReleaseDatePrecision string

const (
	AlbumGetResponseReleaseDatePrecisionYear  AlbumGetResponseReleaseDatePrecision = "year"
	AlbumGetResponseReleaseDatePrecisionMonth AlbumGetResponseReleaseDatePrecision = "month"
	AlbumGetResponseReleaseDatePrecisionDay   AlbumGetResponseReleaseDatePrecision = "day"
)

// The tracks of the album.
type AlbumGetResponseTracks struct {
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
func (r AlbumGetResponseTracks) RawJSON() string { return r.JSON.raw }
func (r *AlbumGetResponseTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumBulkGetResponse struct {
	Albums []AlbumBulkGetResponseAlbum `json:"albums,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Albums      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AlbumBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumBulkGetResponseAlbum struct {
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
	Tracks AlbumBulkGetResponseAlbumTracks `json:"tracks"`
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
func (r AlbumBulkGetResponseAlbum) RawJSON() string { return r.JSON.raw }
func (r *AlbumBulkGetResponseAlbum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The tracks of the album.
type AlbumBulkGetResponseAlbumTracks struct {
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
func (r AlbumBulkGetResponseAlbumTracks) RawJSON() string { return r.JSON.raw }
func (r *AlbumBulkGetResponseAlbumTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlbumGetParams struct {
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

// URLQuery serializes [AlbumGetParams]'s query parameters as `url.Values`.
func (r AlbumGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlbumBulkGetParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the albums.
	// Maximum: 20 IDs.
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

// URLQuery serializes [AlbumBulkGetParams]'s query parameters as `url.Values`.
func (r AlbumBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlbumListTracksParams struct {
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

// URLQuery serializes [AlbumListTracksParams]'s query parameters as `url.Values`.
func (r AlbumListTracksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
