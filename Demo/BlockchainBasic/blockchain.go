/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/31 11:10
  */

package BlockchainBasic

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index 			int				// 块在整个链的位置
	Timestamp 		string			// 块生成的时间戳
	BPM 			int				// 每分钟心跳
	Hash 			string			// 这个块的 hash 值
	PrevHash 		string			// 前一个块的 sha256 散列值
}

// 一个链
var Blockchain []Block

// hash 计算
func calHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 生成块
func generateBlock(oldBlock Block, BPM int) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calHash(newBlock)
	return newBlock
}

// 校验块
func checkBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// 更长的链代表数据更新，过期的链切换为最新的链
func replaceChain(newBlocks []Block)  {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}