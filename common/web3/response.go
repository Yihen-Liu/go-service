package web3

import (
	"github.com/Yihen-Liu/go-service/common/db"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type LatestBscBlockResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type BscTxCountResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type LatestTMBlockResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Response struct {
			Data             string `json:"data"`
			Version          string `json:"version"`
			AppVersion       string `json:"app_version"`
			LastBlockHeight  string `json:"last_block_height"`
			LastBlockAppHash string `json:"last_block_app_hash"`
		} `json:"response"`
	} `json:"result"`
}

type BscRpcBalance struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type BscRpcTransaction struct {
	db.BscTransactionCommon
	Gas      string `gorm:"column:gas" json:"gas"`
	GasPrice string `gorm:"column:gasPrice" json:"gasPrice"`
	Nonce    string `gorm:"column:nonce" json:"nonce"`
	V        string `gorm:"column:v" json:"v"`
	// GasUsed          string `gorm:"column:gasUsed" json:"gasUsed"`
	Timestamp        string `gorm:"column:timestamp" json:"timestamp"`
	BlockNumber      string `gorm:"column:blockNumber;NOT NULL" json:"blockNumber"`
	TransactionIndex string `gorm:"column:transactionIndex" json:"transactionIndex"`
	Type             string `gorm:"column:type" json:"type"`
}

/*
BscRpcBlock
bsc获取区块rpc接口返回
eth_getBlockByNumber with true
*/
type BscRpcBlock struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		db.BscBlockCommon

		Size      string `gorm:"column:size" json:"size"`
		GasLimit  string `gorm:"column:gasLimit" json:"gasLimit"`
		GasUsed   string `gorm:"column:gasUsed" json:"gasUsed"`
		Timestamp string `gorm:"column:timestamp" json:"timestamp"`
		Number    string `gorm:"column:number;NOT NULL" json:"number"`

		// 数据库未存此字段
		MixHash      string              `json:"mixHash"`
		Transactions []BscRpcTransaction `json:"transactions"`
		// CreditData,CreditValue,CreditMax
		TrustNodeScore string `json:"trustNodeScore"`
		// TODO 暂未发现json uncles具体数据
		Uncles []string `json:"uncles"`
	} `json:"result"`
}

/*
BscRpcBlock
bsc获取区块rpc接口返回
eth_getBlockByNumber with false
*/
type BscRpcBlockFal struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		db.BscBlockCommon

		Size      string `gorm:"column:size" json:"size"`
		GasLimit  string `gorm:"column:gasLimit" json:"gasLimit"`
		GasUsed   string `gorm:"column:gasUsed" json:"gasUsed"`
		Timestamp string `gorm:"column:timestamp" json:"timestamp"`
		Number    string `gorm:"column:number;NOT NULL" json:"number"`

		// 数据库未存此字段
		MixHash      string   `json:"mixHash"`
		Transactions []string `json:"transactions"`
		// CreditData,CreditValue,CreditMax
		TrustNodeScore string `json:"trustNodeScore"`
		// TODO 暂未发现json uncles具体数据
		Uncles []string `json:"uncles"`
	} `json:"result"`
}

// BscRpcTXResp bsc获取交易rpc接口返回
type BscRpcTXResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BscRpcTransaction
	} `json:"result"`
}

type TMRpcBlock struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total int    `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Block struct {
			Header struct {
				Version struct {
					Block string `json:"block"`
					App   string `json:"app"`
				} `json:"version"`
				ChainID     string    `json:"chain_id"`
				Height      string    `json:"height"`
				Time        time.Time `json:"time"`
				LastBlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"last_block_id"`
				LastCommitHash     string `json:"last_commit_hash"`
				DataHash           string `json:"data_hash"`
				ValidatorsHash     string `json:"validators_hash"`
				NextValidatorsHash string `json:"next_validators_hash"`
				ConsensusHash      string `json:"consensus_hash"`
				AppHash            string `json:"app_hash"`
				LastResultsHash    string `json:"last_results_hash"`
				EvidenceHash       string `json:"evidence_hash"`
				ProposerAddress    string `json:"proposer_address"`
				TrustHash          string `json:"trust_hash"`
			} `json:"header"`
			Data struct {
				Txs []string `json:"txs"`
			} `json:"data"`
			Evidence struct {
				Evidence []interface{} `json:"evidence"`
			} `json:"evidence"`
			LastCommit struct {
				Height  string `json:"height"`
				Round   int    `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total int    `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"block_id"`
				Signatures []struct {
					BlockIDFlag      int       `json:"block_id_flag"`
					ValidatorAddress string    `json:"validator_address"`
					Timestamp        time.Time `json:"timestamp"`
					Signature        string    `json:"signature"`
				} `json:"signatures"`
			} `json:"last_commit"`
			TrustData  []string `json:"trust_data"`
			TrustNodes []struct {
				Pubkey string `json:"pubkey"`
				Score  string `json:"score"`
				IP     string `json:"ip"`
			} `json:"trust_nodes"`
		} `json:"block"`
	} `json:"result"`
}

type BscRpcRecipient struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Receipt
	} `json:"result"`
}

type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              string `json:"type,omitempty"`
	Status            string `json:"status"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" gencodec:"required"`
	EffectiveGasPrice string `json:"effectivegasprice" gencodec:"required"`
	Bloom             string `json:"logsBloom"         gencodec:"required"`
	Logs              []*Log `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`
	// ContractAddress common.Address `json:"contractAddress"`
	From            string `json:"from"`
	To              string `json:"to"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         string `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash string `json:"blockHash,omitempty"`
	// BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	TransactionIndex string `json:"transactionIndex"`
}

type Log struct {
	// Consensus fields:
	// address of the contract that generated the event
	Address string `json:"address" gencodec:"required"`
	// list of topics provided by the contract.
	Topics []string `json:"topics" gencodec:"required"`
	// supplied by the contract, usually ABI-encoded
	// Data []byte `json:"data" gencodec:"required"`
	Data string `json:"data" gencodec:"required"`

	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber string `json:"blockNumber"`
	// hash of the transaction
	TxHash string `json:"transactionHash" gencodec:"required"`
	// index of the transaction in the block
	TxIndex string `json:"transactionIndex"`
	// hash of the block in which the transaction was included
	BlockHash common.Hash `json:"blockHash"`
	// index of the log in the block
	Index string `json:"logIndex"`

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool `json:"removed"`
}
