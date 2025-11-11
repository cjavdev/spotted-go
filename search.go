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
	"github.com/stainless-sdks/spotted-go/shared/constant"
)

// SearchService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Get Spotify catalog information about albums, artists, playlists, tracks, shows,
// episodes or audiobooks that match a keyword string. Audiobooks are only
// available within the US, UK, Canada, Ireland, New Zealand and Australia markets.
func (r *SearchService) Query(ctx context.Context, query SearchQueryParams, opts ...option.RequestOption) (res *SearchQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchQueryResponse struct {
	Albums     SearchQueryResponseAlbums     `json:"albums"`
	Artists    SearchQueryResponseArtists    `json:"artists"`
	Audiobooks SearchQueryResponseAudiobooks `json:"audiobooks"`
	Episodes   SearchQueryResponseEpisodes   `json:"episodes"`
	Playlists  shared.PagingPlaylistObject   `json:"playlists"`
	Shows      SearchQueryResponseShows      `json:"shows"`
	Tracks     SearchQueryResponseTracks     `json:"tracks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Albums      respjson.Field
		Artists     respjson.Field
		Audiobooks  respjson.Field
		Episodes    respjson.Field
		Playlists   respjson.Field
		Shows       respjson.Field
		Tracks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseAlbums struct {
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
	Total int64                           `json:"total,required"`
	Items []SearchQueryResponseAlbumsItem `json:"items"`
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
func (r SearchQueryResponseAlbums) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseAlbums) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseAlbumsItem struct {
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
	Type constant.Album `json:"type,required"`
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
func (r SearchQueryResponseAlbumsItem) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseAlbumsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseArtists struct {
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
	Total int64                 `json:"total,required"`
	Items []shared.ArtistObject `json:"items"`
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
func (r SearchQueryResponseArtists) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseArtists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseAudiobooks struct {
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
	Total int64                  `json:"total,required"`
	Items []shared.AudiobookBase `json:"items"`
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
func (r SearchQueryResponseAudiobooks) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseAudiobooks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseEpisodes struct {
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
	Total int64                            `json:"total,required"`
	Items []shared.SimplifiedEpisodeObject `json:"items"`
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
func (r SearchQueryResponseEpisodes) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseEpisodes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseShows struct {
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
	Total int64             `json:"total,required"`
	Items []shared.ShowBase `json:"items"`
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
func (r SearchQueryResponseShows) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseShows) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryResponseTracks struct {
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
	Total int64                `json:"total,required"`
	Items []shared.TrackObject `json:"items"`
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
func (r SearchQueryResponseTracks) RawJSON() string { return r.JSON.raw }
func (r *SearchQueryResponseTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchQueryParams struct {
	// Your search query.
	//
	// You can narrow down your search using field filters. The available filters are
	// `album`, `artist`, `track`, `year`, `upc`, `tag:hipster`, `tag:new`, `isrc`, and
	// `genre`. Each field filter only applies to certain result types.
	//
	// The `artist` and `year` filters can be used while searching albums, artists and
	// tracks. You can filter on a single `year` or a range (e.g. 1955-1960).<br /> The
	// `album` filter can be used while searching albums and tracks.<br /> The `genre`
	// filter can be used while searching artists and tracks.<br /> The `isrc` and
	// `track` filters can be used while searching tracks.<br /> The `upc`, `tag:new`
	// and `tag:hipster` filters can only be used while searching albums. The `tag:new`
	// filter will return albums released in the past two weeks and `tag:hipster` can
	// be used to return only albums with the lowest 10% popularity.<br />
	Q string `query:"q,required" json:"-"`
	// A comma-separated list of item types to search across. Search results include
	// hits from all the specified item types. For example: `q=abacab&type=album,track`
	// returns both albums and tracks matching "abacab".
	//
	// Any of "album", "artist", "playlist", "track", "show", "episode", "audiobook".
	Type []string `query:"type,omitzero,required" json:"-"`
	// The maximum number of results to return in each item type.
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
	// The index of the first result to return. Use with limit to get the next page of
	// search results.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// If `include_external=audio` is specified it signals that the client can play
	// externally hosted audio content, and marks the content as playable in the
	// response. By default externally hosted audio content is marked as unplayable in
	// the response.
	//
	// Any of "audio".
	IncludeExternal SearchQueryParamsIncludeExternal `query:"include_external,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchQueryParams]'s query parameters as `url.Values`.
func (r SearchQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// If `include_external=audio` is specified it signals that the client can play
// externally hosted audio content, and marks the content as playable in the
// response. By default externally hosted audio content is marked as unplayable in
// the response.
type SearchQueryParamsIncludeExternal string

const (
	SearchQueryParamsIncludeExternalAudio SearchQueryParamsIncludeExternal = "audio"
)
