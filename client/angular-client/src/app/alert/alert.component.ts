import { Component, OnInit, Input, Output } from '@angular/core';
import {AlertService} from "./alert.service";

/**
 *
 */
@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css'],
  providers: [AlertService]
})
export class AlertComponent implements OnInit {

  errors: String = "";

  constructor( private _alertService: AlertService ) {
  }

  /**
   *
   */
  ngOnInit() {
    console.log("Alert::ngOnInit");
    this._alertService.errors.subscribe(
      errors => {
        this.errors = errors;
        console.log("Alert::Alert:" + this.errors);
      },
      err => {
        console.log("Alert::Error:" + err);
        console.error(err);
      },
      () => {
        console.log("done loading");
      }
    );
  }

  close() {
    console.log("Alert::close");
    this._alertService.clearErrors();
  }
}
