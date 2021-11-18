import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, of } from 'rxjs';
import { StateService } from '../service/state.service';

@Component({
  selector: 'app-wallet-list',
  templateUrl: './wallet-list.component.html',
  styleUrls: ['./wallet-list.component.css'],
})
export class WalletListComponent implements OnInit {
  wallets$: Observable<string[]> = of();
  loading$: Observable<boolean> = of();

  constructor(private stateService: StateService, private router: Router) {}

  ngOnInit(): void {
    this.stateService.readWallets();
    this.wallets$ = this.stateService.wallets$;
    this.loading$ = this.stateService.isLoading();
  }

  trackByWalletName(index: number, wallet: string): string {
    return wallet;
  }

  navigateToWallet(wallet: string): void {
    this.stateService.startLoading();
    this.router.navigate(['wallet', wallet]);
  }
}
