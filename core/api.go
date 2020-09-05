package core

import "myhttps/memory"

/**
@Comment 保存用户提交数据，如果存在返回true,不存在返回false
*/
func CheckData(list []string) []bool {
	result := make([]bool, len(list))
	for key, val := range list {
		if memory.MemStorage.IsExist(val) {
			result[key] = true
			continue
		}
		memory.MemStorage.Push(val, nil)
		result[key] = false
	}
	return result
}
