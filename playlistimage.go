// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/spotted-go/internal/encoding/json"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/shared"
)

// PlaylistImageService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlaylistImageService] method instead.
type PlaylistImageService struct {
	Options []option.RequestOption
}

// NewPlaylistImageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlaylistImageService(opts ...option.RequestOption) (r PlaylistImageService) {
	r = PlaylistImageService{}
	r.Options = opts
	return
}

// Replace the image used to represent a specific playlist.
func (r *PlaylistImageService) Update(ctx context.Context, playlistID string, body PlaylistImageUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/images", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get the current image associated with a specific playlist.
func (r *PlaylistImageService) List(ctx context.Context, playlistID string, opts ...option.RequestOption) (res *[]shared.ImageObject, err error) {
	opts = slices.Concat(r.Options, opts)
	if playlistID == "" {
		err = errors.New("missing required playlist_id parameter")
		return
	}
	path := fmt.Sprintf("playlists/%s/images", playlistID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PlaylistImageUpdateParams struct {
	// Base64 encoded JPEG image data, maximum payload size is 256 KB.
	Body string
	paramObj
}

func (r PlaylistImageUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PlaylistImageUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
