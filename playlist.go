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

// PlaylistService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistService] method instead.
type PlaylistService struct {
	Options   []option.RequestOption
	Tracks    PlaylistTrackService
	Followers PlaylistFollowerService
	Images    PlaylistImageService
}

// NewPlaylistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaylistService(opts ...option.RequestOption) (r PlaylistService) {
	r = PlaylistService{}
	r.Options = opts
	r.Tracks = NewPlaylistTrackService(opts...)
	r.Followers = NewPlaylistFollowerService(opts...)
	r.Images = NewPlaylistImageService(opts...)
	return
}

// Get a playlist owned by a Spotify user.
func (r *PlaylistService) Get(ctx context.Context, playlistID string, query PlaylistGetParams, opts ...option.RequestOption) (res *PlaylistGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Change a playlist's name and public/private state. (The user must, of course,
// own the playlist.)
func (r *PlaylistService) Update(ctx context.Context, playlistID string, body PlaylistUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type PlaylistGetResponse struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// playlist.
	ID string `json:"id"`
	// `true` if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// The playlist description. _Only returned for modified, verified playlists,
	// otherwise_ `null`.
	Description string `json:"description,nullable"`
	// Known external URLs for this playlist.
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	// Information about the followers of the playlist.
	Followers shared.FollowersObject `json:"followers"`
	// A link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// Images for the playlist. The array may be empty or contain up to three images.
	// The images are returned by size in descending order. See
	// [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**:
	// If returned, the source URL for the image (`url`) is temporary and will expire
	// in less than a day._
	Images []shared.ImageObject `json:"images"`
	// The name of the playlist.
	Name string `json:"name"`
	// The user who owns the playlist
	Owner PlaylistGetResponseOwner `json:"owner"`
	// The playlist's public/private status (if it is added to the user's profile):
	// `true` the playlist is public, `false` the playlist is private, `null` the
	// playlist status is not relevant. For more about public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The version identifier for the current playlist. Can be supplied in other
	// requests to target a specific playlist version
	SnapshotID string `json:"snapshot_id"`
	// The tracks of the playlist.
	Tracks PlaylistGetResponseTracks `json:"tracks"`
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
		Followers     respjson.Field
		Href          respjson.Field
		Images        respjson.Field
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
func (r PlaylistGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PlaylistGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user who owns the playlist
type PlaylistGetResponseOwner struct {
	// The name displayed on the user's profile. `null` if not available.
	DisplayName string `json:"display_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.PlaylistUserObject
}

// Returns the unmodified JSON received from the API
func (r PlaylistGetResponseOwner) RawJSON() string { return r.JSON.raw }
func (r *PlaylistGetResponseOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The tracks of the playlist.
type PlaylistGetResponseTracks struct {
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
	Total int64                        `json:"total,required"`
	Items []shared.PlaylistTrackObject `json:"items"`
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
func (r PlaylistGetResponseTracks) RawJSON() string { return r.JSON.raw }
func (r *PlaylistGetResponseTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlaylistGetParams struct {
	// A comma-separated list of item types that your client supports besides the
	// default `track` type. Valid types are: `track` and `episode`.<br/> _**Note**:
	// This parameter was introduced to allow existing clients to maintain their
	// current behaviour and might be deprecated in the future._<br/> In addition to
	// providing this parameter, make sure that your client properly handles cases of
	// new types in the future by checking against the `type` field of each object.
	AdditionalTypes param.Opt[string] `query:"additional_types,omitzero" json:"-"`
	// Filters for the query: a comma-separated list of the fields to return. If
	// omitted, all fields are returned. For example, to get just the playlist”s
	// description and URI: `fields=description,uri`. A dot separator can be used to
	// specify non-reoccurring fields, while parentheses can be used to specify
	// reoccurring fields within objects. For example, to get just the added date and
	// user ID of the adder: `fields=tracks.items(added_at,added_by.id)`. Use multiple
	// parentheses to drill down into nested objects, for example:
	// `fields=tracks.items(track(name,href,album(name,href)))`. Fields can be excluded
	// by prefixing them with an exclamation mark, for example:
	// `fields=tracks.items(track(name,href,album(!name,href)))`
	Fields param.Opt[string] `query:"fields,omitzero" json:"-"`
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

// URLQuery serializes [PlaylistGetParams]'s query parameters as `url.Values`.
func (r PlaylistGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlaylistUpdateParams struct {
	// If `true`, the playlist will become collaborative and other users will be able
	// to modify the playlist in their Spotify client. <br/> _**Note**: You can only
	// set `collaborative` to `true` on non-public playlists._
	Collaborative param.Opt[bool] `json:"collaborative,omitzero"`
	// Value for playlist description as displayed in Spotify Clients and in the Web
	// API.
	Description param.Opt[string] `json:"description,omitzero"`
	// The new name for the playlist, for example `"My New Playlist Title"`
	Name param.Opt[string] `json:"name,omitzero"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Public param.Opt[bool] `json:"public,omitzero"`
	paramObj
}

func (r PlaylistUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
