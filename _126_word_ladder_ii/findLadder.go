package _126_word_ladder_ii

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordLadders, wordsMap := make([][]string, 0), make(map[string]bool)

	for _, word := range wordList {
		wordsMap[word] = true
	}

	if !wordsMap[endWord] {
		return wordLadders
	}

	// create a queue, track the path
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})

	// queuesLen is used to track how many slices in queue are in the same level
	// if found a result, I still need to finish checking current level cause I
	// need to return all possible paths
	queueLen := 1
	// use to track strings that this level has visited when queuesLne == 0, re-
	// move levelMap keys in wordMap
	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]
		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				if nextWord == endWord {
					// record current path
					path = append(path, endWord)
					wordLadders = append(wordLadders, path)
					continue
				}
				if wordsMap[nextWord] {
					// different from word ladder, don't remove the word from wordsMap im
					// -mediately, same level would reuse the key. delete from wordMap only
					// when currently level is done
					levelMap[nextWord] = true // warning todo
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--
		// if queueLen is 0, means finish traversing current level. if wordLadders is
		// not return wordLadders
		if queueLen == 0 {
			if len(wordLadders) > 0 {
				return wordLadders
			}

			for word := range levelMap {
				delete(wordsMap, word)
			}
			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return wordLadders
}
