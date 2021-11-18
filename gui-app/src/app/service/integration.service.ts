import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { CloseDto, FundsDto, Pool, RemoveDto, SwapDto, WalletDto } from '../model/model';

const API_ROOT = environment.apiBaseUrl;
@Injectable({
  providedIn: 'root'
})
export class IntegrationService {

  constructor(private httpClient: HttpClient) { }


  public readWallets(): Observable<string[]> {
    return this.httpClient.get<string[]>(`${API_ROOT}/wallets`)
  }
  public readFunds(id: string): Observable<WalletDto> {
    return this.httpClient.get<WalletDto>(`${API_ROOT}/${id}/funds`, {responseType: 'json'});
  }



  public addFunds(id: string, fundsDto: FundsDto): Observable<void> {
    return this.httpClient.put<void>(`${API_ROOT}/${id}/add`, fundsDto);
  }
  public createPool(id: string, fundsDto: FundsDto): Observable<void> {
    return this.httpClient.post<void>(`${API_ROOT}/${id}/create`, fundsDto);
  }

  public swapCoin(id: string, swapDto: SwapDto): Observable<void> {
    return this.httpClient.put<void>(`${API_ROOT}/${id}/swap`, swapDto);
  }

  public closePool(id: string, closeDto: CloseDto): Observable<void> {
    return this.httpClient.put<void>(`${API_ROOT}/${id}/close`, closeDto);
  }

  public removeFunds(id: string, removeDto: RemoveDto) {
    return this.httpClient.put<void>(`${API_ROOT}/${id}/remove`, removeDto);
  }

}
