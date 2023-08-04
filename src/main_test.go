package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestMain(t *testing.T) {
	pattern := `[0-9]+[A-Za-z]+`
	regExp := regexp.MustCompile(pattern)

	var lines []string

	lines = append(lines, "123abc")

	result := findMatchedStrings(regExp,lines)

	expected := []string{"123abc"}
	
	if !reflect.DeepEqual(result, expected) { 
		t.Errorf("got %v want %v", result, expected)
    }
}
