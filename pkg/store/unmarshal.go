// Copyright © 2018 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

package store

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/TheThingsNetwork/ttn/pkg/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/mitchellh/mapstructure"
)

// slicify recursively replaces (sub-)maps in m by slices.
func slicify(m map[string]interface{}) interface{} {
	for k, v := range m {
		sm, ok := v.(map[string]interface{})
		if ok {
			v = slicify(sm)
			m[k] = v
		}
	}

	sl := make([]interface{}, len(m))
	for k, v := range m {
		i, err := strconv.Atoi(k)
		if err != nil {
			return m
		}
		if len(sl) <= i {
			sl = append(sl, make([]interface{}, 1+i-len(sl))...)
		}
		sl[i] = v
	}
	return sl
}

func unflattened(m map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(m))
	for k, v := range m {
		skeys := strings.Split(k, Separator)
		parent := out
		for _, sk := range skeys[:len(skeys)-1] {
			sm, ok := parent[sk]
			if !ok {
				sm = make(map[string]interface{})
				parent[sk] = sm
			}
			parent = sm.(map[string]interface{})
		}
		parent[skeys[len(skeys)-1]] = v
	}
	return out
}

// MapUnmarshaler is the interface implemented by an object that can
// unmarshal a map[string]interface{} representation of itself.
//
// UnmarshalMap must be able to decode the form generated by MarshalMap.
// UnmarshalMap must deep copy the data if it wishes to retain the data
// after returning.
type MapUnmarshaler interface {
	UnmarshalMap(map[string]interface{}) error
}

// UnmarshalMap parses the map-encoded data and stores the result
// in the value pointed to by v.
//
// UnmarshalMap uses the inverse of the encodings that
// Marshal uses.
func UnmarshalMap(m map[string]interface{}, v interface{}) error {
	m = unflattened(m)
	switch t := v.(type) {
	case MapUnmarshaler:
		return t.UnmarshalMap(m)
	default:
		if len(m) == 0 {
			return nil
		}
		dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			ZeroFields:       true,
			Result:           v,
		})
		if err != nil {
			panic(errors.NewWithCause("Failed to intialize decoder", err))
		}
		return dec.Decode(slicify(m))
	}
}

func typeByFlatName(typ reflect.Type, name string) (reflect.Type, bool) {
	if name == "" {
		panic(errors.New("Empty name specified"))
	}
	for _, name := range strings.Split(name, Separator) {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		switch typ.Kind() {
		case reflect.Struct:
			f, ok := typ.FieldByName(name)
			if !ok {
				return nil, false
			}
			typ = f.Type
		case reflect.Slice, reflect.Array:
			typ = typ.Elem()
		case reflect.Map:
			typ = typ.Key()
		default:
			return nil, false
		}
	}
	return typ, true
}

func bytesToType(b []byte, typ reflect.Type) (interface{}, error) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	enc := Encoding(b[0])
	b = b[1:]
	if len(b) == 0 {
		return reflect.New(typ).Elem().Interface(), nil
	}

	switch enc {
	case RawEncoding:
		switch k := typ.Kind(); k {
		case reflect.String:
			return string(b), nil
		case reflect.Bool:
			return strconv.ParseBool(string(b))
		case reflect.Int:
			return strconv.ParseInt(string(b), 10, 64)
		case reflect.Int8:
			return strconv.ParseInt(string(b), 10, 8)
		case reflect.Int16:
			return strconv.ParseInt(string(b), 10, 16)
		case reflect.Int32:
			return strconv.ParseInt(string(b), 10, 32)
		case reflect.Int64:
			return strconv.ParseInt(string(b), 10, 64)
		case reflect.Uint:
			return strconv.ParseUint(string(b), 10, 64)
		case reflect.Uint8:
			return strconv.ParseUint(string(b), 10, 8)
		case reflect.Uint16:
			return strconv.ParseUint(string(b), 10, 16)
		case reflect.Uint32:
			return strconv.ParseUint(string(b), 10, 32)
		case reflect.Uint64:
			return strconv.ParseUint(string(b), 10, 64)
		case reflect.Float32:
			return strconv.ParseFloat(string(b), 32)
		case reflect.Float64:
			return strconv.ParseFloat(string(b), 64)
		case reflect.Slice, reflect.Array:
			elem := typ.Elem()
			if elem.Kind() == reflect.Uint8 {
				// Handle byte slices/arrays directly
				if k == reflect.Slice {
					return b, nil
				}
				rv := reflect.Indirect(reflect.New(typ))
				for i := 0; i < rv.Len(); i++ {
					rv.Index(i).SetUint(uint64(b[i]))
				}
				return rv.Interface(), nil
			}
		}
		return nil, errors.Errorf("can not decode raw bytes to value of type %s", typ)
	case BinaryEncoding:
		v, ok := reflect.New(typ).Interface().(encoding.BinaryUnmarshaler)
		if !ok {
			var expected encoding.BinaryUnmarshaler
			return nil, errors.Errorf("expected %s to implement %T", typ, expected)
		}
		return v, v.UnmarshalBinary(b)
	case TextEncoding:
		v, ok := reflect.New(typ).Interface().(encoding.TextUnmarshaler)
		if !ok {
			var expected encoding.TextUnmarshaler
			return nil, errors.Errorf("expected %s to implement %T", typ, expected)
		}
		return v, v.UnmarshalText(b)
	case ProtoEncoding:
		v, ok := reflect.New(typ).Interface().(proto.Unmarshaler)
		if !ok {
			var expected proto.Unmarshaler
			return nil, errors.Errorf("expected %s to implement %T", typ, expected)
		}
		return v, v.Unmarshal(b)
	case JSONEncoding:
		v, ok := reflect.New(typ).Interface().(json.Unmarshaler)
		if !ok {
			var expected json.Unmarshaler
			return nil, errors.Errorf("expected %s to implement %T", typ, expected)
		}
		return v, v.UnmarshalJSON(b)
	case UnknownEncoding:
		return nil, errors.New("value is encoded using unknown encoding")
	default:
		return nil, errors.New("invalid data")
	}
}

// UnmarshalByteMap parses the byte map-encoded data and stores the result
// in the value pointed to by v.
//
// UnmarshalByteMap uses the inverse of the encodings that
// MarshalByteMap uses.
func UnmarshalByteMap(bm map[string][]byte, v interface{}) error {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	im := make(map[string]interface{}, len(bm))
	switch typ.Kind() {
	case reflect.Struct:
		for k, bv := range bm {
			it, ok := typeByFlatName(typ, k)
			if !ok {
				return errors.Errorf("field %s does not exist on type specified", k)
			}
			iv, err := bytesToType(bv, it)
			if err != nil {
				return err
			}
			im[k] = iv
		}
	default:
		panic(errors.Errorf("UnmarshalByteMap: %s is not supported yet", typ))
	}
	return UnmarshalMap(im, v)
}
