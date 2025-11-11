// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/spotted-go/internal/apijson"
	"github.com/stainless-sdks/spotted-go/internal/apiquery"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-go/packages/respjson"
	"github.com/stainless-sdks/spotted-go/shared"
)

// BrowseService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowseService] method instead.
type BrowseService struct {
	Options    []option.RequestOption
	Categories BrowseCategoryService
}

// NewBrowseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrowseService(opts ...option.RequestOption) (r BrowseService) {
	r = BrowseService{}
	r.Options = opts
	r.Categories = NewBrowseCategoryService(opts...)
	return
}

// Get a list of Spotify featured playlists (shown, for example, on a Spotify
// player's 'Browse' tab).
//
// Deprecated: deprecated
func (r *BrowseService) GetFeaturedPlaylists(ctx context.Context, query BrowseGetFeaturedPlaylistsParams, opts ...option.RequestOption) (res *BrowseGetFeaturedPlaylistsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browse/featured-playlists"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of new album releases featured in Spotify (shown, for example, on a
// Spotify player’s “Browse” tab).
func (r *BrowseService) GetNewReleases(ctx context.Context, query BrowseGetNewReleasesParams, opts ...option.RequestOption) (res *BrowseGetNewReleasesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browse/new-releases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BrowseGetFeaturedPlaylistsResponse struct {
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
func (r BrowseGetFeaturedPlaylistsResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowseGetFeaturedPlaylistsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseGetNewReleasesResponse struct {
	Albums BrowseGetNewReleasesResponseAlbums `json:"albums,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Albums      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseGetNewReleasesResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowseGetNewReleasesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseGetNewReleasesResponseAlbums struct {
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
	Total int64                                    `json:"total,required"`
	Items []BrowseGetNewReleasesResponseAlbumsItem `json:"items"`
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
func (r BrowseGetNewReleasesResponseAlbums) RawJSON() string { return r.JSON.raw }
func (r *BrowseGetNewReleasesResponseAlbums) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseGetNewReleasesResponseAlbumsItem struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	ID string `json:"id,required"`
	// The type of the album.
	//
	// Any of "album", "single", "compilation".
	AlbumType string `json:"album_type,required"`
	// The artists of the album. Each artist object includes a link in `href` to more
	// detailed information about the artist.
	Artists []shared.SimplifiedArtistObject `json:"artists,required"`
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
	//
	// Any of "album".
	Type string `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	Uri string `json:"uri,required"`
	// Included in the response when a content restriction is applied.
	Restrictions shared.AlbumRestrictionObject `json:"restrictions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
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
		Restrictions         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowseGetNewReleasesResponseAlbumsItem) RawJSON() string { return r.JSON.raw }
func (r *BrowseGetNewReleasesResponseAlbumsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowseGetFeaturedPlaylistsParams struct {
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

// URLQuery serializes [BrowseGetFeaturedPlaylistsParams]'s query parameters as
// `url.Values`.
func (r BrowseGetFeaturedPlaylistsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrowseGetNewReleasesParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowseGetNewReleasesParams]'s query parameters as
// `url.Values`.
func (r BrowseGetNewReleasesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
