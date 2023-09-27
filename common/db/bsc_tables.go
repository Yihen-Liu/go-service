package db

// BscTransactionCommon 与web3公用的交易相关字段
type BscTransactionCommon struct {
	BlockHash string `gorm:"column:blockHash;index;type:char(66);NOT NULL" json:"blockHash"`
	// json字段为from,原有的python项目使用source,这里使用已有的数据库字段
	From  string `gorm:"column:source;index;type:char(42)" json:"from"`
	To    string `gorm:"column:to;index;type:char(42)" json:"to"`
	Hash  string `gorm:"column:hash;type:char(66);unique" json:"hash"`
	Value string `gorm:"column:value;type:char(66)" json:"value"`
	R     string `gorm:"column:r;type:char(66)" json:"r"`
	S     string `gorm:"column:s;type:char(66)" json:"s"`
	V     string `gorm:"column:v;type:char(66)" json:"v"`
	// json字段为input,原有的python项目使用tx_str,这里使用已有的数据库字段
	Input string `gorm:"column:tx_str;type:text" json:"input"`
}

// BscBlockCommon 与web3 rpc公用的字段
type BscBlockCommon struct {
	// Number           int64  `gorm:"column:number;NOT NULL" json:"number"`
	Hash             string `gorm:"column:hash;type:char(66);uniqueIndex;NOT NULL" json:"hash"`
	ParentHash       string `gorm:"column:parentHash;type:char(66);NOT NULL" json:"parentHash"`
	Nonce            string `gorm:"column:nonce;type:char(66)" json:"nonce"`
	Sha3Uncles       string `gorm:"column:sha3Uncles;type:char(66)" json:"sha3Uncles"`
	LogsBloom        string `gorm:"column:logsBloom" json:"logsBloom"`
	TransactionsRoot string `gorm:"column:transactionsRoot;type:char(66);" json:"transactionsRoot"`
	StateRoot        string `gorm:"column:stateRoot;type:char(66)" json:"stateRoot"`
	ReceiptsRoot     string `gorm:"column:receiptsRoot;type:char(66)" json:"receiptsRoot"`
	Miner            string `gorm:"column:miner;type:char(42)" json:"miner"`
	Difficulty       string `gorm:"column:difficulty;type:char(66)" json:"difficulty"`
	TotalDifficulty  string `gorm:"column:totalDifficulty;type:char(66)" json:"totalDifficulty"`
	ExtraData        string `gorm:"column:extraData" json:"extraData"`
	Size             int64  `gorm:"column:size" json:"size"`
	// GasLimit         int64  `gorm:"column:gasLimit" json:"gasLimit"`
	// GasUsed          int64  `gorm:"column:gasUsed" json:"gasUsed"`
	// Timestamp        int64  `gorm:"column:timestamp" json:"timestamp"`
}
