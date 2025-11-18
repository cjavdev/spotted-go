// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/apiquery"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/pagination"
	"github.com/cjavdev/spotted-go/packages/param"
	"github.com/cjavdev/spotted-go/packages/respjson"
	"github.com/cjavdev/spotted-go/shared"
)

// MeAudiobookService contains methods and other services that help with
// interacting with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeAudiobookService] method instead.
type MeAudiobookService struct {
	Options []option.RequestOption
}

// NewMeAudiobookService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeAudiobookService(opts ...option.RequestOption) (r MeAudiobookService) {
	r = MeAudiobookService{}
	r.Options = opts
	return
}

// Get a list of the audiobooks saved in the current Spotify user's 'Your Music'
// library.
func (r *MeAudiobookService) List(ctx context.Context, query MeAudiobookListParams, opts ...option.RequestOption) (res *pagination.CursorURLPage[MeAudiobookListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "me/audiobooks"
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

// Get a list of the audiobooks saved in the current Spotify user's 'Your Music'
// library.
func (r *MeAudiobookService) ListAutoPaging(ctx context.Context, query MeAudiobookListParams, opts ...option.RequestOption) *pagination.CursorURLPageAutoPager[MeAudiobookListResponse] {
	return pagination.NewCursorURLPageAutoPager(r.List(ctx, query, opts...))
}

// Check if one or more audiobooks are already saved in the current Spotify user's
// library.
func (r *MeAudiobookService) Check(ctx context.Context, query MeAudiobookCheckParams, opts ...option.RequestOption) (res *[]bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "me/audiobooks/contains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove one or more audiobooks from the Spotify user's library.
func (r *MeAudiobookService) Remove(ctx context.Context, body MeAudiobookRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/audiobooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Save one or more audiobooks to the current Spotify user's library.
func (r *MeAudiobookService) Save(ctx context.Context, body MeAudiobookSaveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "me/audiobooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

type MeAudiobookListResponse struct {
	// The date and time the audiobook was saved Timestamps are returned in ISO 8601
	// format as Coordinated Universal Time (UTC) with a zero offset:
	// YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an
	// album release), an additional field indicates the precision; see for example,
	// release_date in an album object.
	AddedAt time.Time `json:"added_at" format:"date-time"`
	// Information about the audiobook.
	Audiobook MeAudiobookListResponseAudiobook `json:"audiobook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedAt     respjson.Field
		Audiobook   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeAudiobookListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeAudiobookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the audiobook.
type MeAudiobookListResponseAudiobook struct {
	// The chapters of the audiobook.
	Chapters MeAudiobookListResponseAudiobookChapters `json:"chapters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	shared.AudiobookBase
}

// Returns the unmodified JSON received from the API
func (r MeAudiobookListResponseAudiobook) RawJSON() string { return r.JSON.raw }
func (r *MeAudiobookListResponseAudiobook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The chapters of the audiobook.
type MeAudiobookListResponseAudiobookChapters struct {
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
	Total int64                     `json:"total,required"`
	Items []SimplifiedChapterObject `json:"items"`
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
func (r MeAudiobookListResponseAudiobookChapters) RawJSON() string { return r.JSON.raw }
func (r *MeAudiobookListResponseAudiobookChapters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeAudiobookListParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index of the first item to return. Default: 0 (the first item). Use with
	// limit to get the next set of items.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeAudiobookListParams]'s query parameters as `url.Values`.
func (r MeAudiobookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAudiobookCheckParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeAudiobookCheckParams]'s query parameters as `url.Values`.
func (r MeAudiobookCheckParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAudiobookRemoveParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeAudiobookRemoveParams]'s query parameters as
// `url.Values`.
func (r MeAudiobookRemoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAudiobookSaveParams struct {
	// A comma-separated list of the
	// [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example:
	// `ids=18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ`. Maximum: 50 IDs.
	IDs string `query:"ids,required" json:"-"`
	paramObj
}

// URLQuery serializes [MeAudiobookSaveParams]'s query parameters as `url.Values`.
func (r MeAudiobookSaveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
