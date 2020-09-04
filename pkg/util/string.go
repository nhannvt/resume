package util

import (
	"regexp"
	"strings"
)

// CamelCaseToSnakeCase converts string format from camelcase to snakecase.
// For example
//
//   TestCamelCaseToSnakeCase -> test_camel_case_to_snake_case
//   testCamelCaseToSnakeCase -> test_camel_case_to_snake_case
func CamelCaseToSnakeCase(camelCase string) string {
	rep := regexp.MustCompile("[A-Z]")
	snakeCase := rep.ReplaceAllString(camelCase, "_$0")
	snakeCase = strings.Trim(snakeCase, "_")
	snakeCase = strings.ToLower(snakeCase)

	return snakeCase
}
