package _678_valid_parenthesis_string

/*
** resolve by stack
** 1. create two slices chSlice and starSlice to store left bracket and start
** 2. traveling string s
** 2.1 if element is ( , save into chSlice
** 2.2 if element is * , save info starSlice
** 2.3 if element is ) , find matched element from chSlice and starSlice.
** 2.4.1 if matched element doesn't exist, return false
** 3. compare chSlice and starSlice
** 3.1 if len(chSlice) > len(starSlice) return false
** 4. pop chSlice and starSlice element
** 4.1 if chSlice top element index greater than starSlice top element index, return false
** 5. return true
 */

type Elem struct {
	Index int
	Value int32
}

func CheckValidString(s string) bool {
	chSlice, starSlice := make([]Elem, 0), make([]Elem, 0)
	for i, ch := range s {
		// stack push
		if ch == '(' {
			chSlice = append(chSlice, Elem{i, ch})
			continue
		} else if ch == '*' {
			starSlice = append(starSlice, Elem{i, ch})
			continue
		}

		// bracket match
		if len(chSlice) > 0 {
			chSlice = chSlice[:len(chSlice)-1]
		} else if len(starSlice) > 0 {
			starSlice = starSlice[:len(starSlice)-1]
		} else {
			return false
		}
	}

	chSliceLen, starSliceLen := len(chSlice), len(starSlice)
	if chSliceLen > starSliceLen {
		return false
	}
	for chSliceLen > 0 && starSliceLen > 0 {
		if chSlice[chSliceLen-1].Index > starSlice[starSliceLen-1].Index {
			return false
		}

		chSliceLen--
		starSliceLen--
	}

	return true
}
