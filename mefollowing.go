// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/spotted-go/internal/apijson"
	"github.com/stainless-sdks/spotted-go/internal/apiquery"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-go/packages/respjson"
	"github.com/stainless-sdks/spotted-go/shared"
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
func (r *MeFollowingService) List(ctx context.Context, query MeFollowingListParams, opts ...option.RequestOption) (res *MeFollowingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check to see if the current user is following one or more artists or other
// Spotify users.
func (r *MeFollowingService) Check(ctx context.Context, query MeFollowingCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/following/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add the current user as a follower of one or more artists or other Spotify
// users.
func (r *MeFollowingService) Follow(ctx context.Context, params MeFollowingFollowParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Remove the current user as a follower of one or more artists or other Spotify
// users.
func (r *MeFollowingService) Unfollow(ctx context.Context, params MeFollowingUnfollowParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/following"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

type MeFollowingListResponse struct {
	Artists MeFollowingListResponseArtists `json:"artists,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artists     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeFollowingListResponseArtists struct {
	// The cursors used to find the next set of items.
	Cursors MeFollowingListResponseArtistsCursors `json:"cursors"`
	// A link to the Web API endpoint returning the full result of the request.
	Href  string                `json:"href"`
	Items []shared.ArtistObject `json:"items"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int64 `json:"limit"`
	// URL to the next page of items. ( `null` if none)
	Next string `json:"next"`
	// The total number of items available to return.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		Href        respjson.Field
		Items       respjson.Field
		Limit       respjson.Field
		Next        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingListResponseArtists) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingListResponseArtists) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The cursors used to find the next set of items.
type MeFollowingListResponseArtistsCursors struct {
	// The cursor to use as key to find the next page of items.
	After string `json:"after"`
	// The cursor to use as key to find the previous page of items.
	Before string `json:"before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeFollowingListResponseArtistsCursors) RawJSON() string { return r.JSON.raw }
func (r *MeFollowingListResponseArtistsCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeFollowingListParams struct {
	// The ID type: currently only `artist` is supported.
	//
	// Any of "artist".
	Type MeFollowingListParamsType `query:"type,omitzero,required" json:"-"`
	// The last artist ID retrieved from the previous request.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeFollowingListParams]'s query parameters as `url.Values`.
func (r MeFollowingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The ID type: currently only `artist` is supported.
type MeFollowingListParamsType string

const (
	MeFollowingListParamsTypeArtist MeFollowingListParamsType = "artist"
)

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
	// A comma-separated list of the artist or the user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). A maximum of 50
	// IDs can be sent in one request.
	QueryIDs string `query:"ids,required" json:"-"`
	// The ID type.
	//
	// Any of "artist", "user".
	Type MeFollowingFollowParamsType `query:"type,omitzero,required" json:"-"`
	// A JSON array of the artist or user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `{ids:["74ASZWbe4lXaubB36ztrGX", "08td7MxkoHQkXnWAYD8d6Q"]}`. A maximum of 50
	// IDs can be sent in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	BodyIDs []string `json:"ids,omitzero,required"`
	paramObj
}

func (r MeFollowingFollowParams) MarshalJSON() (data []byte, err error) {
	type shadow MeFollowingFollowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeFollowingFollowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MeFollowingFollowParams]'s query parameters as
// `url.Values`.
func (r MeFollowingFollowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The ID type.
type MeFollowingFollowParamsType string

const (
	MeFollowingFollowParamsTypeArtist MeFollowingFollowParamsType = "artist"
	MeFollowingFollowParamsTypeUser   MeFollowingFollowParamsType = "user"
)

type MeFollowingUnfollowParams struct {
	// A comma-separated list of the artist or the user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q`. A maximum of 50 IDs can be
	// sent in one request.
	QueryIDs string `query:"ids,required" json:"-"`
	// The ID type: either `artist` or `user`.
	//
	// Any of "artist", "user".
	Type MeFollowingUnfollowParamsType `query:"type,omitzero,required" json:"-"`
	// A JSON array of the artist or user
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `{ids:["74ASZWbe4lXaubB36ztrGX", "08td7MxkoHQkXnWAYD8d6Q"]}`. A maximum of 50
	// IDs can be sent in one request. _**Note**: if the `ids` parameter is present in
	// the query string, any IDs listed here in the body will be ignored._
	BodyIDs []string `json:"ids,omitzero"`
	paramObj
}

func (r MeFollowingUnfollowParams) MarshalJSON() (data []byte, err error) {
	type shadow MeFollowingUnfollowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeFollowingUnfollowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MeFollowingUnfollowParams]'s query parameters as
// `url.Values`.
func (r MeFollowingUnfollowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The ID type: either `artist` or `user`.
type MeFollowingUnfollowParamsType string

const (
	MeFollowingUnfollowParamsTypeArtist MeFollowingUnfollowParamsType = "artist"
	MeFollowingUnfollowParamsTypeUser   MeFollowingUnfollowParamsType = "user"
)
