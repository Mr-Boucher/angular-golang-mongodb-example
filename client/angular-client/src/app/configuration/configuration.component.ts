import { Component, OnInit } from '@angular/core';
import {ConfigurationService} from "./configuration.service";
import {Configuration} from "./configuration.service";
import {selector} from "rxjs/operator/publish";
import {Subject} from "rxjs/Subject";


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
  colorList:HTMLSelectElement;
  webconfigurations:HTMLElement;
  subject:Subject<HTMLSelectElement> = new Subject();
  divWatch:Subject<HTMLElement> = new Subject();
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
  updateBackgroundColor(color:string) {
    console.log("hitting background color");
    console.log("selected color is " + color);
    var divElement = document.getElementById("webconfigurations");
    divElement.style.backgroundColor = color;
    this.webconfigurations = divElement;
    this.divWatch.next(this.webconfigurations);
    console.log("color has been updated");

  }
}
