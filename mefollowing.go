// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
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
	"github.com/cjavdev/spotted-go/shared/constant"
)

// MeFollowingService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeFollowingService] method instead.
type MeFollowingService struct {
	Options []option.RequestOption
}

// NewMeFollowingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeFollowingService(opts ...option.RequestOption) (r MeFollowingService) {
	r = MeFollowingService{}
	r.Options = opts
	return
}

// Get the current user's followed artists.
func (r *MeFollowingService) BulkGet(ctx context.Context, query MeFollowingBulkGetParams, opts ...option.RequestOption) (res *MeFollowingBulkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check to see if the current user is following one or more artists or other
// Spotify users.
//
// **Note:** This endpoint is deprecated. Use
// [Check User's Saved Items](/documentation/web-api/reference/check-library-contains)
// instead.
//
// Deprecated: deprecated
func (r *MeFollowingService) Check(ctx context.Context, query MeFollowingCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/following/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add the current user as a follower of one or more artists or other Spotify
// users.
//
// **Note:** This endpoint is deprecated. Use
// [Save Items to Library](/documentation/web-api/reference/save-library-items)
// instead.
//
// Deprecated: deprecated
func (r *MeFollowingService) Follow(ctx context.Context, body MeFollowingFollowParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Remove the current user as a follower of one or more artists or other Spotify
// users.
//
// **Note:** This endpoint is deprecated. Use
// [Remove Items from Library](/documentation/web-api/reference/remove-library-items)
// instead.
//
// Deprecated: deprecated
func (r *MeFollowingService) Unfollow(ctx context.Context, body MeFollowingUnfollowParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type MeFollowingBulkGetResponse struct {
	Artists MeFollowingBulkGetResponseArtists `json:"artists,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artists     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingBulkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingBulkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeFollowingBulkGetResponseArtists struct {
	// The cursors used to find the next set of items.
	Cursors MeFollowingBulkGetResponseArtistsCursors `json:"cursors"`
	// A link to the Web API endpoint returning the full result of the request.
	Href  string                `json:"href"`
	Items []shared.ArtistObject `json:"items"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int64 `json:"limit"`
	// URL to the next page of items. ( `null` if none)
	Next string `json:"next"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// The total number of items available to return.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		Href        respjson.Field
		Items       respjson.Field
		Limit       respjson.Field
		Next        respjson.Field
		Published   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingBulkGetResponseArtists) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingBulkGetResponseArtists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The cursors used to find the next set of items.
type MeFollowingBulkGetResponseArtistsCursors struct {
	// The cursor to use as key to find the next page of items.
	After string `json:"after"`
	// The cursor to use as key to find the previous page of items.
	Before string `json:"before"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published bool `json:"published"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		Published   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingBulkGetResponseArtistsCursors) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingBulkGetResponseArtistsCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeFollowingBulkGetParams struct {
	// The last artist ID retrieved from the previous request.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID type: currently only `artist` is supported.
	//
	// This field can be elided, and will marshal its zero value as "artist".
	Type constant.Artist `query:"type,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeFollowingBulkGetParams]'s query parameters as
// `url.Values`.
func (r MeFollowingBulkGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeFollowingCheckParams struct {
	// A comma-separated list of the artist or the user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids) to check. For
	// example: `ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q`. A maximum of 50
	// IDs can be sent in one request.
	IDs string `query:"ids,required" json:"-"`
	// The ID type: either `artist` or `user`.
	//
	// Any of "artist", "user".
	Type MeFollowingCheckParamsType `query:"type,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeFollowingCheckParams]'s query parameters as `url.Values`.
func (r MeFollowingCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The ID type: either `artist` or `user`.
type MeFollowingCheckParamsType string

const (
	MeFollowingCheckParamsTypeArtist MeFollowingCheckParamsType = "artist"
	MeFollowingCheckParamsTypeUser   MeFollowingCheckParamsType = "user"
)

type MeFollowingFollowParams struct {
	// A JSON array of the artist or user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `{ids:["74ASZWbe4lXaubB36ztrGX", "08td7MxkoHQkXnWAYD8d6Q"]}`. A maximum of 50
	// IDs can be sent in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	IDs []string `json:"ids,omitzero,required"`
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published param.Opt[bool] `json:"published,omitzero"`
	paramObj
}

func (r MeFollowingFollowParams) MarshalJSON() (data []byte, err error) {
	type shadow MeFollowingFollowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeFollowingFollowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeFollowingUnfollowParams struct {
	// The playlist's public/private status (if it should be added to the user's
	// profile or not): `true` the playlist will be public, `false` the playlist will
	// be private, `null` the playlist status is not relevant. For more about
	// public/private status, see
	// [Working with Playlists](/documentation/web-api/concepts/playlists)
	Published param.Opt[bool] `json:"published,omitzero"`
	// A JSON array of the artist or user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `{ids:["74ASZWbe4lXaubB36ztrGX", "08td7MxkoHQkXnWAYD8d6Q"]}`. A maximum of 50
	// IDs can be sent in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	IDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeFollowingUnfollowParams) MarshalJSON() (data []byte, err error) {
	type shadow MeFollowingUnfollowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeFollowingUnfollowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
