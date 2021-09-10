package _068_text_justification

import (
	"strings"
)

const (
	SPACE = " "
	EMPTY = ""
)

func fullJustify(words []string, maxWidth int) []string {
	/*
	** 贪心算法求解：
	** 1.1 当前行为最后一行，每个单词之间只有一个空格，右边补充剩余空格
	** 2.1 当前行只能放置一个单词，单词左对齐，右边补充剩余空格
	** 3.1 当前行有多个单词且不是最后一行
	** 3.2. 每一行尽可能多的放置单次（不能改变单词词组的顺序）
	** 3.2. maxWidth - len(words) 得到空格数 spaceNums
	** 3.3. avgSpaces = spaceNums / (numsWords - 1)
	** 3.4. extraSpaces = spaceNums % (numsWords - 1)
	** 3.5. 前 extraSpaces 个单词填 avgSpaces + 1 个空格，其余填 avgSpaces 个空格
	 */

	rowWrods, rowLens := make([]string, 0), 0
	rowText, ans := EMPTY, make([]string, 0)
	for index, word := range words {
		rowWrods = append(rowWrods, word)
		rowLens += len(word) + 1

		// 单个单词一行
		if rowLens-1 > maxWidth && len(rowWrods) == 2 {
			rowText = rowWrods[0] + strings.Repeat(SPACE, maxWidth-len(rowWrods[0]))
			ans = append(ans, rowText)
			rowWrods = rowWrods[index+1:]
			continue
		}

		// 最后一行
		if rowLens-1 <= maxWidth && index == len(words)-1 {
			rowText = EMPTY
			for _, word := range rowWrods {
				rowText += word + SPACE
			}

			// 去掉多余的空格
			if len(rowText) > maxWidth {
				rowText = rowText[:len(rowText)-1]
			}

			rowText = rowText + strings.Repeat(SPACE, maxWidth-len(rowText))
			ans = append(ans, rowText)
			continue
		}

		// 既不是单个单词一行也不是最后一行
		spaceNums, numWords := maxWidth-len(rowWrods), len(rowWrods)
		avgSpaces := spaceNums / (numWords - 1)
		extraSpaces := spaceNums % (numWords - 1)

		rowText = EMPTY
		for index, word := range rowWrods {

			if index < extraSpaces {
				rowText += word + strings.Repeat(SPACE, avgSpaces+1)
			} else {
				rowText += word + strings.Repeat(SPACE, avgSpaces)
			}
		}

		rowText = rowText[:maxWidth]
		ans = append(ans, rowText)
	}

	return ans
}
