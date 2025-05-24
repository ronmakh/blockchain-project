package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"blockchain-project/model"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Blockchain []model.Block
var PendingTransactions []model.Transaction
var Wallets = make(map[string]*model.Wallet)
var mutex = &sync.Mutex{}

const FIXED_MINING_REWARD = 10

func init() {
	Blockchain = append(Blockchain, genesisBlock())
}

func genesisBlock() model.Block {
	block := model.Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: nil,
		PrevHash:     "",
		Nonce:        0,
		Difficulty:   4,
	}
	block.Hash = calculateHash(block)
	return block
}

func calculateHash(block model.Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%v", block.Transactions) + block.PrevHash + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Difficulty)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func generateBlock(oldBlock model.Block) model.Block {
	newBlock := model.Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: PendingTransactions,
		PrevHash:     oldBlock.Hash,
		Nonce:        0,
		Difficulty:   4,
	}
	PendingTransactions = nil
	targetPrefix := strings.Repeat("0", newBlock.Difficulty)

	for {
		newBlock.Hash = calculateHash(newBlock)
		if strings.HasPrefix(newBlock.Hash, targetPrefix) {
			break
		}
		newBlock.Nonce++
	}

	mutex.Lock()
	for _, tx := range newBlock.Transactions {
		if senderWallet, ok := Wallets[tx.Sender]; ok && senderWallet.Balance >= tx.Amount {
			senderWallet.Balance -= tx.Amount
			Wallets[tx.Receiver].Balance += tx.Amount
		}
	}
	minerAddress := fmt.Sprintf("miner_%d", rand.Intn(1000000))
	if _, exists := Wallets[minerAddress]; !exists {
		Wallets[minerAddress] = &model.Wallet{Address: minerAddress, Balance: 0}
	}
	Wallets[minerAddress].Balance += FIXED_MINING_REWARD
	mutex.Unlock()

	return newBlock
}

func AddTransaction(sender, receiver string, amount float64) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := Wallets[sender]; !exists {
		Wallets[sender] = &model.Wallet{Address: sender, Balance: 100}
	}
	if _, exists := Wallets[receiver]; !exists {
		Wallets[receiver] = &model.Wallet{Address: receiver, Balance: 0}
	}

	if Wallets[sender].Balance >= amount {
		PendingTransactions = append(PendingTransactions, model.Transaction{Sender: sender, Receiver: receiver, Amount: amount})
	}
}

func MineBlock() {
	mutex.Lock()
	lastBlock := Blockchain[len(Blockchain)-1]
	mutex.Unlock()

	newBlock := generateBlock(lastBlock)

	mutex.Lock()
	Blockchain = append(Blockchain, newBlock)
	mutex.Unlock()
}

func GetBlockchain() []model.Block {
	return Blockchain
}

func GetWallets() map[string]*model.Wallet {
	return Wallets
}
