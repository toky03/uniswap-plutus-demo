import { Pipe, PipeTransform } from '@angular/core';
import { Token } from '../model/model';

@Pipe({
  name: 'sort'
})
export class SortPipe implements PipeTransform {

  transform(wallets: string[] | null): string[] {
  if(!wallets){
    return [];
  }
  return wallets.slice(0).sort(sortNumber)

}
}

@Pipe({
  name: 'sortToken'
})
export class SortTokenPipe implements PipeTransform {

  transform(tokens: Token[] | undefined): Token[] {
  if(!tokens){
    return [];
  }
  return tokens.slice(0).sort((a: Token, b: Token) => (''+a.tokenName).localeCompare(b.tokenName));

}
}

const sortNumber = (a: string, b: string) => {
    
  if(isNaN(+a)){
    return -1;
  } else if(isNaN(+b)){
    return 1;
  }
  return +a - +b
};
