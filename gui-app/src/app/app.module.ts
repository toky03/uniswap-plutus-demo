import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule} from '@angular/common/http';
import {MatGridListModule} from '@angular/material/grid-list';
import { MatCardModule } from '@angular/material/card';
import {MatButtonModule} from '@angular/material/button'
import {MatInputModule} from '@angular/material/input';
import {MatIconModule} from '@angular/material/icon';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatSelectModule} from '@angular/material/select';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatTabsModule} from '@angular/material/tabs';


import { SortPipe, SortTokenPipe } from './service/sort.pipe';
import { WalletOverviewComponent } from './wallet/wallet-overview/wallet-overview.component';
import { WalletListComponent } from './wallet-list/wallet-list.component';
import { TokenOverviewComponent } from './wallet/token-overview/token-overview.component';
import { MatOptionModule } from '@angular/material/core';
import { AddFundsComponent } from './wallet/wallet-actions/add-funds/add-funds.component';
import { ReactiveFormsModule } from '@angular/forms';
import { SwapComponent } from './wallet/wallet-actions/swap/swap.component';
import { RemoveComponent } from './wallet/wallet-actions/remove/remove.component';
import { ClosePoolComponent } from './wallet/wallet-actions/close-pool/close-pool.component';


@NgModule({
  declarations: [
    AppComponent,
    SortPipe,
    SortTokenPipe,
    WalletOverviewComponent,
    WalletListComponent,
    TokenOverviewComponent,
    AddFundsComponent,
    SwapComponent,
    RemoveComponent,
    ClosePoolComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatGridListModule,
    MatCardModule,
    MatButtonModule,
    MatInputModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatOptionModule,
    MatSelectModule,
    MatTooltipModule,
    MatTabsModule,
    ReactiveFormsModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
