package leetCode

const (
	BLANK = byte(' ')
)
func ReplaceBlank(src []byte) []byte {
	blankCount := 0
	for _, value := range src {
		if value == BLANK {
			blankCount++
		}
	}

	srcP, desP := len(src)-1, len(src)+2*blankCount-1
	src = append(src, make([]byte, 2*blankCount)...)

	for ;srcP >= 0 && srcP != desP; {
		if src[srcP] != BLANK {
			src[desP] = src[srcP]
			srcP--
			desP--
		} else {
			src[desP] = byte('0')
			desP--
			src[desP] = byte('2')
			desP--
			src[desP] = byte('%')
			desP--
			srcP--
		}
	}

	return src
}
