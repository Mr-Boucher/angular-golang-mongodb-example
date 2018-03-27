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

  alert: String = "";

  constructor( private _alertService: AlertService ) {
  }

  /**
   *
   */
  ngOnInit() {
    console.log("Alert::ngOnInit");
    this._alertService.errors.subscribe(
      alert => {
        this.alert = alert;
        console.log("subscribe result");
      },
      err => {
        console.error(err);
      },
      () => {
        console.log("done loading");
      }
    );
  }

  close() {
    console.log("Alert::close");
  }

  hasErrors(): boolean {
    return alert.length > 0
  }
}
