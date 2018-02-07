package block

import "fmt"

/**
创世块
 */
var GenesisBlock = &Block{
    Index: 0,
    PreviousHash: "0",
    Timestamp: 1517714059,
    Data: "dazhenghu genesis block",
    Hash: "6f19a25ee96fb031e12bd9f40b23d0456bb17542022d78945d09e13420c3e539",
}

/**
区块数据结构
 */
type Block struct {
    Index int64 `json:"index"` // 区块索引
    PreviousHash string `json:"PreviousHash"` // 前块哈希
    Timestamp int64 `json:"TimeStamp"` // 时间戳
    Data string `json:"Data"` // 区块数据
    Hash string `json:"hash"` // 区块哈希
}

func (b *Block)String() string {
    return fmt.Sprintf("index: %d,previousHash:%s,timestamp:%d,data:%s,hash:%s", b.Index, b.PreviousHash, b.Timestamp, b.Data, b.Hash)
}
