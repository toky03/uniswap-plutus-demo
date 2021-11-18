import { ConditionalExpr } from '@angular/compiler';

export interface WalletMap {
  [key: string]: string;
}

export interface WalletDto {
  poolWallet:Pool;
  coinWallet: CoinWallet;
}

export interface CoinWallet {
  walletName: string;
  walletUuid: string;
  tokens?: Token[];
}

export interface Pool {
  tokens: Token[];
}

export interface Token {
  currencySymbol: string;
  tokenName: string;
  amount: number;
}

export interface SwapDto {
  swapCoin: Coin;
  toCoinTokenName: string;
}

export interface RemoveDto {
  tokenNameCoinA: string;
  tokenNameCoinB: string;
  removeDiff: number;
}

export interface CloseDto {
  tokenNameCoinA: string;
  tokenNameCoinB: string;
}


export interface FundsDto {
  coinA: Coin;
  coinB: Coin;
}

export interface Coin {
  tokenName: string;
  amount: number;
}
