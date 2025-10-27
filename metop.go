// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/spotted-go/internal/apiquery"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/pagination"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-go/shared"
)

// MeTopService contains methods and other services that help with interacting with
// the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeTopService] method instead.
type MeTopService struct {
	Options []option.RequestOption
}

// NewMeTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeTopService(opts ...option.RequestOption) (r MeTopService) {
	r = MeTopService{}
	r.Options = opts
	return
}

// Get the current user's top artists based on calculated affinity.
func (r *MeTopService) ListTopArtists(ctx context.Context, query MeTopListTopArtistsParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.ArtistObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/top/artists"
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

// Get the current user's top artists based on calculated affinity.
func (r *MeTopService) ListTopArtistsAutoPaging(ctx context.Context, query MeTopListTopArtistsParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.ArtistObject] {
	return pagination.NewCursorURLPageAutoPager(r.ListTopArtists(ctx, query, opts...))
}

// Get the current user's top tracks based on calculated affinity.
func (r *MeTopService) ListTopTracks(ctx context.Context, query MeTopListTopTracksParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[shared.TrackObject], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/top/tracks"
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

// Get the current user's top tracks based on calculated affinity.
func (r *MeTopService) ListTopTracksAutoPaging(ctx context.Context, query MeTopListTopTracksParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[shared.TrackObject] {
	return pagination.NewCursorURLPageAutoPager(r.ListTopTracks(ctx, query, opts...))
}

type MeTopListTopArtistsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Over what time frame the affinities are computed. Valid values: `long_term`
	// (calculated from ~1 year of data and including all new data as it becomes
	// available), `medium_term` (approximately last 6 months), `short_term`
	// (approximately last 4 weeks). Default: `medium_term`
	TimeRange param.Opt[string] `query:"time_range,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeTopListTopArtistsParams]'s query parameters as
// `url.Values`.
func (r MeTopListTopArtistsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeTopListTopTracksParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Over what time frame the affinities are computed. Valid values: `long_term`
	// (calculated from ~1 year of data and including all new data as it becomes
	// available), `medium_term` (approximately last 6 months), `short_term`
	// (approximately last 4 weeks). Default: `medium_term`
	TimeRange param.Opt[string] `query:"time_range,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeTopListTopTracksParams]'s query parameters as
// `url.Values`.
func (r MeTopListTopTracksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
