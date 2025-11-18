// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/pagination"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/shared"
)

// MePlaylistService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMePlaylistService] method instead.
type MePlaylistService struct {
	Options []option.RequestOption
}

// NewMePlaylistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMePlaylistService(opts ...option.RequestOption) (r MePlaylistService) {
	r = MePlaylistService{}
	r.Options = opts
	return
}

// Get a list of the playlists owned or followed by the current Spotify user.
func (r *MePlaylistService) List(ctx context.Context, query MePlaylistListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.SimplifiedPlaylistObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/playlists"
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

// Get a list of the playlists owned or followed by the current Spotify user.
func (r *MePlaylistService) ListAutoPaging(ctx context.Context, query MePlaylistListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.SimplifiedPlaylistObject] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

type MePlaylistListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// 'The index of the first playlist to return. Default: 0 (the first object).
	// Maximum offset: 100.000\. Use with `limit` to get the next set of playlists.'
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MePlaylistListParams]'s query parameters as `url.Values`.
func (r MePlaylistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
