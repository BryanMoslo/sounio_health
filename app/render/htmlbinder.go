package render

import (
	"cotizador_sounio_health/pkg/formam"
	"net/http"

	"github.com/gobuffalo/buffalo/binding"
)

// HTMLContentTypeBinder is in charge of binding HTML request types.
type HTMLBinder struct {
	decoder *formam.Decoder
}

// ContentTypes that will be used to identify HTML requests
func (ht HTMLBinder) ContentTypes() []string {
	return []string{
		"application/html",
		"text/html",
		"application/x-www-form-urlencoded",
		"html",
	}
}

// BinderFunc that will take care of the HTML binding
func (ht HTMLBinder) BinderFunc() binding.Binder {
	return func(req *http.Request, i interface{}) error {
		err := req.ParseForm()
		if err != nil {
			return err
		}

		if err := ht.decoder.Decode(req.Form, i); err != nil {
			return err
		}

		return nil
	}
}

func NewHTMLBinder() HTMLBinder {
	return HTMLBinder{
		decoder: decoder,
	}
}
