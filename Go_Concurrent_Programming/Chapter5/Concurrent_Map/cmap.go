/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/14 18:41
 */

package Concurrent_Map

import (
	"sync/atomic"
	"math"
)

// ConcurrentMap是一个并发安全的字典的接口
type ConcurrentMap interface {
	// 返回并发量
	Concurrency() int
	/*
		Put会推送一个键值对，element不能为nil，
		第一个返回值表示是否新增了键值对
		如果该键已经存在，新元素替换旧元素
	*/
	Put(key string, element interface{}) (bool, error)
	/*
		Get获取与指定键关联的元素
		如果返回nil，则该键不存在
	*/
	Get(key string) interface{}
	/*
		Delete删除指定的键值对
		如果结果为true，则成功
	*/
	Delete(key string) bool
	// Len返回当前字典中键值对数量
	Len() uint64
}

// myConcurrentMap为ConcurrentMap接口的实现类型
type myConcurrentMap struct {
	concurrency		int				// 并发量，也是segments的长度
	segments 		[]Segment		// 一个Segment代表一个散列段，每个散列段都有互斥锁保护——分段锁
									// 同一个散列段只能有一个goroutine读写，但是不同的散列段可以并发访问
	total 			uint64			// 当前字典中键值对的实际数量
}

// 创建一个ConcurrentMap类型的实例，参数pairRedistributor可以为nil
func NewConcurrentMap(
	concurrency int,
	pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	cmap := &myConcurrentMap{
		concurrency: 	concurrency,
		segments: 		make([]Segment, concurrency),
	}
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}

func (cmap *myConcurrentMap) Concurrency() int {
	return cmap.concurrency
}

func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
	p, err := newPair(key, element)
	if err != nil {
		return false, err
	}
	s := cmap.findSegment(p.Hash())
	ok, err := s.Put(p)
	if ok {
		atomic.AddUint64(&cmap.total, 1)
	}
	return ok, err
}

func (cmap *myConcurrentMap) Get(key string) interface{} {
	keyHash := hash(key)
	s := cmap.findSegment(keyHash)
	pair := s.GetWithHash(key, keyHash)
	if pair == nil {
		return nil
	}
	return pair.Element()
}

func (cmap *myConcurrentMap) Delete(key string) bool {
	s := cmap.findSegment(hash(key))
	if s.Delete(key) {
		atomic.AddUint64(&cmap.total, ^uint64(0))
		return true
	}
	return false
}

func (cmap *myConcurrentMap) Len() uint64 {
	return atomic.LoadUint64(&cmap.total)
}

// 根据给定参数寻找并返回对应的散列段
func (cmap *myConcurrentMap) findSegment(keyHash uint64) Segment {
	if cmap.concurrency == 1 {
		return cmap.segments[0]
	}
	var keyHash32 uint32
	if keyHash > math.MaxUint32 {
		keyHash32 = uint32(keyHash >> 32)
	} else {
		keyHash32 = uint32(keyHash)
	}
	return cmap.segments[int(keyHash32 >> 16) % (cmap.concurrency - 1)]
}