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

// ArtistService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArtistService] method instead.
type ArtistService struct {
	Options []option.RequestOption
}

// NewArtistService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewArtistService(opts ...option.RequestOption) (r ArtistService) {
	r = ArtistService{}
	r.Options = opts
	return
}

// Get Spotify catalog information for a single artist identified by their unique
// Spotify ID.
func (r *ArtistService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.ArtistObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("artists/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Spotify catalog information for several artists based on their Spotify IDs.
//
// Deprecated: deprecated
func (r *ArtistService) BulkGet(ctx context.Context, query ArtistBulkGetParams, opts ...option.RequestOption) (res *ArtistBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "artists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Spotify catalog information about an artist's albums.
func (r *ArtistService) ListAlbums(ctx context.Context, id string, query ArtistListAlbumsParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[ArtistListAlbumsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("artists/%s/albums", id)
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

// Get Spotify catalog information about an artist's albums.
func (r *ArtistService) ListAlbumsAutoPaging(ctx context.Context, id string, query ArtistListAlbumsParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[ArtistListAlbumsResponse] {
	return pagination.NewCursorURLPageAutoPager(r.ListAlbums(ctx, id, query, opts...))
}

// Get Spotify catalog information about artists similar to a given artist.
// Similarity is based on analysis of the Spotify community's listening history.
//
// Deprecated: deprecated
func (r *ArtistService) ListRelatedArtists(ctx context.Context, id string, opts ...option.RequestOption) (res *ArtistListRelatedArtistsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("artists/%s/related-artists", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Spotify catalog information about an artist's top tracks by country.
//
// Deprecated: deprecated
func (r *ArtistService) TopTracks(ctx context.Context, id string, query ArtistTopTracksParams, opts ...option.RequestOption) (res *ArtistTopTracksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("artists/%s/top-tracks", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ArtistBulkGetResponse struct {
	Artists []shared.ArtistObject `json:"artists,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artists     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtistBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ArtistBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArtistListAlbumsResponse struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	ID string `json:"id,required"`
	// This field describes the relationship between the artist and the album.
	//
	// Any of "album", "single", "compilation", "appears_on".
	//
	// Deprecated: deprecated
	AlbumGroup ArtistListAlbumsResponseAlbumGroup `json:"album_group,required"`
	// The type of the album.
	//
	// Any of "album", "single", "compilation".
	AlbumType ArtistListAlbumsResponseAlbumType `json:"album_type,required"`
	// The artists of the album. Each artist object includes a link in `href` to more
	// detailed information about the artist.
	Artists []shared.SimplifiedArtistObject `json:"artists,required"`
	// The markets in which the album is available:
	// [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// _**NOTE**: an album is considered available in a market when at least 1 of its
	// tracks is available in that market._
	//
	// Deprecated: deprecated
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
	ReleaseDatePrecision ArtistListAlbumsResponseReleaseDatePrecision `json:"release_date_precision,required"`
	// The number of tracks in the album.
	TotalTracks int64 `json:"total_tracks,required"`
	// The object type.
	Type constant.Album `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	Uri string `json:"uri,required"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions shared.AlbumRestrictionObject `json:"restrictions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlbumGroup           respjson.Field
		AlbumType            respjson.Field
		Artists              respjson.Field
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
		Published            respjson.Field
		Restrictions         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtistListAlbumsResponse) RawJSON() string { return r.JSON.raw }
func (r *ArtistListAlbumsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This field describes the relationship between the artist and the album.
type ArtistListAlbumsResponseAlbumGroup string

const (
	ArtistListAlbumsResponseAlbumGroupAlbum       ArtistListAlbumsResponseAlbumGroup = "album"
	ArtistListAlbumsResponseAlbumGroupSingle      ArtistListAlbumsResponseAlbumGroup = "single"
	ArtistListAlbumsResponseAlbumGroupCompilation ArtistListAlbumsResponseAlbumGroup = "compilation"
	ArtistListAlbumsResponseAlbumGroupAppearsOn   ArtistListAlbumsResponseAlbumGroup = "appears_on"
)

// The type of the album.
type ArtistListAlbumsResponseAlbumType string

const (
	ArtistListAlbumsResponseAlbumTypeAlbum       ArtistListAlbumsResponseAlbumType = "album"
	ArtistListAlbumsResponseAlbumTypeSingle      ArtistListAlbumsResponseAlbumType = "single"
	ArtistListAlbumsResponseAlbumTypeCompilation ArtistListAlbumsResponseAlbumType = "compilation"
)

// The precision with which `release_date` value is known.
type ArtistListAlbumsResponseReleaseDatePrecision string

const (
	ArtistListAlbumsResponseReleaseDatePrecisionYear  ArtistListAlbumsResponseReleaseDatePrecision = "year"
	ArtistListAlbumsResponseReleaseDatePrecisionMonth ArtistListAlbumsResponseReleaseDatePrecision = "month"
	ArtistListAlbumsResponseReleaseDatePrecisionDay   ArtistListAlbumsResponseReleaseDatePrecision = "day"
)

type ArtistListRelatedArtistsResponse struct {
	Artists []shared.ArtistObject `json:"artists,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artists     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtistListRelatedArtistsResponse) RawJSON() string { return r.JSON.raw }
func (r *ArtistListRelatedArtistsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArtistTopTracksResponse struct {
	Tracks []shared.TrackObject `json:"tracks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tracks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtistTopTracksResponse) RawJSON() string { return r.JSON.raw }
func (r *ArtistTopTracksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArtistBulkGetParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) for the artists.
	// Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [ArtistBulkGetParams]'s query parameters as `url.Values`.
func (r ArtistBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ArtistListAlbumsParams struct {
	// A comma-separated list of keywords that will be used to filter the response. If
	// not supplied, all album types will be returned. <br/> Valid values are:<br/>-
	// `album`<br/>- `single`<br/>- `appears_on`<br/>- `compilation`<br/>For example:
	// `include_groups=album,single`.
	IncludeGroups param.Opt[string] `query:"include_groups,omitzero" json:"-"`
	// The maximum number of items to return. Default: 5. Minimum: 1. Maximum: 10.
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

// URLQuery serializes [ArtistListAlbumsParams]'s query parameters as `url.Values`.
func (r ArtistListAlbumsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ArtistTopTracksParams struct {
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

// URLQuery serializes [ArtistTopTracksParams]'s query parameters as `url.Values`.
func (r ArtistTopTracksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
