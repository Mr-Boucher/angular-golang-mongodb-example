import { Injectable } from '@angular/core';
import {HttpService} from "../http.service";
import {Subject} from "rxjs/Subject";

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
export class ConfigurationService {

  configurationUrl = "configuration/:id";

  subject:Subject<Configuration[]> = new Subject();
  _configurations:Configuration[] = []; //Make sure it is defaulted to an empty array else it will be undefined causing errors


  /**
   * Constructor
   *
   * @param httpService
   */
  constructor(private httpService:HttpService) {
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Configuration[]>}
   */
  get configurations() {
    return this.subject.asObservable();
  }

  /**
   *
   * @param configuration
   */
  update(configuration:Configuration):void {
    this.httpService.update(configuration, this.configurationUrl, this.subject, this._configurations);
  }

  load():void {

  }
}
