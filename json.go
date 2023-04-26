package xlint

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/jsonlint"
)

func ValidateJSON[T byteseq.Byteseq](x T) (int, error) {
	return jsonlint.Validate(x)
}
