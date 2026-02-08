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
	"github.com/cjavdev/spotted-go/shared/constant"
)

// ChapterService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChapterService] method instead.
type ChapterService struct {
	Options []option.RequestOption
}

// NewChapterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChapterService(opts ...option.RequestOption) (r ChapterService) {
	r = ChapterService{}
	r.Options = opts
	return
}

// Get Spotify catalog information for a single audiobook chapter. Chapters are
// only available within the US, UK, Canada, Ireland, New Zealand and Australia
// markets.
func (r *ChapterService) Get(ctx context.Context, id string, query ChapterGetParams, opts ...option.RequestOption) (res *ChapterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("chapters/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Spotify catalog information for several audiobook chapters identified by
// their Spotify IDs. Chapters are only available within the US, UK, Canada,
// Ireland, New Zealand and Australia markets.
//
// Deprecated: deprecated
func (r *ChapterService) BulkGet(ctx context.Context, query ChapterBulkGetParams, opts ...option.RequestOption) (res *ChapterBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "chapters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ChapterGetResponse struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// chapter.
	ID string `json:"id,required"`
	// A URL to a 30 second preview (MP3 format) of the chapter. `null` if not
	// available.
	//
	// Deprecated: deprecated
	AudioPreviewURL string `json:"audio_preview_url,required"`
	// The audiobook for which the chapter belongs.
	Audiobook shared.AudiobookBase `json:"audiobook,required"`
	// The number of the chapter
	ChapterNumber int64 `json:"chapter_number,required"`
	// A description of the chapter. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// The chapter length in milliseconds.
	DurationMs int64 `json:"duration_ms,required"`
	// Whether or not the chapter has explicit content (true = yes it does; false = no
	// it does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this chapter.
	ExternalURLs shared.ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the chapter.
	Href string `json:"href,required"`
	// A description of the chapter. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the chapter in various sizes, widest first.
	Images []shared.ImageObject `json:"images,required"`
	// True if the chapter is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable,required"`
	// A list of the languages used in the chapter, identified by their
	// [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The name of the chapter.
	Name string `json:"name,required"`
	// The date the chapter was first released, for example `"1981-12-15"`. Depending
	// on the precision, it might be shown as `"1981"` or `"1981-12"`.
	ReleaseDate string `json:"release_date,required"`
	// The precision with which `release_date` value is known.
	//
	// Any of "year", "month", "day".
	ReleaseDatePrecision ChapterGetResponseReleaseDatePrecision `json:"release_date_precision,required"`
	// The object type.
	Type constant.Episode `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// chapter.
	Uri string `json:"uri,required"`
	// A list of the countries in which the chapter can be played, identified by their
	// [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions shared.ChapterRestrictionObject `json:"restrictions"`
	// The user's most recent position in the chapter. Set if the supplied access token
	// is a user token and has the scope 'user-read-playback-position'.
	ResumePoint shared.ResumePointObject `json:"resume_point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AudioPreviewURL      respjson.Field
		Audiobook            respjson.Field
		ChapterNumber        respjson.Field
		Description          respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsPlayable           respjson.Field
		Languages            respjson.Field
		Name                 respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AvailableMarkets     respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		ResumePoint          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChapterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The precision with which `release_date` value is known.
type ChapterGetResponseReleaseDatePrecision string

const (
	ChapterGetResponseReleaseDatePrecisionYear  ChapterGetResponseReleaseDatePrecision = "year"
	ChapterGetResponseReleaseDatePrecisionMonth ChapterGetResponseReleaseDatePrecision = "month"
	ChapterGetResponseReleaseDatePrecisionDay   ChapterGetResponseReleaseDatePrecision = "day"
)

type ChapterBulkGetResponse struct {
	Chapters []ChapterBulkGetResponseChapter `json:"chapters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ChapterBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterBulkGetResponseChapter struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// chapter.
	ID string `json:"id,required"`
	// A URL to a 30 second preview (MP3 format) of the chapter. `null` if not
	// available.
	//
	// Deprecated: deprecated
	AudioPreviewURL string `json:"audio_preview_url,required"`
	// The audiobook for which the chapter belongs.
	Audiobook shared.AudiobookBase `json:"audiobook,required"`
	// The number of the chapter
	ChapterNumber int64 `json:"chapter_number,required"`
	// A description of the chapter. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// The chapter length in milliseconds.
	DurationMs int64 `json:"duration_ms,required"`
	// Whether or not the chapter has explicit content (true = yes it does; false = no
	// it does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this chapter.
	ExternalURLs shared.ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the chapter.
	Href string `json:"href,required"`
	// A description of the chapter. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the chapter in various sizes, widest first.
	Images []shared.ImageObject `json:"images,required"`
	// True if the chapter is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable,required"`
	// A list of the languages used in the chapter, identified by their
	// [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The name of the chapter.
	Name string `json:"name,required"`
	// The date the chapter was first released, for example `"1981-12-15"`. Depending
	// on the precision, it might be shown as `"1981"` or `"1981-12"`.
	ReleaseDate string `json:"release_date,required"`
	// The precision with which `release_date` value is known.
	//
	// Any of "year", "month", "day".
	ReleaseDatePrecision string `json:"release_date_precision,required"`
	// The object type.
	Type constant.Episode `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// chapter.
	Uri string `json:"uri,required"`
	// A list of the countries in which the chapter can be played, identified by their
	// [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions shared.ChapterRestrictionObject `json:"restrictions"`
	// The user's most recent position in the chapter. Set if the supplied access token
	// is a user token and has the scope 'user-read-playback-position'.
	ResumePoint shared.ResumePointObject `json:"resume_point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AudioPreviewURL      respjson.Field
		Audiobook            respjson.Field
		ChapterNumber        respjson.Field
		Description          respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsPlayable           respjson.Field
		Languages            respjson.Field
		Name                 respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AvailableMarkets     respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		ResumePoint          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterBulkGetResponseChapter) RawJSON() string { return r.JSON.raw }
func (r *ChapterBulkGetResponseChapter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterGetParams struct {
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

// URLQuery serializes [ChapterGetParams]'s query parameters as `url.Values`.
func (r ChapterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ChapterBulkGetParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=0IsXVP0JmcB2adSE338GkK,3ZXb8FKZGU0EHALYX6uCzU`. Maximum: 50 IDs.
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

// URLQuery serializes [ChapterBulkGetParams]'s query parameters as `url.Values`.
func (r ChapterBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
