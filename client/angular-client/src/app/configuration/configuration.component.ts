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
  providers: [ConfigurationService]
})
export class ConfigurationComponent implements OnInit {

  configurations:Configuration[];

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
  update(data, $event) {
    this._configurationService.update(data)
  }
}
