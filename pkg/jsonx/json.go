package jsonx

import (
    "encoding/json"
    "io"
)

func Decode(r io.Reader, v any) error {
    dec := json.NewDecoder(r)
    dec.DisallowUnknownFields()
    return dec.Decode(v)
}
