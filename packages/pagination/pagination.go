// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"net/url"

	"github.com/stainless-sdks/spotted-go/internal/apijson"
	"github.com/stainless-sdks/spotted-go/internal/requestconfig"
	"github.com/stainless-sdks/spotted-go/option"
	"github.com/stainless-sdks/spotted-go/packages/param"
	"github.com/stainless-sdks/spotted-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorURLPage[T any] struct {
	Next  string `json:"next"`
	Items []T    `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorURLPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorURLPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorURLPage[T]) GetNextPage() (res *CursorURLPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	rawURL := r.Next
	if len(rawURL) == 0 {
		return nil, nil
	}

	cfg := r.cfg.Clone(r.cfg.Context)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	for key, values := range parsedURL.Query() {
		err = cfg.Apply(option.WithQuery(key, values[0]))
		if err != nil {
			return nil, err
		}
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorURLPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorURLPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorURLPageAutoPager[T any] struct {
	page *CursorURLPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorURLPageAutoPager[T any](page *CursorURLPage[T], err error) *CursorURLPageAutoPager[T] {
	return &CursorURLPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorURLPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorURLPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorURLPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorURLPageAutoPager[T]) Index() int {
	return r.run
}

type AlbumsCursorURLPage[T any] struct {
	Next  string `json:"next"`
	Items []T    `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r AlbumsCursorURLPage[T]) RawJSON() string { return r.JSON.raw }
func (r *AlbumsCursorURLPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AlbumsCursorURLPage[T]) GetNextPage() (res *AlbumsCursorURLPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	rawURL := r.Next
	if len(rawURL) == 0 {
		return nil, nil
	}

	cfg := r.cfg.Clone(r.cfg.Context)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	for key, values := range parsedURL.Query() {
		err = cfg.Apply(option.WithQuery(key, values[0]))
		if err != nil {
			return nil, err
		}
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AlbumsCursorURLPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AlbumsCursorURLPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AlbumsCursorURLPageAutoPager[T any] struct {
	page *AlbumsCursorURLPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewAlbumsCursorURLPageAutoPager[T any](page *AlbumsCursorURLPage[T], err error) *AlbumsCursorURLPageAutoPager[T] {
	return &AlbumsCursorURLPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AlbumsCursorURLPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AlbumsCursorURLPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *AlbumsCursorURLPageAutoPager[T]) Err() error {
	return r.err
}

func (r *AlbumsCursorURLPageAutoPager[T]) Index() int {
	return r.run
}
