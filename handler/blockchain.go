package handler

import (
	"blockchain-project/model"
	"blockchain-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTransaction(c *gin.Context) {
	var tx model.Transaction
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction format"})
		return
	}
	service.AddTransaction(tx.Sender, tx.Receiver, tx.Amount)
	c.JSON(http.StatusCreated, gin.H{"status": "Transaction added"})
}

func HandleMine(c *gin.Context) {
	service.MineBlock()
	c.JSON(http.StatusOK, gin.H{"status": "Block mined successfully"})
}

func GetBlockchain(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetBlockchain())
}

func GetWallets(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetWallets())
}
