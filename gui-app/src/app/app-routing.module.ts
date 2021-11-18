import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { WalletListComponent } from './wallet-list/wallet-list.component';
import { WalletOverviewComponent } from './wallet/wallet-overview/wallet-overview.component';

const routes: Routes = [
  {
    path: 'wallet',
    children: [
      {
        path: ':id',
        component: WalletOverviewComponent
      },
      {
        path: '',
        component: WalletListComponent
      }
    ]
  },
  {
    path: '**',
    component: WalletListComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
