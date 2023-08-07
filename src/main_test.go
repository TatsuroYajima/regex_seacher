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
