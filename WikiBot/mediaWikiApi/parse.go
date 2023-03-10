package mwApi

import (
	"golang.org/x/exp/slices"
)

/*
	Parse is the parameter set for parsing a page of the wiki.

It implements the QueryMapper interface.

It defaults to:

	pase action
	content model of text
	retrieving the properties for wikitext and categories.
	It does NOT have a PageId/Title/Page set and so will need this set before being called.
*/
type Parse struct {
	action       string `default:"parse"`
	PageId       string
	Title        string
	Page         string
	ContentModel string `default:"text"`
	Prop         string `default:"wikitext|categories|templates"` // A pipe separated list of properties. see https://www.mediawiki.org/wiki/API:Parsing_wikitext#parse
}

/*
	Parse.Map() outputs a parameter map for the Parse Query.

It will only allow one of the following: PageId, Title, Page.
It will take the first it encounters in that order, discarding the others.
*/
func (pa Parse) Map() map[string]string {
	fields, output := PrepMap(Parse{})

	// these values are mutually exclusive to one another. If more than one are set we don't want
	// them all sent to the api parameters.
	contentIdentifiers := []string{"PageId", "Page", "Title"}

	alreadyHaveContentIdentifier := false

	for _, field := range fields {

		if field.Name == "Parse" {
			continue
		}

		// since reflection returns the slice of fields in the order in which they are defined on the struct
		// we can use that to define priority as well
		if slices.Contains(contentIdentifiers, field.Name) {

			if alreadyHaveContentIdentifier {
				continue
			}

			alreadyHaveContentIdentifier = !isFieldBlank(pa, field)
		}

		GetKeyAndValue(pa, field, output)

	}

	//ok := verifyDependantKeysPresent(output, parseKeyDependencies)

	return output

}
