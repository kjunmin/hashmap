package hashmap

import (
	"errors"
	"reflect"
)

type Key interface{}
type Value interface{}

type HashMaper interface {
	Insert(key Key, value Value) error
	Get(key Key) (value Value, err error)
	Erase(key Key) error
	Count() int
}

type HashMap struct {
	hashFunc         func(blockSize int, key Key) (hashKey uint, bucketIdx uint)
	defaultBlockSize int
	buckets          []*BucketNode
	size             int
}

type BucketNode struct {
	hashKey uint
	key     Key
	value   Value
	next    *BucketNode
}

func Init(blockSize int, fn func(blockSize int, key Key) (hashKey uint, bucketIdx uint)) *HashMap {
	hashMap := new(HashMap)
	hashMap.hashFunc = fn
	hashMap.defaultBlockSize = blockSize
	hashMap.buckets = make([]*BucketNode, blockSize)
	hashMap.size = 0

	return hashMap
}

// Hashmap methods
func (h *HashMap) Get(key Key) (value Value, err error) {
	hashKey, bucketIdx := h.hashFunc(len(h.buckets), key)
	bucketNode := h.buckets[bucketIdx]
	for bucketNode != nil {
		if bucketNode.hashKey == hashKey && reflect.DeepEqual(key, bucketNode.key) {
			return bucketNode.value, nil
		}
		bucketNode = bucketNode.next
	}
	return nil, errors.New("Key not found")
}

func (h *HashMap) Insert(key Key, value Value) error {
	hashKey, bucketIdx := h.hashFunc(len(h.buckets), key)
	bucketNode := h.buckets[bucketIdx]
	newNode := &BucketNode{
		hashKey: hashKey,
		key:     key,
		value:   value,
		next:    bucketNode,
	}
	h.buckets[bucketIdx] = newNode
	h.size++
	return nil
}

func (h *HashMap) Erase(key Key) error {
	hashKey, bucketIdx := h.hashFunc(len(h.buckets), key)
	bucketNode := h.buckets[bucketIdx]
	if bucketNode == nil {
		return errors.New("Key to delete not found")
	}

	var prev *BucketNode
	for bucketNode != nil {
		if bucketNode.hashKey == hashKey && reflect.DeepEqual(key, bucketNode.key) {
			if prev == nil && bucketNode.next == nil {
				h.buckets[bucketIdx] = nil
			} else if prev == nil {
				h.buckets[bucketIdx] = bucketNode.next
			} else {
				prev.next = bucketNode.next
			}
		}
		prev = bucketNode
		bucketNode = bucketNode.next
	}
	h.size--
	return nil
}

func (h *HashMap) Count() int {
	return h.size
}
