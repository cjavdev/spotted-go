// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/spotted-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Album string     // Always "album"
type Artist string    // Always "artist"
type Audiobook string // Always "audiobook"
type Episode string   // Always "episode"
type Show string      // Always "show"

func (c Album) Default() Album         { return "album" }
func (c Artist) Default() Artist       { return "artist" }
func (c Audiobook) Default() Audiobook { return "audiobook" }
func (c Episode) Default() Episode     { return "episode" }
func (c Show) Default() Show           { return "show" }

func (c Album) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Artist) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Audiobook) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Episode) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Show) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
