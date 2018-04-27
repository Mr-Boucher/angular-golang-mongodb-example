import { Component, OnInit } from '@angular/core';
import {ConfigurationService} from "./configuration.service";
import {Configuration} from "./configuration.service";


/**
 *
 */
@Component({
  selector: 'app-configuration',
  templateUrl: './configuration.component.html',
  styleUrls: ['./configuration.component.css'],
  providers: []
})
export class ConfigurationComponent implements OnInit {

  configurations:Configuration[];

  colors:String[] = [
    "pink",
    "green",
    "blue"
  ];

  /**
   *
   * @param _configurationService
   */
  constructor(private _configurationService:ConfigurationService) {
  }

  /**
   *
   */
  ngOnInit() {
    console.log("ngOnInit");
    this._configurationService.configurations.subscribe(
      data => {
        this.configurations = data;
        console.log("subscribe result")
      },
      err => console.error(err),
      () => console.log('done loading courses')
    );
  }

  /**
   *
   * @param $event
   */
  refresh($event) {
    this._configurationService.load();
  }

  /**
   *
   * @param data
   * @param $event
   */
  update(data, value:String, $event) {
    this._configurationService.update(data, value);
  }

  /**
   *
   */
  updateBackgroundColor() {
    console.log("hitting background color");
    console.log(this.colors);
  }
}
