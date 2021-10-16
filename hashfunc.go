package hashmap

import (
	"bytes"
	"reflect"
)

// Hashing function
func HashFunc(blockSize int, key Key) (hashKey uint, bucketIdx uint) {
	var buf bytes.Buffer
	buf.Write([]byte(reflect.ValueOf(key).String()))

	hash := dbj2hash(&buf)

	return hash, (hash % uint(blockSize))
}

func dbj2hash(buf *bytes.Buffer) uint {
	var hash uint = 5381

	for _, c := range buf.Bytes() {
		hash = (hash << 5) + hash + uint(c)
	}

	return hash
}
