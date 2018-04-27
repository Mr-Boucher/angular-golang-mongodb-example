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
export class ConfigurationService {

  objectUrl = "configuration";

  subject:Subject<Configuration[]> = new Subject();
  _configurations:Configuration[] = []; //Make sure it is defaulted to an empty array else it will be undefined causing errors


  /**
   * Constructor
   *
   * @param httpService
   */
  constructor(private httpService:HttpService, private _alertService:AlertService) {
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
  update(configuration:Configuration, value:String):void {

    let newConfig = new ConfigurationProperty();
    newConfig.value = value; //only set the value because the Id is created on the server
    configuration.properties = newConfig;
    this.httpService.update(configuration, this.objectUrl, this.updateItem.bind(this));
  }

  load():void {
    console.log("load data");
    this.httpService.load(this.objectUrl, this.retrieve.bind( this ));
  }

  /**
   *
   * @param result
   */
  retrieve( result:any ):void {
    console.log("Configuration retrieve" );

    this._configurations = [];
    for (var obj in result)
    {
      console.log("Configuration retrieve data:" + obj + obj['id'] + "," + obj['properties']);
      let newObj:Configuration = new Configuration();
      newObj.id = obj;
      newObj.properties = result[obj];
      this._configurations.push(newObj);
    }

    //Emit the data to the subject so the data will refresh with the new value set
    this.subject.next(this._configurations);

  }

  /**
   *
   * @param {Configuration} item
   */
  updateItem(item:Configuration) {
      console.log("Update item on the client");

      console.log("Retrieved updated configuration");
  }

  updateBackgroundColor(color:String) {

  }
}
