package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

var genesisBlock = &Block{
	Index:        0,
	PreviousHash: "000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	Timestamp:    123456789,
	Data:         "genesis block",
}

type Block struct {
	Index        int64  `json:"index"`
	PreviousHash string `json:"previousHash"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
}

func createBlockFromBytes(b []byte) (*Block, error) {
	var index int64
	if err := binary.Read(bytes.NewReader(b[:8]), binary.LittleEndian, &index); err != nil {
		return nil, err
	}

	var timestamp int64
	if err := binary.Read(bytes.NewReader(b[40:48]), binary.LittleEndian, &timestamp); err != nil {
		return nil, err
	}

	return &Block{
		Index:        index,
		PreviousHash: hex.EncodeToString(b[8:40]),
		Timestamp:    timestamp,
		Data:         string(b[48:]),
	}, nil
}

func (block *Block) bytes() ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.LittleEndian, block.Index); err != nil {
		return nil, err
	}

	previousHashBytes, err := hex.DecodeString(block.PreviousHash)
	if err != nil {
		return nil, err
	}

	if _, err := buf.Write(previousHashBytes); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, block.Timestamp); err != nil {
		return nil, err
	}

	if _, err := buf.WriteString(block.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (block *Block) hex() (string, error) {
	b, err := block.bytes()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func (block *Block) hash() (string, error) {
	b, err := block.bytes()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", sha256.Sum256(b)), nil
}
