package service

import (
	"errors"
	"time"

	"github.com/toky03/oracle-swap-demo/adapter"
	"github.com/toky03/oracle-swap-demo/model"
)

type oracleAdapter interface {
	ReadFunds(string) error
	ReadStatus(string) (model.WalletStatus, error)
	ReadStatusPool(string) (model.WalletStatusPool, error)
	AddFunds(string, model.AddPayload) error
	ClosePool(string, model.ClosePayload) error
	CreatePool(string, model.CreatePayload) error
	ReadPools(string) error
	RemoveFunds(string, model.RemovePayload) error
	Swap(string, model.SwapPayload) error
}

type oracleServiceImpl struct {
	oracleAdapter oracleAdapter
	wallets       map[string]string
	symbol        model.UnCurrencySymbol
}

func CreateOracleService() (*oracleServiceImpl, error) {
	wallets, err := readWallets()
	if err != nil {
		return &oracleServiceImpl{}, err
	}
	symbol, err := readSymbol()
	if err != nil {
		return &oracleServiceImpl{}, err
	}
	oracleAdapter := adapter.CrateAdapter()
	return &oracleServiceImpl{
		oracleAdapter: oracleAdapter,
		wallets:       wallets,
		symbol:        symbol,
	}, nil
}

func (s *oracleServiceImpl) ReadWalletNames() ([]string, error) {
	wallets := make([]string, 0, len(s.wallets))
	for name, _ := range s.wallets {
		wallets = append(wallets, name)
	}
	return wallets, nil
}

func (s *oracleServiceImpl) AddFunds(walletId string, inputCoins model.FundsDto) error {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.AddFunds(walletUuid, inputCoins.ToAddPayload(s.symbol.UnCurrencySymbol))
	if err != nil {
		return err
	}
	_, err = s.oracleAdapter.ReadStatus(walletUuid)
	return err
}

func (s *oracleServiceImpl) ClosePool(walletId string, inputCoins model.CloseDto) error {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.ClosePool(walletUuid, inputCoins.ToClosePayload(s.symbol.UnCurrencySymbol))
	if err != nil {
		return err
	}
	_, err = s.oracleAdapter.ReadStatus(walletUuid)
	return err
}

func (s *oracleServiceImpl) CreatePool(walletId string, input model.FundsDto) error {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.CreatePool(walletUuid, input.ToCreatePayload(s.symbol.UnCurrencySymbol))
	time.Sleep(2 * time.Second)
	if err != nil {
		return err
	}
	_, err = s.oracleAdapter.ReadStatus(walletUuid)
	return err
}

func (s *oracleServiceImpl) ReadFunds(walletId string) (model.WalletDto, error) {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return model.WalletDto{}, errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.ReadFunds(walletUuid)
	if err != nil {
		return model.WalletDto{}, errors.New("error posting for funds, wallet: " + walletId + ", error:" + err.Error())
	}
	time.Sleep(2 * time.Second)
	walletStatus, err := s.oracleAdapter.ReadStatus(walletUuid)
	if err != nil {
		return model.WalletDto{}, errors.New("error with get request for status, wallet: " + walletId + ", error:" + err.Error())
	}
	coinWallet := model.CoinWallet{
		WalletName: walletId,
		WalletUuid: walletUuid,
		Tokens:     model.FromOracleDto(walletStatus),
	}

	err = s.oracleAdapter.ReadPools(walletUuid)
	if err != nil {
		return model.WalletDto{}, errors.New("error posting for funds, wallet: " + walletId + ", error:" + err.Error())
	}
	time.Sleep(2 * time.Second)
	poolStatus, err := s.oracleAdapter.ReadStatusPool(walletUuid)
	if err != nil {
		return model.WalletDto{}, errors.New("error with get request for status, wallet: " + walletId + ", error:" + err.Error())
	}
	pool := model.Pool{
		Tokens: model.FromOracleDtoPool(poolStatus),
	}

	return model.WalletDto{
		Pool:       pool,
		CoinWallet: coinWallet,
	}, nil

}

func (s *oracleServiceImpl) ReadPools(walletId string) (model.Pool, error) {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return model.Pool{}, errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.ReadPools(walletUuid)
	if err != nil {
		return model.Pool{}, errors.New("error posting for funds, wallet: " + walletId + ", error:" + err.Error())
	}
	time.Sleep(2 * time.Second)
	walletStatus, err := s.oracleAdapter.ReadStatusPool(walletUuid)
	if err != nil {
		return model.Pool{}, errors.New("error with get request for status, wallet: " + walletId + ", error:" + err.Error())
	}
	return model.Pool{
		Tokens: model.FromOracleDtoPool(walletStatus),
	}, nil
}

func (s *oracleServiceImpl) RemoveFunds(walletId string, input model.RemoveDto) error {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return errors.New("No wallet found for id " + walletId)
	}
	err := s.oracleAdapter.RemoveFunds(walletUuid, input.ToRemovePayload(s.symbol.UnCurrencySymbol))
	if err != nil {
		return err
	}
	_, err = s.oracleAdapter.ReadStatus(walletUuid)
	return err
}

func (s *oracleServiceImpl) Swap(walletId string, input model.SwapDto) error {
	walletUuid := s.wallets[walletId]
	if walletUuid == "" {
		return errors.New("No wallet found for id " + walletId)
	}

	err := s.oracleAdapter.Swap(walletUuid, input.ToSwapPayload(s.symbol.UnCurrencySymbol))
	time.Sleep(2 * time.Second)
	if err != nil {
		return err
	}
	_, err = s.oracleAdapter.ReadStatus(walletUuid)
	return err
}
