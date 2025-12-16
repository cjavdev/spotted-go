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
)

// PlaylistFollowerService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistFollowerService] method instead.
type PlaylistFollowerService struct {
	Options []option.RequestOption
}

// NewPlaylistFollowerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPlaylistFollowerService(opts ...option.RequestOption) (r PlaylistFollowerService) {
	r = PlaylistFollowerService{}
	r.Options = opts
	return
}

// Check to see if the current user is following a specified playlist.
func (r *PlaylistFollowerService) Check(ctx context.Context, playlistID string, query PlaylistFollowerCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/followers/contains", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add the current user as a follower of a playlist.
func (r *PlaylistFollowerService) Follow(ctx context.Context, playlistID string, body PlaylistFollowerFollowParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/followers", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Remove the current user as a follower of a playlist.
func (r *PlaylistFollowerService) Unfollow(ctx context.Context, playlistID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/followers", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PlaylistFollowerCheckParams struct {
	// **Deprecated** A single item list containing current user's
	// [Spotify Username](/documentation/web-api/concepts/spotify-uris-ids). Maximum: 1
	// id.
	IDs param.Opt[string] `query:"ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PlaylistFollowerCheckParams]'s query parameters as
// `url.Values`.
func (r PlaylistFollowerCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlaylistFollowerFollowParams struct {
	// Defaults to `true`. If `true` the playlist will be included in user's public
	// playlists (added to profile), if `false` it will remain private. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Public param.Opt[bool] `json:"public,omitzero"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published param.Opt[bool] `json:"published,omitzero"`
	paramObj
}

func (r PlaylistFollowerFollowParams) MarshalJSON() (data []byte, err error) {
	type shadow PlaylistFollowerFollowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlaylistFollowerFollowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
