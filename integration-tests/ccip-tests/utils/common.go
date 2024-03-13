package utils

import (
	"path/filepath"
	"runtime"
	"sync"
)

func ProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "/..")
}

func DeleteNilEntriesFromMap(inputMap *sync.Map) *sync.Map {
	// checks for nil entry in map, store all not-nil entries to another map and deallocates previous map
	newMap := &sync.Map{}
	inputMap.Range(func(key, value any) bool {
		if value != nil {
			newMap.Store(key, value)
		}
		return true
	})
	return newMap
}
