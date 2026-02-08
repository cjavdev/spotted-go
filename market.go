// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted

import (
	"context"
	"net/http"
	"slices"

	"github.com/cjavdev/spotted-go/internal/apijson"
	"github.com/cjavdev/spotted-go/internal/requestconfig"
	"github.com/cjavdev/spotted-go/option"
	"github.com/cjavdev/spotted-go/packages/respjson"
)

// MarketService contains methods and other services that help with interacting
// with the spotted API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMarketService] method instead.
type MarketService struct {
	Options []option.RequestOption
}

// NewMarketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMarketService(opts ...option.RequestOption) (r MarketService) {
	r = MarketService{}
	r.Options = opts
	return
}

// Get the list of markets where Spotify is available.
//
// Deprecated: deprecated
func (r *MarketService) List(ctx context.Context, opts ...option.RequestOption) (res *MarketListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "markets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MarketListResponse struct {
	Markets []string `json:"markets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MarketListResponse) RawJSON() string { return r.JSON.raw }
func (r *MarketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
