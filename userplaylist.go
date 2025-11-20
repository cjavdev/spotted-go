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
)

// UserPlaylistService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserPlaylistService] method instead.
type UserPlaylistService struct {
	Options []option.RequestOption
}

// NewUserPlaylistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserPlaylistService(opts ...option.RequestOption) (r UserPlaylistService) {
	r = UserPlaylistService{}
	r.Options = opts
	return
}

// Create a playlist for a Spotify user. (The playlist will be empty until you
// [add tracks](/documentation/web-api/reference/add-tracks-to-playlist).) Each
// user is generally limited to a maximum of 11000 playlists.
func (r *UserPlaylistService) New(ctx context.Context, userID string, body UserPlaylistNewParams, opts ...option.RequestOption) (res *UserPlaylistNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/playlists", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of the playlists owned or followed by a Spotify user.
func (r *UserPlaylistService) List(ctx context.Context, userID string, query UserPlaylistListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.SimplifiedPlaylistObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s/playlists", userID)
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

// Get a list of the playlists owned or followed by a Spotify user.
func (r *UserPlaylistService) ListAutoPaging(ctx context.Context, userID string, query UserPlaylistListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.SimplifiedPlaylistObject] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, userID, query, opts...))
}

type UserPlaylistNewResponse struct {
	// The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// playlist.
	ID string `json:"id"`
	// The playlist's public/private status (if it is added to the user's profile):
	// `true` the playlist is public, `false` the playlist is private, `null` the
	// playlist status is not relevant. For more about public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	ComponentsSchemasPropertiesPublished bool `json:"$.components.schemas.*.properties.published"`
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
	Owner UserPlaylistNewResponseOwner `json:"owner"`
	// The version identifier for the current playlist. Can be supplied in other
	// requests to target a specific playlist version
	SnapshotID string `json:"snapshot_id"`
	// The tracks of the playlist.
	Tracks UserPlaylistNewResponseTracks `json:"tracks"`
	// The object type: "playlist"
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// playlist.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                                   respjson.Field
		ComponentsSchemasPropertiesPublished respjson.Field
		Collaborative                        respjson.Field
		Description                          respjson.Field
		ExternalURLs                         respjson.Field
		Followers                            respjson.Field
		Href                                 respjson.Field
		Images                               respjson.Field
		Name                                 respjson.Field
		Owner                                respjson.Field
		SnapshotID                           respjson.Field
		Tracks                               respjson.Field
		Type                                 respjson.Field
		Uri                                  respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserPlaylistNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserPlaylistNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user who owns the playlist
type UserPlaylistNewResponseOwner struct {
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
func (r UserPlaylistNewResponseOwner) RawJSON() string { return r.JSON.raw }
func (r *UserPlaylistNewResponseOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The tracks of the playlist.
type UserPlaylistNewResponseTracks struct {
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
func (r UserPlaylistNewResponseTracks) RawJSON() string { return r.JSON.raw }
func (r *UserPlaylistNewResponseTracks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPlaylistNewParams struct {
	// The name for the new playlist, for example `"Your Coolest Playlist"`. This name
	// does not need to be unique; a user may have several playlists with the same
	// name.
	Name string `json:"name,required"`
	// Defaults to `true`. The playlist's public/private status (if it should be added
	// to the user's profile or not): `true` the playlist will be public, `false` the
	// playlist will be private. To be able to create private playlists, the user must
	// have granted the `playlist-modify-private`
	// [scope](/documentation/web-api/concepts/scopes/#list-of-scopes). For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	PathsRequestBodyContentApplicationJsonSchemaPropertiesPublished param.Opt[bool] `json:"$.paths['*'].*.requestBody.content['application/json'].schema.properties.published,omitzero"`
	// Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to
	// create a collaborative playlist you must also set `public` to `false`. To create
	// collaborative playlists you must have granted `playlist-modify-private` and
	// `playlist-modify-public`
	// [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._
	Collaborative param.Opt[bool] `json:"collaborative,omitzero"`
	// value for playlist description as displayed in Spotify Clients and in the Web
	// API.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r UserPlaylistNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserPlaylistNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserPlaylistNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserPlaylistListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first playlist to return. Default: 0 (the first object).
	// Maximum offset: 100.000\. Use with `limit` to get the next set of playlists.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserPlaylistListParams]'s query parameters as `url.Values`.
func (r UserPlaylistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
