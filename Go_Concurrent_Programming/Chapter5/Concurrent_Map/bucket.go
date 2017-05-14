/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/14 22:21
 */

package Concurrent_Map

import (
	"sync"
	"sync/atomic"
	"bytes"
)

// Bucket时并发安全的散列桶接口
type Bucket interface {
	// Put插入一个键值对
	// 如果调用该方法之前已经锁定lock，则不应传入lock；反之必须传入对应lock
	Put(p Pair, lock sync.Locker) (bool, error)
	// Get获取指定键值对
	Get(key string) Pair
	// GetFirstPair返回第一个键值对
	GetFirstPair() Pair
	// Delete删除指定的键值对
	// 如果调用该方法之前已经锁定lock，则不应传入lock；反之必须传入对应lock
	Delete(key string, lock sync.Locker) bool
	// Clear清空当前桶
	// 如果调用该方法之前已经锁定lock，则不应传入lock；反之必须传入对应lock
	Clear(locker sync.Locker)
	// Size返回当前散列桶尺寸
	Size() uint64
	// String返回当前散列桶字符串表示形式
	String() string
}

// bucket为并发安全的散列桶的类型
type bucket struct {
	firstValue 	atomic.Value	// 键值对列表的表头
	size 		uint64
}

// 占位符，原子值不能存储nil，当散列桶为空时占位
var placeholder Pair = &pair{}

// 创建一个Bucket类型实例
func newBucket() Bucket {
	b := &bucket{}
	b.firstValue.Store(placeholder)
	return b
}

func (b *bucket) Put(p Pair, lock sync.Locker) (bool, error) {
	if p == nil {
		return false, newIllegalParameterError("pair is nil")
	}
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	firstPair := b.GetFirstPair()
	if firstPair == nil {
		b.firstValue.Store(p)
		atomic.AddUint64(&b.size, 1)
		return true, nil
	}
	var target Pair
	key := p.Key()
	for v := firstPair; v != nil; v = v.Next() {
		if v.Key() == key {
			target = v
			break
		}
	}
	if target != nil {
		target.SetElement(p.Element())
		return false, nil
	}
	p.SetNext(firstPair)
	b.firstValue.Store(p)
	atomic.AddUint64(&b.size, 1)
	return true, nil
}

func (b *bucket) Get(key string) Pair {
	firstPair := b.GetFirstPair()
	if firstPair == nil {
		return nil
	}
	for v := firstPair; v != nil; v = v.Next() {
		if v.Key() == key {
			return v
		}
	}
	return nil
}

func (b *bucket) GetFirstPair() Pair {
	if v := b.firstValue.Load(); v == nil {
		return nil
	} else if p, ok := v.(Pair); !ok || p == placeholder {
		return nil
	} else {
		return p
	}
}

func (b *bucket) Delete(key string, lock sync.Locker) bool {
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	firstPair := b.GetFirstPair()
	if firstPair == nil {
		return false
	}
	var prevPairs []Pair
	var target Pair
	var breakpoint Pair
	for v := firstPair; v != nil; v = v.Next() {
		if v.Key() == key {
			target = v
			breakpoint = v.Next()
			break
		}
		prevPairs = append(prevPairs, v)
	}
	if target == nil {
		return false
	}
	newFirstPair := breakpoint
	for i := len(prevPairs) - 1; i >= 0; i-- {
		pairCopy := prevPairs[i].Copy()
		pairCopy.SetNext(newFirstPair)
		newFirstPair = pairCopy
	}
	if newFirstPair != nil {
		b.firstValue.Store(newFirstPair)
	} else {
		b.firstValue.Store(placeholder)
	}
	atomic.AddUint64(&b.size, ^uint64(0))
	return true
}

func (b *bucket) Clear(lock sync.Locker) {
	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}
	atomic.StoreUint64(&b.size, 0)
	b.firstValue.Store(placeholder)
}

func (b *bucket) Size() uint64 {
	return atomic.LoadUint64(&b.size)
}

func (b *bucket) String() string {
	var buf bytes.Buffer
	buf.WriteString("[ ")
	for v := b.GetFirstPair(); v != nil; v = v.Next() {
		buf.WriteString(v.String())
		buf.WriteString(" ")
	}
	buf.WriteString("]")
	return buf.String()
}
