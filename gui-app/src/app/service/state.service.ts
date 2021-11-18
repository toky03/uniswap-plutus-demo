import { Injectable } from '@angular/core';
import {
  BehaviorSubject,
  combineLatest,
  Observable,
  of,
  ReplaySubject,
  Subject,
} from 'rxjs';
import { switchMap, take, tap } from 'rxjs/operators';
import {
  CloseDto,
  FundsDto,
  Pool,
  RemoveDto,
  SwapDto,
  WalletDto,
} from '../model/model';
import { IntegrationService } from './integration.service';

@Injectable({
  providedIn: 'root',
})
export class StateService {
  wallets$: Subject<string[]> = new ReplaySubject(1);
  loading$: Subject<boolean> = new BehaviorSubject<boolean>(true);
  currentWallet$: Subject<WalletDto> = new Subject();

  constructor(private integrationService: IntegrationService) {}


  public readWallets(): void {
    console.log('start loading')
    this.loading$.next(true);
    this.integrationService.readWallets().subscribe(
      (wallets: string[]) => {
        this.wallets$.next(wallets);
        this.loading$.next(false);
      },
      (error: any) => {
        this.loading$.next(false);
      }
    );
  }
  public reloadFunds(walletId: string): void {
    this.loading$.next(true);
    this.integrationService.readFunds(walletId).subscribe(
      (wallet: WalletDto) => {
        this.currentWallet$.next(wallet);
        this.loading$.next(false);
      },
      (error: any) => {
        this.loading$.next(false);
      }
    );
  }

  public isLoading(): Observable<boolean> {
    return this.loading$.asObservable();
  }

  public startLoading(): void {
    this.loading$.next(true);
  }

  addFunds(walletId: string, fundsDto: FundsDto) {
    this.integrationService
      .addFunds(walletId, fundsDto)
      .subscribe(() => this.reloadFunds(walletId));
  }
  createPool(walletId: string, fundsDto: FundsDto) {
    this.integrationService
      .createPool(walletId, fundsDto)
      .subscribe(() => this.reloadFunds(walletId));
  }

  swap(walletId: string, swapDto: SwapDto) {
    this.integrationService
      .swapCoin(walletId, swapDto)
      .subscribe(() => this.reloadFunds(walletId));
  }

  removeFunds(walletId: string, removeDto: RemoveDto) {
    this.integrationService
      .removeFunds(walletId, removeDto)
      .subscribe(() => this.reloadFunds(walletId));
  }

  closePool(walletId: string, closeDto: CloseDto) {
    this.integrationService
      .closePool(walletId, closeDto)
      .subscribe(() => this.reloadFunds(walletId));
  }
}
