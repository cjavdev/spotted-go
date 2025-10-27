// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/spotted-go/internal/apijson"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/respjson"
	"github.com/stainless-sdks/spotted-go/shared"
)

// MeService contains methods and other services that help with interacting with
// the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options    []option.RequestOption
	Audiobooks MeAudiobookService
	Playlists  MePlaylistService
	Top        MeTopService
	Albums     MeAlbumService
	Tracks     MeTrackService
	Episodes   MeEpisodeService
	Shows      MeShowService
	Following  MeFollowingService
	Player     MePlayerService
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.Options = opts
	r.Audiobooks = NewMeAudiobookService(opts...)
	r.Playlists = NewMePlaylistService(opts...)
	r.Top = NewMeTopService(opts...)
	r.Albums = NewMeAlbumService(opts...)
	r.Tracks = NewMeTrackService(opts...)
	r.Episodes = NewMeEpisodeService(opts...)
	r.Shows = NewMeShowService(opts...)
	r.Following = NewMeFollowingService(opts...)
	r.Player = NewMePlayerService(opts...)
	return
}

// Get detailed profile information about the current user (including the current
// user's username).
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MeGetResponse struct {
	// The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for the
	// user.
	ID string `json:"id"`
	// The country of the user, as set in the user's account profile. An
	// [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// _This field is only available when the current user has granted access to the
	// [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes)
	// scope._
	Country string `json:"country"`
	// The name displayed on the user's profile. `null` if not available.
	DisplayName string `json:"display_name"`
	// The user's email address, as entered by the user when creating their account.
	// _**Important!** This email address is unverified; there is no proof that it
	// actually belongs to the user._ _This field is only available when the current
	// user has granted access to the
	// [user-read-email](/documentation/web-api/concepts/scopes/#list-of-scopes)
	// scope._
	Email string `json:"email"`
	// The user's explicit content settings. _This field is only available when the
	// current user has granted access to the
	// [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes)
	// scope._
	ExplicitContent MeGetResponseExplicitContent `json:"explicit_content"`
	// Known external URLs for this user.
	ExternalURLs shared.ExternalURLObject `json:"external_urls"`
	// Information about the followers of the user.
	Followers shared.FollowersObject `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The user's profile image.
	Images []shared.ImageObject `json:"images"`
	// The user's Spotify subscription level: "premium", "free", etc. (The subscription
	// level "open" can be considered the same as "free".) _This field is only
	// available when the current user has granted access to the
	// [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes)
	// scope._
	Product string `json:"product"`
	// The object type: "user"
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the
	// user.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Country         respjson.Field
		DisplayName     respjson.Field
		Email           respjson.Field
		ExplicitContent respjson.Field
		ExternalURLs    respjson.Field
		Followers       respjson.Field
		Href            respjson.Field
		Images          respjson.Field
		Product         respjson.Field
		Type            respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user's explicit content settings. _This field is only available when the
// current user has granted access to the
// [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes)
// scope._
type MeGetResponseExplicitContent struct {
	// When `true`, indicates that explicit content should not be played.
	FilterEnabled bool `json:"filter_enabled"`
	// When `true`, indicates that the explicit content setting is locked and can't be
	// changed by the user.
	FilterLocked bool `json:"filter_locked"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FilterEnabled respjson.Field
		FilterLocked  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponseExplicitContent) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponseExplicitContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
