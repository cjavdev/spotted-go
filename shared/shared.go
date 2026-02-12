// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type AlbumRestrictionObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The reason for the restriction. Albums may be restricted if the content is not
	// available in a given market, to the user's subscription type, or when the user's
	// account is set to not play explicit content. Additional reasons may be added in
	// the future.
	//
	// Any of "market", "product", "explicit".
	Reason AlbumRestrictionObjectReason `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlbumRestrictionObject) RawJSON() string { return r.JSON.raw }
func (r *AlbumRestrictionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason for the restriction. Albums may be restricted if the content is not
// available in a given market, to the user's subscription type, or when the user's
// account is set to not play explicit content. Additional reasons may be added in
// the future.
type AlbumRestrictionObjectReason string

const (
	AlbumRestrictionObjectReasonMarket   AlbumRestrictionObjectReason = "market"
	AlbumRestrictionObjectReasonProduct  AlbumRestrictionObjectReason = "product"
	AlbumRestrictionObjectReasonExplicit AlbumRestrictionObjectReason = "explicit"
)

type ArtistObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// artist.
	ID string `json:"id"`
	// Known external URLs for this artist.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// Information about the followers of the artist.
	//
	// Deprecated: deprecated
	Followers FollowersObject `json:"followers"`
	// A list of the genres the artist is associated with. If not yet classified, the
	// array is empty.
	//
	// Deprecated: deprecated
	Genres []string `json:"genres"`
	// A link to the Web API endpoint providing full details of the artist.
	Href string `json:"href"`
	// Images of the artist in various sizes, widest first.
	Images []ImageObject `json:"images"`
	// The name of the artist.
	Name string `json:"name"`
	// The popularity of the artist. The value will be between 0 and 100, with 100
	// being the most popular. The artist's popularity is calculated from the
	// popularity of all the artist's tracks.
	//
	// Deprecated: deprecated
	Popularity int64 `json:"popularity"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The object type.
	//
	// Any of "artist".
	Type ArtistObjectType `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// artist.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ExternalURLs respjson.Field
		Followers    respjson.Field
		Genres       respjson.Field
		Href         respjson.Field
		Images       respjson.Field
		Name         respjson.Field
		Popularity   respjson.Field
		Published    respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArtistObject) RawJSON() string { return r.JSON.raw }
func (r *ArtistObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type ArtistObjectType string

const (
	ArtistObjectTypeArtist ArtistObjectType = "artist"
)

type AudiobookBase struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// audiobook.
	ID string `json:"id,required"`
	// The author(s) for the audiobook.
	Authors []AuthorObject `json:"authors,required"`
	// A list of the countries in which the audiobook can be played, identified by
	// their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	// code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets,required"`
	// The copyright statements of the audiobook.
	Copyrights []CopyrightObject `json:"copyrights,required"`
	// A description of the audiobook. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// Whether or not the audiobook has explicit content (true = yes it does; false =
	// no it does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this audiobook.
	ExternalURLs ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the audiobook.
	Href string `json:"href,required"`
	// A description of the audiobook. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the audiobook in various sizes, widest first.
	Images []ImageObject `json:"images,required"`
	// A list of the languages used in the audiobook, identified by their
	// [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The media type of the audiobook.
	MediaType string `json:"media_type,required"`
	// The name of the audiobook.
	Name string `json:"name,required"`
	// The narrator(s) for the audiobook.
	Narrators []NarratorObject `json:"narrators,required"`
	// The publisher of the audiobook.
	//
	// Deprecated: deprecated
	Publisher string `json:"publisher,required"`
	// The number of chapters in this audiobook.
	TotalChapters int64 `json:"total_chapters,required"`
	// The object type.
	Type constant.Audiobook `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// audiobook.
	Uri string `json:"uri,required"`
	// The edition of the audiobook.
	Edition string `json:"edition"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Authors          respjson.Field
		AvailableMarkets respjson.Field
		Copyrights       respjson.Field
		Description      respjson.Field
		Explicit         respjson.Field
		ExternalURLs     respjson.Field
		Href             respjson.Field
		HTMLDescription  respjson.Field
		Images           respjson.Field
		Languages        respjson.Field
		MediaType        respjson.Field
		Name             respjson.Field
		Narrators        respjson.Field
		Publisher        respjson.Field
		TotalChapters    respjson.Field
		Type             respjson.Field
		Uri              respjson.Field
		Edition          respjson.Field
		Published        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudiobookBase) RawJSON() string { return r.JSON.raw }
func (r *AudiobookBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthorObject struct {
	// The name of the author.
	Name string `json:"name"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthorObject) RawJSON() string { return r.JSON.raw }
func (r *AuthorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterRestrictionObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The reason for the restriction. Supported values:
	//
	//   - `market` - The content item is not available in the given market.
	//   - `product` - The content item is not available for the user's subscription
	//     type.
	//   - `explicit` - The content item is explicit and the user's account is set to not
	//     play explicit content.
	//   - `payment_required` - Payment is required to play the content item.
	//
	// Additional reasons may be added in the future. **Note**: If you use this field,
	// make sure that your application safely handles unknown values.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterRestrictionObject) RawJSON() string { return r.JSON.raw }
func (r *ChapterRestrictionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CopyrightObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The copyright text for this content.
	Text string `json:"text"`
	// The type of copyright: `C` = the copyright, `P` = the sound recording
	// (performance) copyright.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CopyrightObject) RawJSON() string { return r.JSON.raw }
func (r *CopyrightObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EpisodeObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episode.
	ID string `json:"id,required"`
	// A URL to a 30 second preview (MP3 format) of the episode. `null` if not
	// available.
	//
	// Deprecated: deprecated
	AudioPreviewURL string `json:"audio_preview_url,required"`
	// A description of the episode. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// The episode length in milliseconds.
	DurationMs int64 `json:"duration_ms,required"`
	// Whether or not the episode has explicit content (true = yes it does; false = no
	// it does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this episode.
	ExternalURLs ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the episode.
	Href string `json:"href,required"`
	// A description of the episode. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the episode in various sizes, widest first.
	Images []ImageObject `json:"images,required"`
	// True if the episode is hosted outside of Spotify's CDN.
	IsExternallyHosted bool `json:"is_externally_hosted,required"`
	// True if the episode is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable,required"`
	// A list of the languages used in the episode, identified by their
	// [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The name of the episode.
	Name string `json:"name,required"`
	// The date the episode was first released, for example `"1981-12-15"`. Depending
	// on the precision, it might be shown as `"1981"` or `"1981-12"`.
	ReleaseDate string `json:"release_date,required"`
	// The precision with which `release_date` value is known.
	//
	// Any of "year", "month", "day".
	ReleaseDatePrecision EpisodeObjectReleaseDatePrecision `json:"release_date_precision,required"`
	// The show on which the episode belongs.
	Show ShowBase `json:"show,required"`
	// The object type.
	Type constant.Episode `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episode.
	Uri string `json:"uri,required"`
	// The language used in the episode, identified by a
	// [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated
	// and might be removed in the future. Please use the `languages` field instead.
	//
	// Deprecated: deprecated
	Language string `json:"language"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions EpisodeRestrictionObject `json:"restrictions"`
	// The user's most recent position in the episode. Set if the supplied access token
	// is a user token and has the scope 'user-read-playback-position'.
	ResumePoint ResumePointObject `json:"resume_point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		IsPlayable           respjson.Field
		Languages            respjson.Field
		Name                 respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Show                 respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		Language             respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		ResumePoint          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EpisodeObject) RawJSON() string { return r.JSON.raw }
func (r *EpisodeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (EpisodeObject) ImplPlaylistTrackObjectItemUnion()                  {}
func (EpisodeObject) ImplPlaylistTrackObjectTrackUnion()                 {}
func (EpisodeObject) ImplMePlayerGetCurrentlyPlayingResponseItemUnion()  {}
func (EpisodeObject) ImplMePlayerGetStateResponseItemUnion()             {}
func (EpisodeObject) ImplMePlayerQueueGetResponseCurrentlyPlayingUnion() {}
func (EpisodeObject) ImplMePlayerQueueGetResponseQueueUnion()            {}

// The precision with which `release_date` value is known.
type EpisodeObjectReleaseDatePrecision string

const (
	EpisodeObjectReleaseDatePrecisionYear  EpisodeObjectReleaseDatePrecision = "year"
	EpisodeObjectReleaseDatePrecisionMonth EpisodeObjectReleaseDatePrecision = "month"
	EpisodeObjectReleaseDatePrecisionDay   EpisodeObjectReleaseDatePrecision = "day"
)

type EpisodeRestrictionObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The reason for the restriction. Supported values:
	//
	//   - `market` - The content item is not available in the given market.
	//   - `product` - The content item is not available for the user's subscription
	//     type.
	//   - `explicit` - The content item is explicit and the user's account is set to not
	//     play explicit content.
	//
	// Additional reasons may be added in the future. **Note**: If you use this field,
	// make sure that your application safely handles unknown values.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EpisodeRestrictionObject) RawJSON() string { return r.JSON.raw }
func (r *EpisodeRestrictionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalIDObject struct {
	// [International Article Number](http://en.wikipedia.org/wiki/International_Article_Number_%28EAN%29)
	Ean string `json:"ean"`
	// [International Standard Recording Code](http://en.wikipedia.org/wiki/International_Standard_Recording_Code)
	Isrc string `json:"isrc"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// [Universal Product Code](http://en.wikipedia.org/wiki/Universal_Product_Code)
	Upc string `json:"upc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ean         respjson.Field
		Isrc        respjson.Field
		Published   respjson.Field
		Upc         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalIDObject) RawJSON() string { return r.JSON.raw }
func (r *ExternalIDObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalURLObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the
	// object.
	Spotify string `json:"spotify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Spotify     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalURLObject) RawJSON() string { return r.JSON.raw }
func (r *ExternalURLObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FollowersObject struct {
	// This will always be set to null, as the Web API does not support it at the
	// moment.
	Href string `json:"href,nullable"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The total number of followers.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Published   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FollowersObject) RawJSON() string { return r.JSON.raw }
func (r *FollowersObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageObject struct {
	// The image height in pixels.
	Height int64 `json:"height,required"`
	// The source URL of the image.
	URL string `json:"url,required"`
	// The image width in pixels.
	Width int64 `json:"width,required"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageObject) RawJSON() string { return r.JSON.raw }
func (r *ImageObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedTrackObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	ID string `json:"id"`
	// Known external URLs for this track.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The object type: "track".
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ExternalURLs respjson.Field
		Href         respjson.Field
		Published    respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedTrackObject) RawJSON() string { return r.JSON.raw }
func (r *LinkedTrackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NarratorObject struct {
	// The name of the Narrator.
	Name string `json:"name"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NarratorObject) RawJSON() string { return r.JSON.raw }
func (r *NarratorObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PagingPlaylistObject struct {
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
	Total int64                      `json:"total,required"`
	Items []SimplifiedPlaylistObject `json:"items"`
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
func (r PagingPlaylistObject) RawJSON() string { return r.JSON.raw }
func (r *PagingPlaylistObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTrackObject struct {
	// The date and time the track or episode was added. _**Note**: some very old
	// playlists may return `null` in this field._
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// The Spotify user who added the track or episode. _**Note**: some very old
	// playlists may return `null` in this field._
	AddedBy PlaylistUserObject `json:"added_by"`
	// Whether this track or episode is a
	// [local file](/documentation/web-api/concepts/playlists/#local-files) or not.
	IsLocal bool `json:"is_local"`
	// Information about the track or episode.
	Item PlaylistTrackObjectItemUnion `json:"item"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// **Deprecated:** Use `item` instead. Information about the track or episode.
	//
	// Deprecated: deprecated
	Track PlaylistTrackObjectTrackUnion `json:"track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		AddedBy     respjson.Field
		IsLocal     respjson.Field
		Item        respjson.Field
		Published   respjson.Field
		Track       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistTrackObject) RawJSON() string { return r.JSON.raw }
func (r *PlaylistTrackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlaylistTrackObjectItemUnion contains all possible properties and values from
// [TrackObject], [EpisodeObject].
//
// Use the [PlaylistTrackObjectItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlaylistTrackObjectItemUnion struct {
	ID string `json:"id"`
	// This field is from variant [TrackObject].
	Album TrackObjectAlbum `json:"album"`
	// This field is from variant [TrackObject].
	Artists []SimplifiedArtistObject `json:"artists"`
	// This field is from variant [TrackObject].
	AvailableMarkets []string `json:"available_markets"`
	// This field is from variant [TrackObject].
	DiscNumber int64 `json:"disc_number"`
	DurationMs int64 `json:"duration_ms"`
	Explicit   bool  `json:"explicit"`
	// This field is from variant [TrackObject].
	ExternalIDs ExternalIDObject `json:"external_ids"`
	// This field is from variant [TrackObject].
	ExternalURLs ExternalURLObject `json:"external_urls"`
	Href         string            `json:"href"`
	// This field is from variant [TrackObject].
	IsLocal    bool `json:"is_local"`
	IsPlayable bool `json:"is_playable"`
	// This field is from variant [TrackObject].
	LinkedFrom LinkedTrackObject `json:"linked_from"`
	Name       string            `json:"name"`
	// This field is from variant [TrackObject].
	Popularity int64 `json:"popularity"`
	// This field is from variant [TrackObject].
	PreviewURL string `json:"preview_url"`
	Published  bool   `json:"published"`
	// This field is a union of [TrackRestrictionObject], [EpisodeRestrictionObject]
	Restrictions PlaylistTrackObjectItemUnionRestrictions `json:"restrictions"`
	// This field is from variant [TrackObject].
	TrackNumber int64 `json:"track_number"`
	// Any of "track", "episode".
	Type string `json:"type"`
	Uri  string `json:"uri"`
	// This field is from variant [EpisodeObject].
	AudioPreviewURL string `json:"audio_preview_url"`
	// This field is from variant [EpisodeObject].
	Description string `json:"description"`
	// This field is from variant [EpisodeObject].
	HTMLDescription string `json:"html_description"`
	// This field is from variant [EpisodeObject].
	Images []ImageObject `json:"images"`
	// This field is from variant [EpisodeObject].
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// This field is from variant [EpisodeObject].
	Languages []string `json:"languages"`
	// This field is from variant [EpisodeObject].
	ReleaseDate string `json:"release_date"`
	// This field is from variant [EpisodeObject].
	ReleaseDatePrecision EpisodeObjectReleaseDatePrecision `json:"release_date_precision"`
	// This field is from variant [EpisodeObject].
	Show ShowBase `json:"show"`
	// This field is from variant [EpisodeObject].
	Language string `json:"language"`
	// This field is from variant [EpisodeObject].
	ResumePoint ResumePointObject `json:"resume_point"`
	JSON        struct {
		ID                   respjson.Field
		Album                respjson.Field
		Artists              respjson.Field
		AvailableMarkets     respjson.Field
		DiscNumber           respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalIDs          respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		IsLocal              respjson.Field
		IsPlayable           respjson.Field
		LinkedFrom           respjson.Field
		Name                 respjson.Field
		Popularity           respjson.Field
		PreviewURL           respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		TrackNumber          respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		Languages            respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Show                 respjson.Field
		Language             respjson.Field
		ResumePoint          respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPlaylistTrackObjectItem is implemented by each variant of
// [PlaylistTrackObjectItemUnion] to add type safety for the return type of
// [PlaylistTrackObjectItemUnion.AsAny]
type anyPlaylistTrackObjectItem interface {
	ImplPlaylistTrackObjectItemUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := PlaylistTrackObjectItemUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PlaylistTrackObjectItemUnion) AsAny() anyPlaylistTrackObjectItem {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u PlaylistTrackObjectItemUnion) AsTrack() (v TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlaylistTrackObjectItemUnion) AsEpisode() (v EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlaylistTrackObjectItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PlaylistTrackObjectItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlaylistTrackObjectItemUnionRestrictions is an implicit subunion of
// [PlaylistTrackObjectItemUnion]. PlaylistTrackObjectItemUnionRestrictions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PlaylistTrackObjectItemUnion].
type PlaylistTrackObjectItemUnionRestrictions struct {
	Published bool   `json:"published"`
	Reason    string `json:"reason"`
	JSON      struct {
		Published respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (r *PlaylistTrackObjectItemUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlaylistTrackObjectTrackUnion contains all possible properties and values from
// [TrackObject], [EpisodeObject].
//
// Use the [PlaylistTrackObjectTrackUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlaylistTrackObjectTrackUnion struct {
	ID string `json:"id"`
	// This field is from variant [TrackObject].
	Album TrackObjectAlbum `json:"album"`
	// This field is from variant [TrackObject].
	Artists []SimplifiedArtistObject `json:"artists"`
	// This field is from variant [TrackObject].
	AvailableMarkets []string `json:"available_markets"`
	// This field is from variant [TrackObject].
	DiscNumber int64 `json:"disc_number"`
	DurationMs int64 `json:"duration_ms"`
	Explicit   bool  `json:"explicit"`
	// This field is from variant [TrackObject].
	ExternalIDs ExternalIDObject `json:"external_ids"`
	// This field is from variant [TrackObject].
	ExternalURLs ExternalURLObject `json:"external_urls"`
	Href         string            `json:"href"`
	// This field is from variant [TrackObject].
	IsLocal    bool `json:"is_local"`
	IsPlayable bool `json:"is_playable"`
	// This field is from variant [TrackObject].
	LinkedFrom LinkedTrackObject `json:"linked_from"`
	Name       string            `json:"name"`
	// This field is from variant [TrackObject].
	Popularity int64 `json:"popularity"`
	// This field is from variant [TrackObject].
	PreviewURL string `json:"preview_url"`
	Published  bool   `json:"published"`
	// This field is a union of [TrackRestrictionObject], [EpisodeRestrictionObject]
	Restrictions PlaylistTrackObjectTrackUnionRestrictions `json:"restrictions"`
	// This field is from variant [TrackObject].
	TrackNumber int64 `json:"track_number"`
	// Any of "track", "episode".
	Type string `json:"type"`
	Uri  string `json:"uri"`
	// This field is from variant [EpisodeObject].
	AudioPreviewURL string `json:"audio_preview_url"`
	// This field is from variant [EpisodeObject].
	Description string `json:"description"`
	// This field is from variant [EpisodeObject].
	HTMLDescription string `json:"html_description"`
	// This field is from variant [EpisodeObject].
	Images []ImageObject `json:"images"`
	// This field is from variant [EpisodeObject].
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// This field is from variant [EpisodeObject].
	Languages []string `json:"languages"`
	// This field is from variant [EpisodeObject].
	ReleaseDate string `json:"release_date"`
	// This field is from variant [EpisodeObject].
	ReleaseDatePrecision EpisodeObjectReleaseDatePrecision `json:"release_date_precision"`
	// This field is from variant [EpisodeObject].
	Show ShowBase `json:"show"`
	// This field is from variant [EpisodeObject].
	Language string `json:"language"`
	// This field is from variant [EpisodeObject].
	ResumePoint ResumePointObject `json:"resume_point"`
	JSON        struct {
		ID                   respjson.Field
		Album                respjson.Field
		Artists              respjson.Field
		AvailableMarkets     respjson.Field
		DiscNumber           respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalIDs          respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		IsLocal              respjson.Field
		IsPlayable           respjson.Field
		LinkedFrom           respjson.Field
		Name                 respjson.Field
		Popularity           respjson.Field
		PreviewURL           respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		TrackNumber          respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		Languages            respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Show                 respjson.Field
		Language             respjson.Field
		ResumePoint          respjson.Field
		raw                  string
	} `json:"-"`
}

// anyPlaylistTrackObjectTrack is implemented by each variant of
// [PlaylistTrackObjectTrackUnion] to add type safety for the return type of
// [PlaylistTrackObjectTrackUnion.AsAny]
type anyPlaylistTrackObjectTrack interface {
	ImplPlaylistTrackObjectTrackUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := PlaylistTrackObjectTrackUnion.AsAny().(type) {
//	case shared.TrackObject:
//	case shared.EpisodeObject:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PlaylistTrackObjectTrackUnion) AsAny() anyPlaylistTrackObjectTrack {
	switch u.Type {
	case "track":
		return u.AsTrack()
	case "episode":
		return u.AsEpisode()
	}
	return nil
}

func (u PlaylistTrackObjectTrackUnion) AsTrack() (v TrackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlaylistTrackObjectTrackUnion) AsEpisode() (v EpisodeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlaylistTrackObjectTrackUnion) RawJSON() string { return u.JSON.raw }

func (r *PlaylistTrackObjectTrackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlaylistTrackObjectTrackUnionRestrictions is an implicit subunion of
// [PlaylistTrackObjectTrackUnion]. PlaylistTrackObjectTrackUnionRestrictions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PlaylistTrackObjectTrackUnion].
type PlaylistTrackObjectTrackUnionRestrictions struct {
	Published bool   `json:"published"`
	Reason    string `json:"reason"`
	JSON      struct {
		Published respjson.Field
		Reason    respjson.Field
		raw       string
	} `json:"-"`
}

func (r *PlaylistTrackObjectTrackUnionRestrictions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistTracksRefObject struct {
	// A link to the Web API endpoint where full details of the playlist's tracks can
	// be retrieved.
	Href string `json:"href"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Number of tracks in the playlist.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Published   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistTracksRefObject) RawJSON() string { return r.JSON.raw }
func (r *PlaylistTracksRefObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistUserObject struct {
	// The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this
	// user.
	ID string `json:"id"`
	// Known public external URLs for this user.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The object type.
	//
	// Any of "user".
	Type PlaylistUserObjectType `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this
	// user.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ExternalURLs respjson.Field
		Href         respjson.Field
		Published    respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlaylistUserObject) RawJSON() string { return r.JSON.raw }
func (r *PlaylistUserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type PlaylistUserObjectType string

const (
	PlaylistUserObjectTypeUser PlaylistUserObjectType = "user"
)

type ResumePointObject struct {
	// Whether or not the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The user's most recent position in the episode in milliseconds.
	ResumePositionMs int64 `json:"resume_position_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullyPlayed      respjson.Field
		Published        respjson.Field
		ResumePositionMs respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResumePointObject) RawJSON() string { return r.JSON.raw }
func (r *ResumePointObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShowBase struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the show.
	ID string `json:"id,required"`
	// A list of the countries in which the show can be played, identified by their
	// [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets,required"`
	// The copyright statements of the show.
	Copyrights []CopyrightObject `json:"copyrights,required"`
	// A description of the show. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// Whether or not the show has explicit content (true = yes it does; false = no it
	// does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this show.
	ExternalURLs ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the show.
	Href string `json:"href,required"`
	// A description of the show. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the show in various sizes, widest first.
	Images []ImageObject `json:"images,required"`
	// True if all of the shows episodes are hosted outside of Spotify's CDN. This
	// field might be `null` in some cases.
	IsExternallyHosted bool `json:"is_externally_hosted,required"`
	// A list of the languages used in the show, identified by their
	// [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The media type of the show.
	MediaType string `json:"media_type,required"`
	// The name of the episode.
	Name string `json:"name,required"`
	// The publisher of the show.
	//
	// Deprecated: deprecated
	Publisher string `json:"publisher,required"`
	// The total number of episodes in the show.
	TotalEpisodes int64 `json:"total_episodes,required"`
	// The object type.
	Type constant.Show `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// show.
	Uri string `json:"uri,required"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AvailableMarkets   respjson.Field
		Copyrights         respjson.Field
		Description        respjson.Field
		Explicit           respjson.Field
		ExternalURLs       respjson.Field
		Href               respjson.Field
		HTMLDescription    respjson.Field
		Images             respjson.Field
		IsExternallyHosted respjson.Field
		Languages          respjson.Field
		MediaType          respjson.Field
		Name               respjson.Field
		Publisher          respjson.Field
		TotalEpisodes      respjson.Field
		Type               respjson.Field
		Uri                respjson.Field
		Published          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShowBase) RawJSON() string { return r.JSON.raw }
func (r *ShowBase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimplifiedArtistObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// artist.
	ID string `json:"id"`
	// Known external URLs for this artist.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the artist.
	Href string `json:"href"`
	// The name of the artist.
	Name string `json:"name"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The object type.
	//
	// Any of "artist".
	Type SimplifiedArtistObjectType `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// artist.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ExternalURLs respjson.Field
		Href         respjson.Field
		Name         respjson.Field
		Published    respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplifiedArtistObject) RawJSON() string { return r.JSON.raw }
func (r *SimplifiedArtistObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type SimplifiedArtistObjectType string

const (
	SimplifiedArtistObjectTypeArtist SimplifiedArtistObjectType = "artist"
)

type SimplifiedEpisodeObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episode.
	ID string `json:"id,required"`
	// A URL to a 30 second preview (MP3 format) of the episode. `null` if not
	// available.
	//
	// Deprecated: deprecated
	AudioPreviewURL string `json:"audio_preview_url,required"`
	// A description of the episode. HTML tags are stripped away from this field, use
	// `html_description` field in case HTML tags are needed.
	Description string `json:"description,required"`
	// The episode length in milliseconds.
	DurationMs int64 `json:"duration_ms,required"`
	// Whether or not the episode has explicit content (true = yes it does; false = no
	// it does not OR unknown).
	Explicit bool `json:"explicit,required"`
	// External URLs for this episode.
	ExternalURLs ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the episode.
	Href string `json:"href,required"`
	// A description of the episode. This field may contain HTML tags.
	HTMLDescription string `json:"html_description,required"`
	// The cover art for the episode in various sizes, widest first.
	Images []ImageObject `json:"images,required"`
	// True if the episode is hosted outside of Spotify's CDN.
	IsExternallyHosted bool `json:"is_externally_hosted,required"`
	// True if the episode is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable,required"`
	// A list of the languages used in the episode, identified by their
	// [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.
	Languages []string `json:"languages,required"`
	// The name of the episode.
	Name string `json:"name,required"`
	// The date the episode was first released, for example `"1981-12-15"`. Depending
	// on the precision, it might be shown as `"1981"` or `"1981-12"`.
	ReleaseDate string `json:"release_date,required"`
	// The precision with which `release_date` value is known.
	//
	// Any of "year", "month", "day".
	ReleaseDatePrecision SimplifiedEpisodeObjectReleaseDatePrecision `json:"release_date_precision,required"`
	// The object type.
	Type constant.Episode `json:"type,required"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// episode.
	Uri string `json:"uri,required"`
	// The language used in the episode, identified by a
	// [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated
	// and might be removed in the future. Please use the `languages` field instead.
	//
	// Deprecated: deprecated
	Language string `json:"language"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions EpisodeRestrictionObject `json:"restrictions"`
	// The user's most recent position in the episode. Set if the supplied access token
	// is a user token and has the scope 'user-read-playback-position'.
	ResumePoint ResumePointObject `json:"resume_point"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AudioPreviewURL      respjson.Field
		Description          respjson.Field
		DurationMs           respjson.Field
		Explicit             respjson.Field
		ExternalURLs         respjson.Field
		Href                 respjson.Field
		HTMLDescription      respjson.Field
		Images               respjson.Field
		IsExternallyHosted   respjson.Field
		IsPlayable           respjson.Field
		Languages            respjson.Field
		Name                 respjson.Field
		ReleaseDate          respjson.Field
		ReleaseDatePrecision respjson.Field
		Type                 respjson.Field
		Uri                  respjson.Field
		Language             respjson.Field
		Published            respjson.Field
		Restrictions         respjson.Field
		ResumePoint          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplifiedEpisodeObject) RawJSON() string { return r.JSON.raw }
func (r *SimplifiedEpisodeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The precision with which `release_date` value is known.
type SimplifiedEpisodeObjectReleaseDatePrecision string

const (
	SimplifiedEpisodeObjectReleaseDatePrecisionYear  SimplifiedEpisodeObjectReleaseDatePrecision = "year"
	SimplifiedEpisodeObjectReleaseDatePrecisionMonth SimplifiedEpisodeObjectReleaseDatePrecision = "month"
	SimplifiedEpisodeObjectReleaseDatePrecisionDay   SimplifiedEpisodeObjectReleaseDatePrecision = "day"
)

type SimplifiedPlaylistObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// playlist.
	ID string `json:"id"`
	// `true` if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// The playlist description. _Only returned for modified, verified playlists,
	// otherwise_ `null`.
	Description string `json:"description"`
	// Known external URLs for this playlist.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// Images for the playlist. The array may be empty or contain up to three images.
	// The images are returned by size in descending order. See
	// [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**:
	// If returned, the source URL for the image (`url`) is temporary and will expire
	// in less than a day._
	Images []ImageObject `json:"images"`
	// A collection containing a link ( `href` ) to the Web API endpoint where full
	// details of the playlist's items can be retrieved, along with the `total` number
	// of items in the playlist. Note, a track object may be `null`. This can happen if
	// a track is no longer available.
	Items PlaylistTracksRefObject `json:"items"`
	// The name of the playlist.
	Name string `json:"name"`
	// The user who owns the playlist
	Owner SimplifiedPlaylistObjectOwner `json:"owner"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The version identifier for the current playlist. Can be supplied in other
	// requests to target a specific playlist version
	SnapshotID string `json:"snapshot_id"`
	// **Deprecated:** Use `items` instead. A collection containing a link ( `href` )
	// to the Web API endpoint where full details of the playlist's tracks can be
	// retrieved, along with the `total` number of tracks in the playlist. Note, a
	// track object may be `null`. This can happen if a track is no longer available.
	//
	// Deprecated: deprecated
	Tracks PlaylistTracksRefObject `json:"tracks"`
	// The object type: "playlist"
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// playlist.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Collaborative respjson.Field
		Description   respjson.Field
		ExternalURLs  respjson.Field
		Href          respjson.Field
		Images        respjson.Field
		Items         respjson.Field
		Name          respjson.Field
		Owner         respjson.Field
		Published     respjson.Field
		SnapshotID    respjson.Field
		Tracks        respjson.Field
		Type          respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplifiedPlaylistObject) RawJSON() string { return r.JSON.raw }
func (r *SimplifiedPlaylistObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user who owns the playlist
type SimplifiedPlaylistObjectOwner struct {
	// The name displayed on the user's profile. `null` if not available.
	DisplayName string `json:"display_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	PlaylistUserObject
}

// Returns the unmodified JSON received from the API
func (r SimplifiedPlaylistObjectOwner) RawJSON() string { return r.JSON.raw }
func (r *SimplifiedPlaylistObjectOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimplifiedTrackObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	ID string `json:"id"`
	// The artists who performed the track. Each artist object includes a link in
	// `href` to more detailed information about the artist.
	Artists []SimplifiedArtistObject `json:"artists"`
	// A list of the countries in which the track can be played, identified by their
	// [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually `1` unless the album consists of more than one disc).
	DiscNumber int64 `json:"disc_number"`
	// The track length in milliseconds.
	DurationMs int64 `json:"duration_ms"`
	// Whether or not the track has explicit lyrics ( `true` = yes it does; `false` =
	// no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this track.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
	// Part of the response when
	// [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied.
	// If `true`, the track is playable in the given market. Otherwise `false`.
	IsPlayable bool `json:"is_playable"`
	// Part of the response when
	// [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied
	// and is only part of the response if the track linking, in fact, exists. The
	// requested track has been replaced with a different track. The track in the
	// `linked_from` object contains information about the originally requested track.
	//
	// Deprecated: deprecated
	LinkedFrom LinkedTrackObject `json:"linked_from"`
	// The name of the track.
	Name string `json:"name"`
	// A URL to a 30 second preview (MP3 format) of the track.
	//
	// Deprecated: deprecated
	PreviewURL string `json:"preview_url,nullable"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions TrackRestrictionObject `json:"restrictions"`
	// The number of the track. If an album has several discs, the track number is the
	// number on the specified disc.
	TrackNumber int64 `json:"track_number"`
	// The object type: "track".
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Artists          respjson.Field
		AvailableMarkets respjson.Field
		DiscNumber       respjson.Field
		DurationMs       respjson.Field
		Explicit         respjson.Field
		ExternalURLs     respjson.Field
		Href             respjson.Field
		IsLocal          respjson.Field
		IsPlayable       respjson.Field
		LinkedFrom       respjson.Field
		Name             respjson.Field
		PreviewURL       respjson.Field
		Published        respjson.Field
		Restrictions     respjson.Field
		TrackNumber      respjson.Field
		Type             respjson.Field
		Uri              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplifiedTrackObject) RawJSON() string { return r.JSON.raw }
func (r *SimplifiedTrackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TrackObject struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	ID string `json:"id"`
	// The album on which the track appears. The album object includes a link in `href`
	// to full information about the album.
	Album TrackObjectAlbum `json:"album"`
	// The artists who performed the track. Each artist object includes a link in
	// `href` to more detailed information about the artist.
	Artists []SimplifiedArtistObject `json:"artists"`
	// A list of the countries in which the track can be played, identified by their
	// [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually `1` unless the album consists of more than one disc).
	DiscNumber int64 `json:"disc_number"`
	// The track length in milliseconds.
	DurationMs int64 `json:"duration_ms"`
	// Whether or not the track has explicit lyrics ( `true` = yes it does; `false` =
	// no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// Known external IDs for the track.
	//
	// Deprecated: deprecated
	ExternalIDs ExternalIDObject `json:"external_ids"`
	// Known external URLs for this track.
	ExternalURLs ExternalURLObject `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
	// Part of the response when
	// [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied.
	// If `true`, the track is playable in the given market. Otherwise `false`.
	IsPlayable bool `json:"is_playable"`
	// Part of the response when
	// [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied,
	// and the requested track has been replaced with different track. The track in the
	// `linked_from` object contains information about the originally requested track.
	//
	// Deprecated: deprecated
	LinkedFrom LinkedTrackObject `json:"linked_from"`
	// The name of the track.
	Name string `json:"name"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being
	// the most popular.<br/>The popularity of a track is a value between 0 and 100,
	// with 100 being the most popular. The popularity is calculated by algorithm and
	// is based, in the most part, on the total number of plays the track has had and
	// how recent those plays are.<br/>Generally speaking, songs that are being played
	// a lot now will have a higher popularity than songs that were played a lot in the
	// past. Duplicate tracks (e.g. the same track from a single and an album) are
	// rated independently. Artist and album popularity is derived mathematically from
	// track popularity. _**Note**: the popularity value may lag actual popularity by a
	// few days: the value is not updated in real time._
	//
	// Deprecated: deprecated
	Popularity int64 `json:"popularity"`
	// A link to a 30 second preview (MP3 format) of the track. Can be `null`
	//
	// Deprecated: deprecated
	PreviewURL string `json:"preview_url,nullable"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions TrackRestrictionObject `json:"restrictions"`
	// The number of the track. If an album has several discs, the track number is the
	// number on the specified disc.
	TrackNumber int64 `json:"track_number"`
	// The object type: "track".
	//
	// Any of "track".
	Type TrackObjectType `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// track.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Album            respjson.Field
		Artists          respjson.Field
		AvailableMarkets respjson.Field
		DiscNumber       respjson.Field
		DurationMs       respjson.Field
		Explicit         respjson.Field
		ExternalIDs      respjson.Field
		ExternalURLs     respjson.Field
		Href             respjson.Field
		IsLocal          respjson.Field
		IsPlayable       respjson.Field
		LinkedFrom       respjson.Field
		Name             respjson.Field
		Popularity       respjson.Field
		PreviewURL       respjson.Field
		Published        respjson.Field
		Restrictions     respjson.Field
		TrackNumber      respjson.Field
		Type             respjson.Field
		Uri              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackObject) RawJSON() string { return r.JSON.raw }
func (r *TrackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (TrackObject) ImplPlaylistTrackObjectItemUnion()                  {}
func (TrackObject) ImplPlaylistTrackObjectTrackUnion()                 {}
func (TrackObject) ImplMePlayerGetCurrentlyPlayingResponseItemUnion()  {}
func (TrackObject) ImplMePlayerGetStateResponseItemUnion()             {}
func (TrackObject) ImplMePlayerQueueGetResponseCurrentlyPlayingUnion() {}
func (TrackObject) ImplMePlayerQueueGetResponseQueueUnion()            {}

// The album on which the track appears. The album object includes a link in `href`
// to full information about the album.
type TrackObjectAlbum struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// album.
	ID string `json:"id,required"`
	// The type of the album.
	//
	// Any of "album", "single", "compilation".
	AlbumType string `json:"album_type,required"`
	// The artists of the album. Each artist object includes a link in `href` to more
	// detailed information about the artist.
	Artists []SimplifiedArtistObject `json:"artists,required"`
	// The markets in which the album is available:
	// [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// _**NOTE**: an album is considered available in a market when at least 1 of its
	// tracks is available in that market._
	//
	// Deprecated: deprecated
	AvailableMarkets []string `json:"available_markets,required"`
	// Known external URLs for this album.
	ExternalURLs ExternalURLObject `json:"external_urls,required"`
	// A link to the Web API endpoint providing full details of the album.
	Href string `json:"href,required"`
	// The cover art for the album in various sizes, widest first.
	Images []ImageObject `json:"images,required"`
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
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// Included in the response when a content restriction is applied.
	Restrictions AlbumRestrictionObject `json:"restrictions"`
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
		Published            respjson.Field
		Restrictions         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackObjectAlbum) RawJSON() string { return r.JSON.raw }
func (r *TrackObjectAlbum) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type: "track".
type TrackObjectType string

const (
	TrackObjectTypeTrack TrackObjectType = "track"
)

type TrackRestrictionObject struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The reason for the restriction. Supported values:
	//
	//   - `market` - The content item is not available in the given market.
	//   - `product` - The content item is not available for the user's subscription
	//     type.
	//   - `explicit` - The content item is explicit and the user's account is set to not
	//     play explicit content.
	//
	// Additional reasons may be added in the future. **Note**: If you use this field,
	// make sure that your application safely handles unknown values.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Published   respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TrackRestrictionObject) RawJSON() string { return r.JSON.raw }
func (r *TrackRestrictionObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
