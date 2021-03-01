package xlint

import (
	"bytes"

	"github.com/koykov/fastconv"
)

var (
	urnPrefixL = []byte("urn:uuid:")
	urnPrefixU = []byte("URN:UUID:")
	payload32  = [16]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	payload36  = [16]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	xvalues    = [256]byte{
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
		255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	}
)

func ValidateUUIDStr(s string) (bool, error) {
	return ValidateUUID(fastconv.S2B(s))
}

func ValidateUUID(p []byte) (bool, error) {
	var c []byte
	switch len(p) {
	case 36:
		// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		c = p
	case 38:
		// {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}
		c = p[1:37]
	case 45:
		// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		if !bytes.Equal(p[:9], urnPrefixL) && !bytes.Equal(p[:9], urnPrefixU) {
			return false, ErrPrefixUUID
		}
		c = p[9:]
	case 32:
		// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
		c = p
	default:
		return false, ErrLengthUUID
	}

	var payload [16]int
	if len(c) == 32 {
		payload = payload32
	} else {
		if c[8] != '-' || c[13] != '-' || c[18] != '-' || c[23] != '-' {
			return false, ErrFormatUUID
		}
		payload = payload36
	}

	for i := 0; i < 16; i++ {
		x := payload[i]
		if _, ok := xtob(c[x], c[x+1]); !ok {
			return false, ErrFormatUUID
		}
	}

	return true, nil
}

func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}
