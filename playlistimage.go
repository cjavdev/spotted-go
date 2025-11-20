// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/shared"
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
