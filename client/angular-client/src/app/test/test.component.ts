import { Component, OnInit } from '@angular/core';
import {TestService} from "./test.service";


/**
 *
 */
@Component({
  selector: 'app-configuration',
  templateUrl: './test.component.html',
  styleUrls: ['./test.component.css'],
  providers: []
})
export class TestComponent implements OnInit {

  /**
   *
   * @param _configurationService
   */
  constructor(private _configurationService:TestService) {
  }

  /**
   *
   */
  ngOnInit() {
    console.log("ngOnInit");

  }

}
