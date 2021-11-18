import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';
import { Pool, RemoveDto, Token } from 'src/app/model/model';

@Component({
  selector: 'app-remove',
  templateUrl: './remove.component.html',
  styleUrls: ['./remove.component.css']
})
export class RemoveComponent implements OnInit {

  @Input('pool') pool: Pool | null = null;
  @Output() removeFunds: EventEmitter<RemoveDto> = new EventEmitter();

  validity$: Observable<string | null> = of();

  inputForm: FormGroup = new FormGroup({});

  availableTokens: string[] = [];

  constructor(private fb: FormBuilder) {}

  ngOnInit(): void {
    this.inputForm = this.fb.group({
      tokenNameCoinA: ['', Validators.required],
      tokenNameCoinB: ['', Validators.required],
      removeDiff: ['', Validators.required],
    });
    this.validity$ = this.inputForm.valueChanges.pipe(
      map((val) => {
        return this.isValid(val);
      })
    );
  }

  ngOnChanges(): void {
    if (this.pool && this.pool.tokens) {
      this.availableTokens = this.pool.tokens
        .filter((token: Token) => token.currencySymbol && token.tokenName)
        .map((token) => token.tokenName);
    }
  }

  isValid(val: any): string | null {
    const funds = this.convertToFormValue(val);
    if (!funds.tokenNameCoinA || !funds.tokenNameCoinB) {
      return 'Not all Tokens selected';
    }
    if (funds.tokenNameCoinA === funds.tokenNameCoinB) {
      return 'First and second token should not be the same';
    }

    const correspondingCoinA = this.pool?.tokens?.find(
      (token) => token.tokenName === funds.tokenNameCoinA
    );
    const correspondingCoinB = this.pool?.tokens?.find(
      (token) => token.tokenName === funds.tokenNameCoinB
    );
    if (!correspondingCoinA || !correspondingCoinB) {
      return 'Coins not found in Pool';
    }
    return null;
  }

  remove(): void {
    this.removeFunds.emit(this.convertToFormValue(this.inputForm.value));
    this.inputForm.reset();
  }

  private convertToFormValue(val: any): RemoveDto {
    return {
      tokenNameCoinA: val.tokenNameCoinA,
      tokenNameCoinB: val.tokenNameCoinB,
      removeDiff: val.removeDiff
    }
  }
}
