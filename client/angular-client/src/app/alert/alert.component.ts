import { Component, OnInit, Input, Output } from '@angular/core';
import {AlertService} from "./alert.service";

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

  private _errors: String = "Uknonw";

  constructor( private _alertService: AlertService ) {
  }

  get errors(): String{
    return this._errors;
  }

  /**
   *
   */
  ngOnInit() {
    console.log("Alert::ngOnInit");
    this._alertService.errorsObserver.subscribe(
      errors => {
        console.log("Alert::Alert errors:" + errors);
        this._errors = "Balh";
        console.log("Alert::Alert:" + this._errors);
      },
      err => {
        console.log("Alert::Error:" + err);
        console.error(err);
      },
      () => {
        console.log("Alert::Alert Done");
      }
    );

    this._errors = this._alertService.errors;
  }

  close() {
    console.log("Alert::close");
    this._alertService.clearErrors();
  }
}
