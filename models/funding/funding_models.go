package funding

import "github.com/amir-the-h/okex"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string           `json:"transId"`
		Ccy     string           `json:"ccy"`
		Amt     okex.JSONFloat64 `json:"amt"`
		From    okex.AccountType `json:"from,string"`
		To      okex.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string           `json:"billId"`
		Ccy    string           `json:"ccy"`
		Bal    okex.JSONFloat64 `json:"bal"`
		BalChg okex.JSONFloat64 `json:"balChg"`
		Type   okex.BillType    `json:"type,string"`
		TS     okex.JSONTime    `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtID    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to,string"`
		TS       okex.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string            `json:"ccy"`
		Chain string            `json:"chain"`
		TxID  string            `json:"txId"`
		From  string            `json:"from"`
		To    string            `json:"to"`
		DepId string            `json:"depId"`
		Amt   okex.JSONFloat64  `json:"amt"`
		State okex.DepositState `json:"state,string"`
		TS    okex.JSONTime     `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		WdID  okex.JSONInt64   `json:"wdId"`
		Amt   okex.JSONFloat64 `json:"amt"`
	}

	WithdrawalHistory struct {
		Ccy              string               `json:"ccy"`              // 币种
		Chain            string               `json:"chain"`            // 币种链信息
		NonTradableAsset bool                 `json:"nonTradableAsset"` // 是否为不可交易资产, true：不可交易资产，false：可交易资产
		Amt              okex.JSONFloat64     `json:"amt"`              // 数量
		TS               okex.JSONTime        `json:"ts"`               // 提币申请时间, 毫秒级时间戳
		From             string               `json:"from"`             // 提币账户
		AreaCodeFrom     string               `json:"areaCodeFrom"`     // 如果from为手机号，该字段为该手机号的区号
		To               string               `json:"to"`               // 收币地址
		AreaCodeTo       string               `json:"areaCodeTo"`       // 如果to为手机号，该字段为该手机号的区号
		Tag              string               `json:"tag,omitempty"`    // 部分币种提币需要标签
		PmtID            string               `json:"pmtId,omitempty"`  // 部分币种提币需要此字段（payment_id）
		Memo             string               `json:"memo,omitempty"`
		AddrEx           map[string]any       `json:"addrEx"`       // 提币地址备注，部分币种提币需要, 币种TONCOIN的提币地址备注标签名为comment,则该字段返回：{'comment':'123456'}
		TxID             string               `json:"txId"`         // 提币哈希记录 内部转账该字段返回""
		Fee              okex.JSONFloat64     `json:"fee"`          // 提币手续费数量
		FeeCcy           string               `json:"feeCcy"`       // 提币手续费币种
		State            okex.WithdrawalState `json:"state,string"` // 提币状态
		WdID             okex.JSONInt64       `json:"wdId"`         // 提币申请ID
		ClientId         string               `json:"clientId"`     //客户自定义ID
		Note             string               `json:"note"`         // 备注信息
	}

	PiggyBank struct {
		Ccy  string           `json:"ccy"`
		Amt  okex.JSONFloat64 `json:"amt"`
		Side okex.ActionType  `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string           `json:"ccy"`
		Amt      okex.JSONFloat64 `json:"amt"`
		Earnings okex.JSONFloat64 `json:"earnings"`
	}
)
