package main

import (
	"reflect"
	"testing"
)

func Test_invalid_sql_should_be_matched(t *testing.T) {
	targetString := []string{"SELECT * FROM test_table WHERE id = 123UNION ALL"}

	result := findMatchedString(targetString)

	expected := []string{"SELECT * FROM test_table WHERE id = 123UNION ALL"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}

func Test_quoted_string_should_not_be_matched(t *testing.T) {
	targetString := []string{`2023-07-27 04:01:22 UTC:10.0.1.164(33036):neo_search@neo_search:[28374]:LOG:  duration: 2567.922 ms  statement: SELECT DISTINCT "UOrder"."id" AS "UOrder__expired", to_char("UOrder"."created", 'YYYY/MM/DD'`}

	result := findMatchedString(targetString)

	var expected []string

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}

func Test_valid_sql_should_not_be_matched(t *testing.T) {
	targetString := []string{"SELECT * FROM test_table WHERE id = 123;"}

	result := findMatchedString(targetString)
	
	var expected []string

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v want %v", result, expected)
	}
}
