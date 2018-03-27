import { Component, OnInit, Input, Output } from '@angular/core';

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

  alert = "Testing";
  showIt = false;

  ngOnInit() {
    // copy all inputs to avoid polluting them
    console.log("Alert::ngOnInit");
  }

  close() {
    console.log("Alert::close");
  }
}
