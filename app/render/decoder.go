package render

import (
	"cotizador_sounio_health/pkg/formam"
	"html"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"

	"github.com/pkg/errors"
)

var decoder *formam.Decoder

func init() {
	decoder = formam.NewDecoder(&formam.DecoderOptions{
		TagName:           "form",
		IgnoreUnknownKeys: true,
	})

	decoder.RegisterCustomType(intDecoder, []interface{}{int(0)}, nil)
	decoder.RegisterCustomType(int64Decoder, []interface{}{int64(0)}, nil)
	decoder.RegisterCustomType(float64Decoder, []interface{}{float64(0)}, nil)
	decoder.RegisterCustomType(nullsFloat64Decoder, []interface{}{nulls.Float64{}}, nil)
	decoder.RegisterCustomType(nullsStringDecoder, []interface{}{nulls.String{}}, nil)
	decoder.RegisterCustomType(nullsIntDecoder, []interface{}{nulls.Int{}}, nil)
	decoder.RegisterCustomType(nullsBoolDecoder, []interface{}{nulls.Bool{}}, nil)
	decoder.RegisterCustomType(uuidDecoder, []interface{}{uuid.UUID{}}, nil)
}

func intDecoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return int(0), nil
	}

	val, err := strconv.Atoi(vals[0])
	if err != nil {
		return int(0), nil
	}

	return val, nil
}

func int64Decoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return int64(0), nil
	}

	val, err := strconv.ParseInt(vals[0], 10, 64)
	if err != nil {
		return int64(0), nil
	}

	return val, nil
}

func float64Decoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return float64(0), nil
	}

	val, err := strconv.ParseFloat(vals[0], 64)
	if err != nil {
		return float64(0), nil
	}

	return val, nil
}

func nullsFloat64Decoder(vals []string) (interface{}, error) {
	if vals[0] == "" || vals[0] == "." {
		return nulls.Float64{}, nil
	}

	val, err := strconv.ParseFloat(strings.ReplaceAll(vals[0], ",", ""), 64)
	return nulls.NewFloat64(val), err
}

func nullsStringDecoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return nulls.NewString(""), nil
	}

	return nulls.NewString(html.EscapeString(vals[0])), nil
}

func nullsIntDecoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return nulls.Int{}, nil
	}

	val, convErr := strconv.Atoi(vals[0])
	if convErr != nil {
		return nulls.Int{}, convErr
	}

	return nulls.NewInt(val), nil
}

func nullsBoolDecoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return nulls.Bool{}, nil
	}

	val, convErr := strconv.ParseBool(vals[0])
	if convErr != nil {
		return nulls.Bool{}, convErr
	}

	return nulls.NewBool(val), nil
}

func uuidDecoder(vals []string) (interface{}, error) {
	if vals[0] == "" {
		return uuid.Nil, nil
	}

	return uuid.FromStringOrNil(vals[0]), nil
}

// HTMLBinding is our custom function to parse requests into structs.
func HTMLBinding(req *http.Request, i interface{}) error {
	err := req.ParseForm()
	if err != nil {
		return errors.WithStack(err)
	}

	rx := structs.New(i)
	for _, f := range rx.Fields() {
		if f.Kind() != reflect.Bool {
			continue
		}

		if req.Form.Get(f.Name()) == "" {
			req.Form.Set(f.Name(), "false")
		}
	}

	if err := decoder.Decode(req.Form, i); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
