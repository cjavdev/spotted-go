// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// UserService contains methods and other services that help with interacting with
// the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options   []option.RequestOption
	Playlists UserPlaylistService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	r.Playlists = NewUserPlaylistService(opts...)
	return
}

// Get public profile information about a Spotify user.
func (r *UserService) GetProfile(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserGetProfileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserGetProfileResponse struct {
	// The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this
	// user.
	ID string `json:"id"`
	// The name displayed on the user's profile. `null` if not available.
	DisplayName string `json:"display_name,nullable"`
	// Known public external URLs for this user.
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	// Information about the followers of this user.
	Followers shared.FollowersObject `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The user's profile image.
	Images []shared.ImageObject `json:"images"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The object type.
	//
	// Any of "user".
	Type UserGetProfileResponseType `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this
	// user.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayName  respjson.Field
		ExternalURLs respjson.Field
		Followers    respjson.Field
		Href         respjson.Field
		Images       respjson.Field
		Published    respjson.Field
		Type         respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetProfileResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetProfileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object type.
type UserGetProfileResponseType string

const (
	UserGetProfileResponseTypeUser UserGetProfileResponseType = "user"
)
