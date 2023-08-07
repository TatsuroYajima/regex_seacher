package main

import (
	"reflect"
	"testing"
)

func Test_invalid_sql_should_be_matched(t *testing.T) {
	targetStrings := []string{"SELECT * FROM test_table WHERE id = 123UNION ALL"}

	result := findMatchedStrings(targetStrings)

	expected := []string{"SELECT * FROM test_table WHERE id = 123UNION ALL"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}

func Test_valid_sql_should_not_be_matched(t *testing.T) {
	targetStrings := []string{"SELECT * FROM test_table WHERE id = 123;"}

	result := findMatchedStrings(targetStrings)
	
	var expected []string

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
func Test_string_wrapped_in_single_quotes_should_not_be_matched(t *testing.T) {
	targetString := []string{"\"serial_code\" = ('test0196a740c98318ee084')"}

	result := findMatchedString(targetString)
	
	var expected []string

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}
