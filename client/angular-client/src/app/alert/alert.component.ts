import { Component, OnInit, Input, Output } from '@angular/core';
import {AlertService} from "./alert.service";
import {AppError} from "./alert.service";

/**
 *
 */
@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css'],
  providers: []
})
export class AlertComponent implements OnInit {

  constructor( private _alertService: AlertService ) {
  }

  get errors(): AppError[]{
    return this._alertService.errors;
  }

  hasErrors(): boolean  {
    return this._alertService.hasErrors();
  }

  /**
   *
   */
  ngOnInit() {
    console.log("Alert::ngOnInit");
  }

  close() {
    console.log("Alert::close");
    this._alertService.clearErrors();
  }
}
