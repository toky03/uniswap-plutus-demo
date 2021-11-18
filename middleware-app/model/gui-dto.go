package model

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Pool struct {
	Tokens []Token `json:"tokens"`
}

type FundsDto struct {
	CoinA CoinDto `json:"coinA"`
	CoinB CoinDto `json:"coinB"`
}

type SwapDto struct {
	SwapCoin        CoinDto `json:"swapCoin"`
	ToCoinTokenName string  `json:"toCoinTokenName"`
}

type CloseDto struct {
	TokenNameCoinA string `json:"tokenNameCoinA"`
	TokenNameCoinB string `json:"tokenNameCoinB"`
}

type RemoveDto struct {
	TokenNameCoinA string `json:"tokenNameCoinA"`
	TokenNameCoinB string `json:"tokenNameCoinB"`
	RemoveDiff     uint32 `json:"removeDiff"`
}

func (d *FundsDto) ToAddPayload(symbol string) AddPayload {
	return AddPayload{
		ApAmountA: d.CoinA.Amount,
		ApAmountB: d.CoinB.Amount,
		ApCoinA:   createUnAssetClass(symbol, d.CoinA.TokenName),
		ApCoinB:   createUnAssetClass(symbol, d.CoinB.TokenName),
	}
}

func (d *FundsDto) ToCreatePayload(symbol string) CreatePayload {
	return CreatePayload{
		CpAmountA: d.CoinA.Amount,
		CpAmountB: d.CoinB.Amount,
		CpCoinA:   createUnAssetClass(symbol, d.CoinA.TokenName),
		CpCoinB:   createUnAssetClass(symbol, d.CoinB.TokenName),
	}
}

func (d *CloseDto) ToClosePayload(symbol string) ClosePayload {
	return ClosePayload{
		ClpCoinB: createUnAssetClass(symbol, d.TokenNameCoinB),
		ClpCoinA: createUnAssetClass(symbol, d.TokenNameCoinA),
	}
}

func (d *RemoveDto) ToRemovePayload(symbol string) RemovePayload {
	return RemovePayload{
		RpCoinB: createUnAssetClass(symbol, d.TokenNameCoinB),
		RpCoinA: createUnAssetClass(symbol, d.TokenNameCoinA),
		RpDiff:  d.RemoveDiff,
	}
}

func (d *SwapDto) ToSwapPayload(symbol string) SwapPayload {
	return SwapPayload{
		SpAmountA: d.SwapCoin.Amount,
		SpAmountB: 0,
		SpCoinA:   createUnAssetClass(symbol, d.SwapCoin.TokenName),
		SpCoinB:   createUnAssetClass(symbol, d.ToCoinTokenName),
	}
}

func createUnAssetClass(symbol, tokenName string) Coin {
	return Coin{[]map[string]string{
		{"unCurrencySymbol": symbol},
		{"unTokenName": tokenName},
	}}
}

type CoinDto struct {
	TokenName string `json:"tokenName"`
	Amount    uint32 `json:"amount"`
}

type WalletDto struct {
	Pool Pool `json:"poolWallet"`
	CoinWallet CoinWallet `json:"coinWallet"`
}

type PoolWallet struct {
	WalletName string  `json:"walletName"`
	WalletUuid string  `json:"uuid"`
	Tokens     []Token `json:"tokens"`
}

type Token struct {
	CurrencySymbol string `json:"currencySymbol"`
	TokenName      string `json:"tokenName"`
	Ammount        int64  `json:"amount"`
}

type CoinWallet struct {
	WalletName string  `json:"walletName"`
	WalletUuid string  `json:"uuid"`
	Tokens     []Token `json:"tokens"`
}

type OfferedLovelaces struct {
	OfferedLovelaces int64 `json:"offeredLovelaces"`
}

func FromOracleDto(walletStatus WalletStatus) []Token {
	tokens := make([]Token, 0)
	for _, state := range walletStatus.CicCurrentState.ObservableState.Right.Contents.GetValue {
		currencySymbolRaw := fmt.Sprintf("%v", state[0])
		resultCurrencySymbol := extractText(currencySymbolRaw)
		for _, currencyState := range state[1].([]interface{}) {
			tokenNameRaw := fmt.Sprintf("%v", currencyState.([]interface{})[0])
			ammountRaw := fmt.Sprintf("%f", currencyState.([]interface{})[1])
			tokenName := extractText(tokenNameRaw)
			if tokenName == "" && resultCurrencySymbol == "" {
				tokenName = "Lovelace"
			}
			if tokenName == "" && resultCurrencySymbol != "" {
				tokenName = "Pool Token"
			}
			tokens = append(tokens, Token{
				CurrencySymbol: resultCurrencySymbol,
				TokenName:      tokenName,
				Ammount:        extractAmmount(ammountRaw),
			})
		}
	}
	return tokens
}

func FromOracleDtoPool(walletStatus WalletStatusPool) []Token {
	tokens := make([]Token, 0)
	if len(walletStatus.CicCurrentState.ObservableState.Right.Contents) == 0 {
		return []Token{}
	}
	for _, content := range walletStatus.CicCurrentState.ObservableState.Right.Contents {
		for _, state := range content {
			currencySymbolRaw := fmt.Sprintf("%v", state[0].(map[string]interface{})["unAssetClass"].([]interface{})[0])
			tokenNameRaw := fmt.Sprintf("%v", state[0].(map[string]interface{})["unAssetClass"].([]interface{})[1])
			ammountRaw := fmt.Sprintf("%f", state[1])
			resultCurrencySymbol := extractText(currencySymbolRaw)
			resultTokenName := extractText(tokenNameRaw)
			tokens = append(tokens, Token{
				CurrencySymbol: resultCurrencySymbol,
				TokenName:      resultTokenName,
				Ammount:        extractAmmount(ammountRaw),
			})
		}
	}
	return tokens
}

func extractText(raw string) string {
	currencyRegex := regexp.MustCompile(`map\[\w*:([a-zA-Z0-9_\s]*)\]`)
	result := currencyRegex.FindAllStringSubmatch(raw, 1)
	if result == nil || len(result) < 1 {
		return ""
	}
	return result[0][1]
}

func extractAmmount(raw string) int64 {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		log.Printf("could not convert %v to int", raw)
		return 0
	}
	return int64(val)
}
