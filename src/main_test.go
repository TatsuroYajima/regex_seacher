package main

import (
	"reflect"
	"testing"
)

func Test_123abc_should_be_matched(t *testing.T) {
	targetStrings := []string{"123abc"}

	result := findMatchedStrings(targetStrings)

	expected := []string{"123abc"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}
