package _551_student_abset_record

func checkRecord(s string) bool {
	absent, late, pre := 0, 0, ' '
	for _, v := range s {
		if v == 'A' {
			absent++
		} else if v == 'L' {
			if pre == 'L' || pre == ' ' {
				late++
			} else {
				late = 1
			}
		}

		pre = v
		if absent == 2 || late == 3 {
			return false
		}
	}

	return true
}
