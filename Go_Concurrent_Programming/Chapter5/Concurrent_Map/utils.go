/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/15 2:17
 */

package Concurrent_Map

// 本函数实现了BKDR哈希算法
func hash(str string) uint64 {
	seed := uint64(13131)
	var hash uint64
	for i := 0; i < len(str); i++ {
		hash = hash * seed + uint64(str[i])
	}
	return (hash & 0x7FFFFFFFFFFFFFFF)
}
