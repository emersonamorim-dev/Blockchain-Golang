package main

import (
    "testing"
)

func TestBlockchain(t *testing.T) {
    blockchain := CreateBlockchain(2)

    // test if the genesis block is valid
    if len(blockchain.chain) != 1 {
        t.Errorf("Expected 1 block, got %d", len(blockchain.chain))
    }
    if blockchain.chain[0].hash != "0" {
        t.Errorf("Expected genesis block hash to be '0', got %s", blockchain.chain[0].hash)
    }

    // record transactions on the blockchain for Ana, Gaby and Nathalia
    blockchain.addBlock("Ana", "Gaby", 8)
    blockchain.addBlock("Nathalia", "Gaby", 2)

    // check if the Blockchain is valid; expecting true
    if !blockchain.isValid() {
        t.Errorf("Blockchain is invalid")
    }
}
