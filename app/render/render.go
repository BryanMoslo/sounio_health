package render

import (
	base "cotizador_sounio_health"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/plush/v4"
	"github.com/leekchan/accounting"
)

// Engine for rendering across the app, it provides
// the base for rendering HTML, JSON, XML and other formats
// while also defining thing like the base layout.
var Engine = render.New(render.Options{
	HTMLLayout:   "application.plush.html",
	TemplatesBox: base.Templates,
	AssetsBox:    base.Assets,
	Helpers:      Helpers,
})

// Helpers available for the plush templates, there are
// some helpers that are injected by Buffalo but this is
// the list of custom Helpers.
var Helpers = map[string]interface{}{
	// partialFeeder is the helper used by the render engine
	// to find the partials that will be used, this is important
	"partialFeeder":   base.Templates.FindString,
	"toCurrency":      toCurrency,
	"floatToString":   floatToString,
	"activePathClass": activePathClass,
	"today":           today,
	"terms":           terms,
}

func toCurrency(value float64) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 0}
	return ac.FormatMoney(value)
}

func floatToString(value float64) string {
	return fmt.Sprintf("%.0f", value)
}

func activePathClass(class, basePath string, help plush.HelperContext) string {
	request := help.Value("request").(*http.Request)
	requestURL := request.URL

	exp, err := regexp.Compile(fmt.Sprintf(`^%v$`, basePath))
	if err != nil {
		return ""
	}

	if exp.Match([]byte(requestURL.String())) {
		return class
	}

	return ""
}

func today() string {
	return time.Now().Format("02/01/2006")
}

func terms() []string {
	terms := []string{}
	for i := 36; i <= 72; i++ {
		terms = append(terms, fmt.Sprintf("%v", i))
	}

	return terms
}
