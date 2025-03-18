package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
}

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Wallet struct {
	Address string
	Balance float64
}

var Blockchain []Block
var PendingTransactions []Transaction
var mutex = &sync.Mutex{}
var Wallets = make(map[string]*Wallet)

const FIXED_MINING_REWARD = 10

// CalculateHash generates a SHA-256 hash for a block
func CalculateHash(block Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%v", block.Transactions) + block.PrevHash + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Difficulty)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GenerateBlock creates a new block with proof-of-work
func GenerateBlock(oldBlock Block) Block {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: PendingTransactions,
		PrevHash:     oldBlock.Hash,
		Nonce:        0,
		Difficulty:   4,
	}
	PendingTransactions = nil // Clear pending transactions after being added to a block
	targetPrefix := strings.Repeat("0", newBlock.Difficulty)
	for {
		newBlock.Hash = CalculateHash(newBlock)
		if strings.HasPrefix(newBlock.Hash, targetPrefix) {
			break
		}
		newBlock.Nonce++
	}

	// Process transactions and update wallet balances
	mutex.Lock()
	for _, tx := range newBlock.Transactions {
		if senderWallet, exists := Wallets[tx.Sender]; exists {
			if senderWallet.Balance >= tx.Amount {
				senderWallet.Balance -= tx.Amount
				Wallets[tx.Receiver].Balance += tx.Amount
			}
		}
	}
	// Reward miner with new coins
	minerAddress := fmt.Sprintf("miner_%d", rand.Intn(1000000))
	if _, exists := Wallets[minerAddress]; !exists {
		Wallets[minerAddress] = &Wallet{Address: minerAddress, Balance: 0}
	}
	Wallets[minerAddress].Balance += FIXED_MINING_REWARD
	mutex.Unlock()

	return newBlock
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() Block {
	genesis := Block{Index: 0, Timestamp: time.Now().String(), Transactions: nil, PrevHash: "", Nonce: 0, Difficulty: 4}
	genesis.Hash = CalculateHash(genesis)
	return genesis
}

func addTransaction(sender, receiver string, amount float64) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := Wallets[sender]; !exists {
		Wallets[sender] = &Wallet{Address: sender, Balance: 100} // Default balance
	}
	if _, exists := Wallets[receiver]; !exists {
		Wallets[receiver] = &Wallet{Address: receiver, Balance: 0}
	}
	if Wallets[sender].Balance >= amount {
		PendingTransactions = append(PendingTransactions, Transaction{Sender: sender, Receiver: receiver, Amount: amount})
		fmt.Println("Transaction added successfully.")
	} else {
		fmt.Println("Insufficient balance.")
	}
}

func mineBlock() {
	mutex.Lock()
	lastBlock := Blockchain[len(Blockchain)-1]
	mutex.Unlock()

	newBlock := GenerateBlock(lastBlock) // No lock inside GenerateBlock

	mutex.Lock()
	Blockchain = append(Blockchain, newBlock) // Lock only when updating Blockchain
	mutex.Unlock()

	fmt.Println("Block mined successfully:", newBlock.Hash)
}

func printBalances() {
	fmt.Println("Wallet Balances:")
	for _, wallet := range Wallets {
		fmt.Printf("%s: %.2f\n", wallet.Address, wallet.Balance)
	}
}

func printBlockchain() {
	fmt.Println("Blockchain:")
	for _, block := range Blockchain {
		fmt.Printf("Index: %d, Hash: %s, PrevHash: %s, Transactions: %v\n", block.Index, block.Hash, block.PrevHash, block.Transactions)
	}
}

func cli() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Mine Block")
		fmt.Println("3. View Blockchain")
		fmt.Println("4. View Wallet Balances")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter sender: ")
			scanner.Scan()
			sender := scanner.Text()

			fmt.Print("Enter receiver: ")
			scanner.Scan()
			receiver := scanner.Text()

			fmt.Print("Enter amount: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(scanner.Text(), 64)

			addTransaction(sender, receiver, amount)

		case "2":
			mineBlock()

		case "3":
			printBlockchain()

		case "4":
			printBalances()

		case "5":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}

func main() {
	Blockchain = append(Blockchain, GenesisBlock())
	cli()
}
