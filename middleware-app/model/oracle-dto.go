package model

type AddPayload struct {
	ApAmountA uint32 `json:"apAmountA"`
	ApAmountB uint32 `json:"apAmountB"`
	ApCoinA   Coin   `json:"apCoinA"`
	ApCoinB   Coin   `json:"apCoinB"`
}
type CreatePayload struct {
	CpAmountA uint32 `json:"cpAmountA"`
	CpAmountB uint32 `json:"cpAmountB"`
	CpCoinA   Coin   `json:"cpCoinA"`
	CpCoinB   Coin   `json:"cpCoinB"`
}

type RemovePayload struct {
	RpDiff uint32 `json:"rpDiff"`
	RpCoinA   Coin   `json:"rpCoinA"`
	RpCoinB   Coin   `json:"rpCoinB"`
}

type ClosePayload struct {
	ClpCoinB Coin `json:"clpCoinB"`
	ClpCoinA Coin `json:"clpCoinA"`
}

type SwapPayload struct {
	SpAmountA uint32 `json:"spAmountA"`
	SpAmountB uint32 `json:"spAmountB"`
	SpCoinA   Coin   `json:"spCoinA"`
	SpCoinB   Coin   `json:"spCoinB"`
}

type Coin struct {
	UnAssetClass []map[string]string `json:"unAssetClass"`
}

type AssetClass struct {
	Symbol      string `json:"unCurrencySymbol"`
	UnTokenName string `json:"unTokenName"`
}

type WalletStatus struct {
	CicCurrentState CicCurrentState `json:"cicCurrentState"`
}

type CicCurrentState struct {
	ObservableState ObservableState `json:"observableState"`
}

type ObservableState struct {
	Right Right `json:"Right"`

}

type WalletStatusPool struct {
	CicCurrentState CicCurrentStatePool `json:"cicCurrentState"`
}

type CicCurrentStatePool struct {
	ObservableState ObservableStatePool `json:"observableState"`
}

type ObservableStatePool struct {
	Right RightPool `json:"Right"`

}
type RightPool struct {
	Contents [][][]interface{} `json:"contents"`
	Tag string `json:"tag"`
}

type Right struct {
	Contents Content `json:"contents"`
	Tag string `json:"tag"`
}

type Content struct {
	GetValue [][]interface{} `json:"getValue"`
}

type UnCurrencySymbol struct {
	UnCurrencySymbol string `json:"unCurrencySymbol"`
}

type UnTokenName struct {
	UnTokenName string `json:"unTokenName"`
}
