import { Injectable } from '@angular/core';
import {HttpService} from "../http.service";
import {Subject} from "rxjs/Subject";
import {AlertService} from "../alert/alert.service";
import {Data} from "../data-editor/data-editor.service";

/**
 * single editable property
 */
export class ConfigurationProperty {
  key:String;
  value:any;
}

/**
 * The object the service and component handle
 */
export class Configuration {
  id:String;
  properties:ConfigurationProperty;
}

/**
 *
 */
@Injectable()
export class TestService {

  objectUrl = "test";

  /**
   * Constructor
   *
   * @param httpService
   */
  constructor(private httpService:HttpService, private _alertService:AlertService) {
  }



}
