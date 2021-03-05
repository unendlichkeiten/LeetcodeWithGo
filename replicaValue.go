package leetCode

import "errors"

func ReplicaValue_Hash(array []int)(int, error){
	arrMap := make(map[int]int)

	for _, value := range array {
		if _, ok := arrMap[value]; ok {
			return value, nil
		} else {
			arrMap[value] = 1
		}
	}

	return -1, errors.New("no value replica")
}

func ReplicaValue_Scan(array []int)(int, error){
	for key, value := range array {
		compare:
		if value == key {
			continue
		} else {
			if array[value] == value {
				return value, nil
			} else {
				array[value], value = value, array[value]
				goto compare
			}
		}
	}

	return -1, errors.New("no value replica")
}