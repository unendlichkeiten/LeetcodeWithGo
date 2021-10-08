package _187_repeated_dna_sequences

/*
** 解题思路：滑动窗口
 */
const LEN = 10

func findRepeatedDnaSequences(s string) (ans []string) {
	cnt := map[string]int{}
	for i := 0; i <= len(s)-LEN; i++ {
		sub := s[i : i+LEN]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return
}
