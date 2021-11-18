import { ChangeDetectionStrategy, Component, OnInit } from '@angular/core';
import { Observable, of, ReplaySubject, Subject } from 'rxjs';
import { StateService } from './service/state.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent implements OnInit {
  loading$: Observable<boolean> = of();

  constructor(private stateService: StateService){

  }

  ngOnInit() : void {
   this.loading$ = this.stateService.isLoading();

  }

  
}
