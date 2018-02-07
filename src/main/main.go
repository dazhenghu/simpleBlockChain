package main

import (
    "fmt"
    "github.com/dazhenghu/simpleBlockChain/src/chain"
    "github.com/dazhenghu/simpleBlockChain/src/block"
)

func main()  {
    fmt.Printf("%s\n", chain.CalculateHashForBlock(block.GenesisBlock))
}

