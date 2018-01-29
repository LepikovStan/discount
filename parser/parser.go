package parser

import (
	"io"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
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
	}
}

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
		result[0] = strings.TrimSpace(doc.Find(selectors["name"]).Text())
		result[1] = strings.TrimSpace(doc.Find(selectors["price"]).Text())
		result[2] = strings.TrimSpace(doc.Find(selectors["oldPrice"]).Text())
	}

	return result, nil
}

// New function initialize new Parser instance and return pointer to it
func New() *parser {
	p := new(parser)
	p.init()
	return p
}
