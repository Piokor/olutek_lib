package googleapi

import (
	"fmt"
	"strings"
	"testing"
)

func TestOptionalParamsString(t *testing.T) {
	query := BookQuery{
		"",
		"a",
		"b",
		"c",
		"d",
	}
	createdStr := createOptionalParamsString(&query)
	expectedStrings := []string{
		fmt.Sprintf("+inauthor:%s", query.Author),
		fmt.Sprintf("+intitle:%s", query.Title),
		fmt.Sprintf("+inpublisher:%s", query.Publisher),
		fmt.Sprintf("+isbn:%s", query.Isbn),
	}
	for _, expectedStr := range expectedStrings {
		if !strings.Contains(createdStr, expectedStr) {
			t.Fatalf("Optional params not containing expected %s param", expectedStr)
		}
	}
}

func TestPartialParamsString(t *testing.T) {
	query := BookQuery{
		"",
		"a",
		"b",
		"",
		"",
	}
	createdStr := createOptionalParamsString(&query)
	expectedStrings := []string{
		fmt.Sprintf("+inauthor:%s", query.Author),
		fmt.Sprintf("+intitle:%s", query.Title),
	}
	for _, expectedStr := range expectedStrings {
		if !strings.Contains(createdStr, expectedStr) {
			t.Fatalf("Optional params not containing expected %s param", expectedStr)
		}
	}
	unexpectedStrings := []string{
		"+inpublisher",
		"+isbn",
	}
	for _, unexpectedStr := range unexpectedStrings {
		if strings.Contains(createdStr, unexpectedStr) {
			t.Fatalf("Optional params containing uexpected %s param", unexpectedStr)
		}
	}
}

func TestEmptyOptionalParamsString(t *testing.T) {
	query := BookQuery{
		"",
		"",
		"",
		"",
		"",
	}
	createdStr := createOptionalParamsString(&query)
	if createdStr != "" {
		t.Fatalf("Optional params string is not empty for empty query params")
	}
}
