/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/14 19:33
 */

package Concurrent_Map

import (
	"unsafe"
	"sync/atomic"
	"bytes"
	"fmt"
)

// 单向链接的键值对缺口
type linkedPair interface {
	// Next获取下一个键值对，返回为nil时则在单链表末尾
	Next() Pair
	// SetNext设置下一个键值对
	SetNext(nextPair Pair) error
}

// Pair是一个并发安全的键值对的接口
type Pair interface {
	// 单链键值对接口
	linkedPair
	// Key返回键的值
	Key() string
	// Hash返回键的哈希值
	Hash() uint64
	// Element返回元素的值
	Element() interface{}
	// Set设置元素的值
	SetElement(element interface{}) error
	// Copy生成一个当前键值对的副本并返回
	Copy() Pair
	// String返回当前键值对的字符串表示形式
	String() string
}

// pair代表键值对类型
type pair struct {
	key 	string
	hash 	uint64			// 键的hash值
	element unsafe.Pointer
	next 	unsafe.Pointer
}

// newPair会创建一个Pair类型的实例
func newPair(key string, element interface{}) (Pair, error) {
	p := &pair{
		key:  key,
		hash: hash(key),
	}
	if element == nil {
		return nil, newIllegalParameterError("element is nil")
	}
	p.element = unsafe.Pointer(&element)
	return p, nil
}

func (pair *pair) Key() string {
	return pair.key
}

func (pair *pair) Hash() uint64 {
	return pair.hash
}

func (p *pair) Element() interface{} {
	pointer := atomic.LoadPointer(&p.element)
	if pointer == nil {
		return nil
	}
	return *(*interface{})(pointer)
}

func (p *pair) SetElement(element interface{}) error {
	if element == nil {
		return newIllegalParameterError("element is nil")
	}
	atomic.StorePointer(&p.element, unsafe.Pointer(&element))
	return nil
}

func (p *pair) Next() Pair {
	pointer := atomic.LoadPointer(&p.next)
	if pointer == nil {
		return nil
	}
	return (*pair)(pointer)
}

func (p *pair) SetNext(nextPair Pair) error {
	if nextPair == nil {
		atomic.StorePointer(&p.next, nil)
		return nil
	}
	pp, ok := nextPair.(*pair)
	if !ok {
		return newIllegalPairTypeError(nextPair)
	}
	atomic.StorePointer(&p.next, unsafe.Pointer(pp))
	return nil
}

// Copy 会生成一个当前键-元素对的副本并返回。
func (p *pair) Copy() Pair {
	pCopy, _ := newPair(p.Key(), p.Element())
	return pCopy
}

func (p *pair) String() string {
	return p.genString(false)
}

// genString 用于生成并返回当前键-元素对的字符串形式。
func (p *pair) genString(nextDetail bool) string {
	var buf bytes.Buffer
	buf.WriteString("pair{key:")
	buf.WriteString(p.Key())
	buf.WriteString(", hash:")
	buf.WriteString(fmt.Sprintf("%d", p.Hash()))
	buf.WriteString(", element:")
	buf.WriteString(fmt.Sprintf("%+v", p.Element()))
	if nextDetail {
		buf.WriteString(", next:")
		if next := p.Next(); next != nil {
			if npp, ok := next.(*pair); ok {
				buf.WriteString(npp.genString(nextDetail))
			} else {
				buf.WriteString("<ignore>")
			}
		}
	} else {
		buf.WriteString(", nextKey:")
		if next := p.Next(); next != nil {
			buf.WriteString(next.Key())
		}
	}
	buf.WriteString("}")
	return buf.String()
}
















