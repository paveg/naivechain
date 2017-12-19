package main

import (
	"encoding/hex"
	"testing"
)

const genesisBlockHex = "00000000000000000000000000000000000000000000000000000000000000000000000000000000917c5457000000006d792067656e6573697320626c6f636b2121"

func TestCreateBlockFromBytes(t *testing.T) {
	blockBytes, err := hex.DecodeString(genesisBlockHex)
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	block, err := createBlockFromBytes(blockBytes)
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	if block.Index != genesisBlock.Index {
		t.Errorf("should have %q but %q", genesisBlock.PreviousHash, block.PreviousHash)
	}
	if block.Timestamp != genesisBlock.Timestamp {
		t.Errorf("should have %d, but %d", genesisBlock.Timestamp, block.Timestamp)
	}
	if block.Data != genesisBlock.Data {
		t.Errorf("should have %q but %q", genesisBlock.Data, block.Data)
	}
}

func TestBlockHex(t *testing.T) {
	blockHex, err := genesisBlock.hex()
	if err != nil {
		t.Fatalf("shouild not be fail: %v", err)
	}
	if blockHex != genesisBlockHex {
		t.Errorf("should %q but %q", genesisBlockHex, blockHex)
	}
}
