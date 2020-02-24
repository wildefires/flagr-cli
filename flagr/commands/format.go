package commands

import (
	"fmt"
	"github.com/checkr/goflagr"
	"math"
	"os"
	"strconv"
	"strings"
)
func formatFlags(format string, separator string, showHeaders bool, columns []string, flags []goflagr.Flag) {
	switch format {
	case "table":
		renderTable(separator, showHeaders, columns, flags)
	//case "yaml":
	//	renderYAML(assets)
	//case "json":
	//	renderJSON(assets)
	//case "link":
	//	renderLinks(url, remoteLookup, assets)
	default:
		logAndDie(format + " is not a supported format")
	}
}

func fieldToFlagStruct(field string, flag goflagr.Flag) string {
	switch field {
	case "ID":
		return strconv.FormatInt(flag.Id, 10)
	case "Description":
		length := math.Min(40, float64(len(flag.Description)))
		return flag.Description[:int(length)]
	case "Key":
		length := math.Min(40, float64(len(flag.Key)))
		return flag.Key[:int(length)]
	default:
		logAndDie("")
		return ""
	}
}

func renderTable(separator string, showHeaders bool, columns []string, flags []goflagr.Flag) {
	// Find the longest column in each field so the final output is pretty.
	maxColumnWidth := make(map[string]int)
	for _, column := range columns {
		var max int
		for _, flag := range flags {
			length := len(fieldToFlagStruct(column, flag))
			if length > max {
				max = length
			}
		}

		// If headers are going to be output make sure we take them into
		// account when formatting the table.
		if showHeaders && len(column) > max {
			max = len(column)
		}

		maxColumnWidth[column] = max
	}

	// Make sure we build the formatter back in the correct order.
	// Golang you need more datastructures for real.
	var formatterSlice []string
	for _, col := range columns {
		fmtr := "%-" + strconv.Itoa(maxColumnWidth[col]) + "s"
		formatterSlice = append(formatterSlice, fmtr)
	}

	formatter := strings.Join(formatterSlice, separator)

	if showHeaders {
		headers := make([]interface{}, len(columns))
		for i, v := range columns {
			headers[i] = v
		}

		fmt.Fprintf(os.Stderr, formatter+"\n", headers...)
	}

	for _, asset := range flags {
		// We use an interface instead of a slice becasue Printf requires this.
		var fields []interface{}

		for _, column := range columns {
			fields = append(fields, fieldToFlagStruct(column, asset))
		}

		// Cleanup the output a little bit so we don't have trailing tabs or spaces
		// for assets that may be missing lots of information. See the following issue
		// on github for more info https://github.com/michaeljs1990/collins-go-cli/issues/41
		line_to_print := fmt.Sprintf(formatter, fields...)
		fmt.Println(strings.TrimSpace(line_to_print))
	}
}
