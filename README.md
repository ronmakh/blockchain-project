# Simple Blockchain CLI in Go

This is a simple command-line blockchain application written in Go. The project simulates a basic blockchain system where you can:

- Add transactions
- Mine blocks
- View the blockchain
- View wallet balances

## Features

- ✅ Add new transactions  
- ✅ Mine new blocks to the blockchain  
- ✅ View the entire blockchain  
- ✅ View wallet balances (simple ledger)  
- ✅ Basic CLI interface

## How it works

- **Transactions** are added to a pending transactions pool.
- **Mining a block** will take all pending transactions and add them into a new block, which is appended to the blockchain.
- **Viewing the blockchain** will show all blocks, their transactions, and metadata like hashes.
- **Wallet balances** are calculated by iterating over the blockchain.

---

## CLI Options

When you run the program, you'll get a menu like this:

### 1. Add Transaction
- Enter sender wallet address
- Enter receiver wallet address
- Enter amount

### 2. Mine Block
- Mines a new block with pending transactions.

### 3. View Blockchain
- Prints all blocks with details like transactions, previous hash, and hash.

### 4. View Wallet Balances
- Displays each wallet's balance.

### 5. Exit
- Closes the application.

---

## Getting Started

### Prerequisites

- Go 1.18+

### Running the app

```bash
git clone https://github.com/ronmakh/blockchain-project.git
cd go-blockchain-cli
go run main.go
```

## Near Future Enhancements
- Add Peer-to-Peer (P2P) Networking
- Add cryptographic signatures using public/private key pairs (Ensure only the wallet owner can authorize a transaction)
- Include mempool-like behaviour (queue for unconfirmed transactions)
- Create REST APIs to interact with the blockchain