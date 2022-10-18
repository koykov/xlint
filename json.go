package xlint

import "github.com/koykov/jsonlint"

var _ = ValidateJSONStr

func ValidateJSONStr(s string) (int, error) {
	return jsonlint.ValidateStr(s)
}

func ValidateJSON(s []byte) (int, error) {
	return jsonlint.Validate(s)
}
