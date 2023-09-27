package web3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Yihen-Liu/go-service/common/tools"
	"github.com/Yihen-Liu/go-service/log"

	"strings"
)

var BSCTIMEOUT = 15

// LatestBscBlock
//
//	@Description: 获取bsc最新区块
//	@param rpc
//	@return string: 带有0x前缀的hex string
//	@return error
func LatestBscBlock(rpc string) (string, error) {
	var data = strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_blockNumber",
		"params":[],
		"id":83
	}`)

	var resp LatestBscBlockResp

	body, err := tools.Post(rpc, BSCTIMEOUT, data)

	if err != nil {
		log.Error(err)
		return "", err
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(string(body))
		log.Error(err)
		return "", err
	}

	return resp.Result, nil
}

// GetBscBlockByNum
//
//	@Description: 通过块号获取bsc区块信息
//	@param rpc:bsc rpc接口url
//	@param block: 块号
//	@return *BscRpcBlock
//	@return error
func GetBscBlockByNum(rpc, block string) (*BscRpcBlock, error) {
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getBlockByNumber",
		"params":[
			"%s", 
			true
		],
		"id":1
	}`, block))

	body, err := tools.Post(rpc, BSCTIMEOUT, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var resp BscRpcBlock
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil

}

func GetBscTxRecipient(rpc, txHash string) (*BscRpcRecipient, error) {
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getTransactionReceipt",
		"params":[
			"%s"
		],
		"id":1
	}`, txHash))

	body, err := tools.Post(rpc, BSCTIMEOUT, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// log.Logger.Infof(string(body))
	var resp BscRpcRecipient

	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil

}

func GetBscTxCountByNumber(rpc, block string) (*BscTxCountResp, error) {
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getBlockTransactionCountByNumber",
		"params":[
			"%s"
		],
		"id":1
	}`, block))

	body, err := tools.Post(rpc, BSCTIMEOUT, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var resp BscTxCountResp
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func GetBscBlockByNumFal(rpc, block string) (*BscRpcBlockFal, error) {
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getBlockByNumber",
		"params":[
			"%s", 
			false
		],
		"id":1
	}`, block))

	body, err := tools.Post(rpc, BSCTIMEOUT, data)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var resp BscRpcBlockFal
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil

}

func GetTxByHash(rpc, hash string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getTransactionByHash",
		"params":[
			"%s"
		],
		"id":1
	}`, hash))
	return tools.Post(rpc, BSCTIMEOUT, data)
}

func GetTxrByHash(rpc, hash string) {

}

// GetBalance 返回带有0x的hexstring
func GetBalance(rpc string, addr string) (string, error) {
	log.Debug(addr)
	var data = strings.NewReader(fmt.Sprintf(`{
		"jsonrpc":"2.0",
		"method":"eth_getBalance",
		"params":[
			"%s", 
			"latest"
		],
		"id":1
	}`, addr))
	var resp BscRpcBalance
	body, err := tools.Post(rpc, BSCTIMEOUT, data)
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Info(string(body))
	// 校验返回值
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&resp)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return resp.Result, nil
}
