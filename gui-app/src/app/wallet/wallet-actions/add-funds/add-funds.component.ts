import { tokenName } from '@angular/compiler';
import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Input,
  OnChanges,
  OnInit,
  Output,
} from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';
import { FundsDto, Token, CoinWallet } from 'src/app/model/model';

@Component({
  selector: 'app-add-funds',
  templateUrl: './add-funds.component.html',
  styleUrls: ['./add-funds.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AddFundsComponent implements OnInit, OnChanges {
  @Input('wallet') wallet: CoinWallet | null = null;
  @Input('hasMaxAmount') hasMaxAmount: boolean = false;
  @Output() addFunds: EventEmitter<FundsDto> = new EventEmitter();

  validity$: Observable<string | null> = of();

  inputForm: FormGroup = new FormGroup({});

  availableTokens: string[] = [];

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.inputForm = this.fb.group({
      coinAName: ['', Validators.required],
      coinBName: ['', Validators.required],
      coinAAmount: ['', Validators.required],
      coinBAmount: ['', Validators.required],
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
    console.log('trigger validation')
    const funds = this.convertToFormValue(val);
    if (!funds.coinA.tokenName || !funds.coinB.tokenName) {
      return 'Not all Tokens selected';
    }
    if (funds.coinA.tokenName === funds.coinB.tokenName) {
      return 'First and second token should not be the same';
    }

    const correspondingCoinA = this.wallet?.tokens?.find(
      (token) => token.tokenName === funds.coinA.tokenName
    );
    const correspondingCoinB = this.wallet?.tokens?.find(
      (token) => token.tokenName === funds.coinB.tokenName
    );
    if (!correspondingCoinA || !correspondingCoinB) {
      return 'Coins not found in Wallet';
    }
    if (
      this.hasMaxAmount &&
      (correspondingCoinA.amount < funds.coinA.amount ||
        correspondingCoinB.amount < funds.coinB.amount)
    ) {
      return 'Not enough funds to spend';
    }
    return null;
  }

  sendFunds(): void {
    this.addFunds.emit(this.convertToFormValue(this.inputForm.value));
    this.inputForm.reset();
  }

  private convertToFormValue(val: any): FundsDto {
    return {
      coinA: {
        amount: val.coinAAmount,
        tokenName: val.coinAName,
      },
      coinB: {
        amount: val.coinBAmount,
        tokenName: val.coinBName,
      },
    };
  }
}
