import { Component, Input, OnInit } from '@angular/core';
import {Token} from '../../model/model';

@Component({
  selector: 'app-token-overview',
  templateUrl: './token-overview.component.html',
  styleUrls: ['./token-overview.component.css']
})
export class TokenOverviewComponent implements OnInit {

  @Input('token') token: Token | null = null;

  constructor() { }

  ngOnInit(): void {
  }

}
