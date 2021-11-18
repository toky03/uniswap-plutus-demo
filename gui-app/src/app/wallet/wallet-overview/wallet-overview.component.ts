import {
  ChangeDetectionStrategy,
  Component,
  OnDestroy,
  OnInit,
} from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { combineLatest, Observable, of, Subject } from 'rxjs';
import { map, switchMap, takeUntil } from 'rxjs/operators';
import {
  CloseDto,
  FundsDto,
  Pool,
  RemoveDto,
  SwapDto,
  WalletDto,
} from 'src/app/model/model';
import { StateService } from 'src/app/service/state.service';

@Component({
  selector: 'app-wallet-overview',
  templateUrl: './wallet-overview.component.html',
  styleUrls: ['./wallet-overview.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WalletOverviewComponent implements OnInit, OnDestroy {
  private destroy$: Subject<void> = new Subject();

  private walletId$: Observable<string> = of('');
  wallet$: Observable<WalletDto> = of();
  pool$: Observable<Pool> = of();
  availableTokens$: Observable<string[]> = of();

  constructor(
    private activatedRoute: ActivatedRoute,
    private router: Router,
    private service: StateService
  ) {}

  ngOnInit(): void {
    console.log('init walled overview')
    this.activatedRoute.params
      .pipe(map((p) => p.id))
      .subscribe((walletId) => this.service.reloadFunds(walletId));

    this.wallet$ = this.service.currentWallet$;
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  navigateToHome(): void {
    this.router.navigate(['/']);
  }

  onCreatePool(walletId: string, fundsDto: FundsDto): void {
    this.service.createPool(walletId, fundsDto);
  }

  onAddFunds(walletId: string, fundsDto: FundsDto): void {
    this.service.addFunds(walletId, fundsDto);
  }

  onSwapCoin(walletId: string, swapDto: SwapDto): void {
    this.service.swap(walletId, swapDto);
  }

  onRemoveFunds(walletId: string, removeDto: RemoveDto): void {
    this.service.removeFunds(walletId, removeDto);
  }

  onClosePool(walletId: string, closeDto: CloseDto): void {
    this.service.closePool(walletId, closeDto);
  }
}
