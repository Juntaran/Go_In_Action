/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/1/31 17:02
  */

package Blockchain

import (
	"time"
	"fmt"
	"bytes"
	"encoding/binary"
	"net/url"
	"encoding/json"
	"net/http"
)

type BlockchainService interface {
	// 新增 node 到 nodes
	RegisterNode(address string) bool

	// 判断区块链是否正确
	ValidChain(chain Blockchain) bool

	// 如果冲突，取最长链
	ResolveConflicts() bool

	// 区块链新增一个 block
	NewBlock(proof int64, prevHash string) Block

	// 创建一个新交易到下一个挖掘块
	NewTransaction(tx Transaction) int64

	// 返回最后一个 block
	LastBlock() Block

	// Simple Proof of Work Algorithm:
	// - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
	// - p is the previous proof, and p' is the new proof
	ProofOfWork(lastProof int64)

	// Validates the Proof: Does hash(lastProof, proof) contain 4 leading zeroes?
	VerifyProof(lastProof, proof int64) bool
}

type Block struct {
	Index 			int64			`json:"index"`				// 块在整个链的位置
	Timestamp 		int64			`json:"timestamp"`			// 块生成的时间戳
	Transactions 	[]Transaction	`json:"transactions"`		// 每分钟心跳
	Proof 			int64			`json:"proof"`				// 这个块的 hash 值
	PrevHash 		string			`json:"prevHash"`		// 前一个块的 sha256 散列值
}

type Transaction struct {
	Sender			string			`json:"sender"`
	Reciever		string			`json:"reciever"`
	Amount			int64			`json:"amount"`
}

type Blockchain struct {
	blockchain 		[]Block
	transaction 	[]Transaction
	nodes 			StringSet
}

// 区块链新增一个 block
func (bc *Blockchain) NewBlock(proof int64, prevHash string) Block {
	if prevHash == "" {
		prevBlock := bc.blockchain[len(bc.blockchain) - 1]
		prevHash = calHashForBlock(prevBlock)
	}
	newBlock := Block{
		Index: 			int64(len(bc.blockchain) + 1),
		Timestamp:		time.Now().UnixNano(),
		Transactions: 	bc.transaction,
		Proof: 			proof,
		PrevHash: 		prevHash,
	}
	bc.transaction = nil
	bc.blockchain = append(bc.blockchain, newBlock)
	return newBlock
}

// 创建一个新交易
func (bc *Blockchain) NewTransaction(tx Transaction) int64 {
	bc.transaction = append(bc.transaction, tx)
	return bc.LastBlock().Index + 1
}

// 返回最后一个 block
func (bc *Blockchain) LastBlock() Block {
	return bc.blockchain[len(bc.blockchain) - 1]
}

// Simple Proof of Work Algorithm:
// - Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
// - p is the previous proof, and p' is the new proof
func (bc *Blockchain) ProofOfWork(lastProof int64) int64 {
	var proof int64 = 0
	for !bc.ValidProof(lastProof, proof) {
		proof += 1
	}
	return proof
}

// Validates the Proof: Does hash(lastProof, proof) contain 4 leading zeroes?
func (bc *Blockchain) ValidProof(lastProof, proof int64) bool {
	guess := fmt.Sprintf("%d%d", lastProof, proof)
	guessHash := CalHashSha256([]byte(guess))
	return guessHash[:4] == "0000"
}

// 判断区块链是否正确
func (bc *Blockchain) ValidChain(chain *[]Block) bool {
	lastBlock := (*chain)[0]
	currentIndex := 1
	for currentIndex < len(*chain) {
		block := (*chain)[currentIndex]
		// 检查 block 的 hash 值是否正确
		if block.PrevHash != calHashForBlock(lastBlock) {
			return false
		}
		// 检查 work 的 proof 是否正确
		if !bc.ValidProof(lastBlock.Proof, block.Proof) {
			return false
		}
		lastBlock = block
		currentIndex += 1
	}
	return true
}

func (bc *Blockchain) RegisterNode(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	return bc.nodes.Add(u.Host)
}

func (bc *Blockchain) ResolveConflicts() bool {
	neighbours := bc.nodes
	newChain := make([]Block, 0)
	maxLength := len(bc.blockchain)

	for _, node := range neighbours.Keys() {
		otherBlockchain, err := findExternalChain(node)
		if err != nil {
			continue
		}
		if otherBlockchain.Length {

		}
	}
}

func calHashForBlock(block Block) string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, block)
	return CalHashSha256(buf.Bytes())
}

type blockchainInfo struct {
	Length int     `json:"length"`
	Chain  []Block `json:"chain"`
}

func findExternalChain(address string) (blockchainInfo, error) {
	response, err := http.Get(fmt.Sprintf("http://%s/chain", address))
	if err == nil && response.StatusCode == http.StatusOK {
		var bi blockchainInfo
		if err := json.NewDecoder(response.Body).Decode(&bi); err != nil {
			return blockchainInfo{}, err
		}
		return bi, nil
	}
	return blockchainInfo{}, err
}