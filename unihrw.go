package unihrw

import (
	"bytes"
	"errors"
	"hash"
	"reflect"
	"sort"
)

type (
	Rawer interface{ Raw() []byte }
)

func NewHrw32Map(slice interface{}, h hash.Hash32) map[int]uint32 {
	val := reflect.ValueOf(slice)
	result := make(map[int]uint32, val.Len())
	for i := 0; i < val.Len(); i++ {
		result[i] = getHash32(val.Index(i).Interface().(Rawer).Raw(), h)
	}
	return result
}

func NewHrw64Map(slice interface{}, h hash.Hash64) map[int]uint64 {
	val := reflect.ValueOf(slice)
	result := make(map[int]uint64, val.Len())
	for i := 0; i < val.Len(); i++ {
		result[i] = getHash64(val.Index(i).Interface().(Rawer).Raw(), h)
	}
	return result
}

func getHash32(data []byte, h hash.Hash32) uint32 {
	h.Reset()
	h.Write(data) // we ignore errors here, sorry
	return h.Sum32()
}

func getHash64(data []byte, h hash.Hash64) uint64 {
	h.Reset()
	h.Write(data) // we ignore errors here, sorry
	return h.Sum64()
}

func finalizer32(h uint32) uint32 {
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func finalizer64(h uint64) uint64 {
	h ^= h >> 33
	h = h * 0xff51afd7ed558ccd
	h ^= h >> 33
	h = h * 0xc4ceb9fe1a85ec53
	h ^= h >> 33
	return h
}

func HrwSort32(slice interface{}, object []byte, h hash.Hash32) error {
	pivot := finalizer32(getHash32(object, h))

	val := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)

	switch slice := slice.(type) {
	// todo: add more cases
	case [][]byte:
		sort.Slice(slice, func(i, j int) bool {
			// return getHash32(slice[i], h) < getHash32(slice[j], h)
			return bytes.Compare(slice[i], slice[j]) == -1
		})
	case []int:
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	case []string: // pre defines for hash
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	default:
		if _, ok := val.Index(0).Interface().(Rawer); !ok {
			return errors.New("argument is not slice of rawers")
		}
		hrwmap := NewHrw32Map(slice, h)
		sort.Slice(slice, func(i, j int) bool {
			return hrwmap[i] < hrwmap[j]
		})
	}

	for i := 0; i < val.Len()-1; i++ {
		offset := pivot % uint32(val.Len()-i)
		swap(i, i+int(offset))
	}
	return nil
}

func HrwSort64(slice interface{}, object []byte, h hash.Hash64) error {
	pivot := finalizer64(getHash64(object, h))

	val := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)

	switch slice := slice.(type) {
	// todo: add more cases
	case [][]byte:
		sort.Slice(slice, func(i, j int) bool {
			return bytes.Compare(slice[i], slice[j]) == -1
		})
	case []int:
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	case []string: // pre defines for hash
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	default:
		if _, ok := val.Index(0).Interface().(Rawer); !ok {
			return errors.New("argument is not slice of rawers")
		}
		hrwmap := NewHrw64Map(slice, h)
		sort.Slice(slice, func(i, j int) bool {
			return hrwmap[i] < hrwmap[j]
		})
	}

	for i := 0; i < val.Len()-1; i++ {
		offset := pivot % uint64(val.Len()-i)
		swap(i, i+int(offset))
	}
	return nil
}
