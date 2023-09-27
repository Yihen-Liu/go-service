package web3

import (
	"github.com/Yihen-Liu/go-service/common/tools"
	"github.com/Yihen-Liu/go-service/log"
)

// LatestBlockNumber 返回带有0x的hexstring
func LatestBlockNumber(rpc, chain string) (int64, error) {
	var blockNum int64
	var err error
	switch chain {
	case "bsc":
		block, err := LatestBscBlock(rpc)
		if err != nil {
			return blockNum, err
		}
		blockNum, err = tools.Hex2int64(block)
		if err != nil {
			return blockNum, err
		}

	}
	return blockNum, err
}

func GetBlockByNum(rpc, block, chain string) (interface{}, error) {
	switch chain {
	case "bsc":
		return GetBscBlockByNum(rpc, block)
	}
	return nil, nil
}

func GetTokenAddrssByHash(rpc, txHash string) (map[string]struct{}, error) {
	// tokenAddress := make([]string, 10)
	tokenAddress := make(map[string]struct{})

	bscRpcTxRecipient, err := GetBscTxRecipient(rpc, txHash)
	if err != nil {
		return nil, err
	}

	for _, singleLog := range bscRpcTxRecipient.Result.Logs {
		// tokenAddress = append(tokenAddress, singleLog.Address)
		log.Debugf("get token %s", singleLog.Address)
		tokenAddress[singleLog.Address] = struct{}{}
	}
	return tokenAddress, nil

}
