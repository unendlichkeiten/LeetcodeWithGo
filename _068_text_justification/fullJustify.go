package _068_text_justification

import (
	"strings"
)

const (
	SPACE = " "
	EMPTY = ""
)

func FullJustify(words []string, maxWidth int) []string {
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

	rowWrods, ans, index := make([]string, 0), make([]string, 0), 0
	rowLens, rowLensNoSpaces, rowText := 0, 0, EMPTY

	for {
		rowWrods = append(rowWrods, words[index])
		rowLens += len(words[index]) + 1
		rowLensNoSpaces += len(words[index])

		// 单个单词一行
		if rowLens-1 > maxWidth && len(rowWrods) == 2 {
			rowText = rowWrods[0] + strings.Repeat(SPACE, maxWidth-len(rowWrods[0]))
			rowLens, rowLensNoSpaces = len(words[index])+1, len(words[index])
			rowWrods = rowWrods[1:]
			ans = append(ans, rowText)
		}

		// 既不是单个单词一行也不是最后一行
		if rowLens-1 > maxWidth && rowLens-1-len(words[index])-1 <= maxWidth {
			spaceNums, numWords := maxWidth-rowLensNoSpaces+len(words[index]), len(rowWrods)-1
			avgSpaces := spaceNums / (numWords - 1)
			extraSpaces := spaceNums % (numWords - 1)

			rowText = EMPTY
			for index, rowWord := range rowWrods {
				if index < extraSpaces {
					rowText += rowWord + strings.Repeat(SPACE, avgSpaces+1)
				} else {
					rowText += rowWord + strings.Repeat(SPACE, avgSpaces)
				}
			}

			rowText = rowText[:maxWidth]
			rowLens, rowLensNoSpaces = len(words[index])+1, len(words[index])
			rowWrods = rowWrods[len(rowWrods)-1:]
			ans = append(ans, rowText)
		}

		// 最后一行，这里主要考虑前面处理结束只剩一个单词的情况
		if rowLens-1 <= maxWidth && index == len(words)-1 {
			rowText = EMPTY
			for _, rowWord := range rowWrods {
				rowText += rowWord + SPACE
			}

			// 去掉多余的空格
			if len(rowText) > maxWidth {
				rowText = rowText[:len(rowText)-1]
			}

			rowText = rowText + strings.Repeat(SPACE, maxWidth-len(rowText))
			ans = append(ans, rowText)
			break
		}

		index++
	}

	return ans
}
