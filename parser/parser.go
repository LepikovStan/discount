package parser

import (
	"io"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"regexp"
)

// Parser is the type, contains the basic Parse method
type parser struct {
	selectors map[string]map[string]string
}

func (p *parser) init() {
	p.selectors = map[string]map[string]string{
		"detmir": map[string]string{
			"name": "#product_details_name",
			"oldPrice": ".products_item_c .product_card__price .old_price",
			"price": "#priceId",
		},
		"akusherstvo":   map[string]string{
			"name": ".title-page",
			"oldPrice": ".specialInfoTovarInner .price",
			"price": "#tovar_price_20430",
		},
		"ozon":   map[string]string{
			"name": ".bItemName",
			"oldPrice": ".bPriceDicsount",
			"price": ".eOzonPrice_main",
		},
		"ulmart":   map[string]string{
			"name": ".b-details__name",
			"oldPrice": "",
			"price": ".b-details__price",
		},
		"dochkisinochki":   map[string]string{
			"name": "#nameProduct",
			"oldPrice": "#ProductPrice",
			"price": "#ProductMainPrice",
		},
		"mytoys":   map[string]string{
			"name": "h1",
			"oldPrice": "",
			"price": ".price--normal",
		},
		"helptomama":   map[string]string{
			"name": "#offer_name",
			"oldPrice": "#currentOfferPrice .ElementOldPrice",
			"price": "#currentOfferPrice .ElementCurrentPrice",
		},
	}
}

var reg, err = regexp.Compile("[^0-9]+")

// Parse parsing html and find all links on the page
// return list of finded links or error
func (p *parser) Parse(body io.Reader, ptype string) ([3]string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	result := [3]string{}

	if err != nil {
		fmt.Println("Parser error", err)
		return [3]string{}, err
	}

	if selectors, ok := p.selectors[ptype]; ok != false {
		result[0] = "Moony XL Boys 12-17кг 38 шт"
		result[1] = reg.ReplaceAllString(strings.TrimSpace(doc.Find(selectors["price"]).Text()), "")
		result[2] = reg.ReplaceAllString(strings.TrimSpace(doc.Find(selectors["oldPrice"]).Text()), "")

	}

	return result, nil
}

// New function initialize new Parser instance and return pointer to it
func New() *parser {
	p := new(parser)
	p.init()
	return p
}
