package rewardaccount

import (
	"fmt"
	"testing"
)

func TestVerifyBlocks(t *testing.T) {
	//data
	blocks := make([]BlockEventRecord, 0)
	blocks = append(blocks, BlockEventRecord{BlockId: "12321"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12312"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12233"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12231"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12201"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12123"})
	blocks = append(blocks, BlockEventRecord{BlockId: "12111"})

	account := "abc"
	interval := int64(12)
	b, reward := VerifyBlocks(account, blocks, interval)
	fmt.Println("b:", b)
	fmt.Println("reward:", reward)

}
