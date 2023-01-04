package mwApi

import (
	"reflect"
	"strconv"
	"strings"
)

type Property string

const (
	queryAction            Property = "query"
	editAction             Property = "edit"
	parseAction            Property = "parse"
	botAssert              Property = "bot"
	jsonFormat             Property = "json"
	queryListRecentChanges Property = "recentchanges"
)

type RecentChangesProp string

const (
	rcEditType RecentChangesProp = "edit"
	rcNewType  RecentChangesProp = "new"
	rcLimitMax RecentChangesProp = "max"
)

/*

	NilAction  Action = ""
	Query      Action = "query"
	EditPage   Action = "edit"
	Parse      Action = "parse"
	ShortenUrl Action = "shortenurl"
	Tag        Action = "tag"


	Json       Format = "json"
	PrettyJson Format = "jsonfm"

	Bot  Assert = "bot"
	Anon Assert = "anon"
	User Assert = "user"

	RecentChanges List = "recentchanges"

	RcEditFilter RcType = "edit"


*/

type QueryMapper interface {
	Map() map[string]string
}

type RecentChanges struct {
	action   Property
	list     Property
	Limit    int    `default:"max"`
	TopOnly  bool   `default:"false"`
	Continue bool   `default:"false"`
	Start    string // Timestamp in ISO8601 format YYYY-MM-DDThh:mm:ssZ
}

func (rc RecentChanges) Map() map[string]string {
	fields := reflect.VisibleFields(reflect.TypeOf((struct{ RecentChanges }{})))

	output := make(map[string]string, len(fields))

	for _, field := range fields {
		var value string
		var name string
		switch field.Name {

		case "RecentChanges":
			continue

		case "action":
			value = "query"
			name = field.Name
		case "list":
			value = "recentchanges"
			name = field.Name
		default:
			value = getValueOrDefault(rc, field)

			name = "rc" + strings.ToLower(field.Name)
		}

		output[name] = value

	}

	return output
}

func getValueOrDefault(q QueryMapper, field reflect.StructField) (value string) {
	fieldValue := reflect.ValueOf(q).FieldByName(field.Name)
	if fieldValue.Kind() == reflect.Int {
		value = strconv.FormatInt(fieldValue.Int(), 10)
	}
	if fieldValue.Kind() == reflect.String {
		value = fieldValue.String()
	}

	if value == "0" || value == "" {
		value, _ = field.Tag.Lookup("default")
	}
	return value
}
