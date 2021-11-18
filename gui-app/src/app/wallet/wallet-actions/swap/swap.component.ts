import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';
import { CoinWallet, SwapDto, Token, WalletDto } from 'src/app/model/model';

@Component({
  selector: 'app-swap',
  templateUrl: './swap.component.html',
  styleUrls: ['./swap.component.css'],
})
export class SwapComponent implements OnInit {
  @Input('wallet') wallet: CoinWallet | null = null;
  @Output() swapCoin: EventEmitter<SwapDto> = new EventEmitter();

  validity$: Observable<string | null> = of();

  inputForm: FormGroup = new FormGroup({});

  availableTokens: string[] = [];

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.inputForm = this.fb.group({
      swapCoinAmount: ['', Validators.required],
      swapCoinName: ['', Validators.required],
      toCoinTokenName: ['', Validators.required],
    });
    this.validity$ = this.inputForm.valueChanges.pipe(
      map((val) => {
        return this.isValid(val);
      })
    );
  }

  ngOnChanges(): void {
    if (this.wallet && this.wallet.tokens) {
      this.availableTokens = this.wallet.tokens
        .filter((token: Token) => token.currencySymbol && token.tokenName)
        .map((token) => token.tokenName);
    }
  }

  isValid(val: any): string | null {
    const funds = this.convertToFormValue(val);
    if (!funds.swapCoin.tokenName || !funds.toCoinTokenName) {
      return 'Not all Tokens selected';
    }
    if (funds.swapCoin.tokenName === funds.toCoinTokenName) {
      return 'Swapcoin and to Coin token should not be the same';
    }

    const correspondingSwapToken = this.wallet?.tokens?.find(
      (token) => token.tokenName === funds.swapCoin.tokenName
    );

    if (!correspondingSwapToken || !correspondingSwapToken) {
      return 'Coins not found in Wallet';
    }
    if (correspondingSwapToken.amount < funds.swapCoin.amount
    ) {
      return 'Not enough funds to spend';
    }
    return null;
  }

  sendFunds(): void {
    this.swapCoin.emit(this.convertToFormValue(this.inputForm.value));
    this.inputForm.reset();
  }

  private convertToFormValue(val: any): SwapDto {
    return {
      swapCoin: {
        amount: val.swapCoinAmount,
        tokenName: val.swapCoinName,
      },
      toCoinTokenName: val.toCoinTokenName,
    };
  }
}
