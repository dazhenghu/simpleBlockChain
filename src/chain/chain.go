package chain

import (
    "github.com/dazhenghu/simpleBlockChain/src/block"
    "time"
    "crypto/sha256"
    "fmt"
    "log"
)

/**
区块链是有序list，为了保证有序性，需要对内存中保存区块的数组进行排序，定义[]*block.Block按index排序类型，实现排序接口
 */
type ByIndex []*block.Block

/**
获取链的长度
 */
func (byIndex ByIndex)Len() int  {
    return len(byIndex)
}

/**
位置互换
 */
func (byIndex ByIndex)Swap(i, j int)  {
    byIndex[i], byIndex[j] = byIndex[j], byIndex[i]
}

/**
比较index大小
 */
func (byIndex ByIndex)Less(i, j int) bool  {
    return byIndex[i].Index < byIndex[j].Index
}


var blockChain = ByIndex{block.GenesisBlock} // 链，初始化为创世块



/**
取链中的最后一位区块
 */
func GetLastestBlock() *block.Block {
    return blockChain[len(blockChain) - 1]
}

/**
创建新区块
 */
func GenerateNextBlock(data string) (nextBlock *block.Block)  {
    previousBlock := GetLastestBlock()
    nextBlock = &block.Block{
        Data: data,
        PreviousHash: previousBlock.Hash,
        Index: previousBlock.Index + 1,
        Timestamp: time.Now().Unix(),
    }

    nextBlock.Hash = CalculateHashForBlock(nextBlock)

    return
}

/**
将新区块加入链中
 */
func AddBlock(b *block.Block)  {
    if isValidNewBlock(b, GetLastestBlock()) {
        blockChain = append(blockChain, b)
    }
}

/**
校验相邻的两个区块，新区块是否有效
 */
func isValidNewBlock(nextBlock, previousBlock *block.Block) (ok bool) {

    /**
    index值相邻 and hash值计算正确 and nextBlock的前块hash值匹配
     */
    if nextBlock.Hash == CalculateHashForBlock(nextBlock) &&
        nextBlock.Index == previousBlock.Index + 1 &&
        previousBlock.Hash == nextBlock.PreviousHash {
        ok = true
    }

    return
}

/**
校验区块链的合法性
 */
func isValidChain(bc []*block.Block) bool {
    if bc[0].String() != block.GenesisBlock.String() {
        log.Println("No same GenesisBlock.(创世区块不正确)", bc[0].String())
        return false
    }

    var temp = ByIndex{bc[0]}
    for i := 1; i < len(bc); i++ {
        if isValidNewBlock(bc[i], temp[i - 1]) {
            temp = append(temp, bc[i])
        } else {
            return false
        }
    }

    return true
}

/**
根据区块生成hash值
 */
func CalculateHashForBlock(b *block.Block) string {
    shaVal := sha256.Sum256([]byte(fmt.Sprintf("%d%s%d%s", b.Index, b.PreviousHash, b.Timestamp, b.Data)))
    return fmt.Sprintf("%x", shaVal)
}
