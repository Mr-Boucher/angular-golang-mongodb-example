import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import "rxjs/add/operator/map";
import {Subject} from "rxjs/Subject";
import {HttpService} from "../http.service";
import {Observable} from "rxjs/Observable";
import {catchError} from "rxjs/operators";
import {AlertService} from "../alert/alert.service";

/**
 * The object the service and component handle
 */
export class Data {
  id:String;
  value:String;
}

/**
 * Supports the CRUD of data objects
 */
@Injectable()
export class DataEditorService {

  objectUrl = "data";
  searchUrl = this.objectUrl + "?search=";
  deleteUrl = this.objectUrl + "/";

  subject:Subject<Data[]> = new Subject();
  _data:Data[] = []; //Make sure it is defaulted to an empty array else it will be undefined causing errors

  /**
   *
   * @param httpService
   */
  constructor(private httpService:HttpService, private _alertService:AlertService) {
    this.load();
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Data[]>}
   */
  get data() {
    return this.subject.asObservable();
  }

  /**
   *
   */
  search( searchCriteria:string ):void {
    console.log("Search Criteria" + searchCriteria );
    let sc = new Array(searchCriteria);
    this.httpService.load(this.searchUrl + sc, this.subject, this._alertService, this._data );
  }

  /**
   *
   */
  load( ):void {
    console.log("Load" );
    this.httpService.load(this.objectUrl, this.subject, this._alertService, this._data );
  }

  /**
   *
   * @param value
   */
  add(value:String):void {

    //create the data object
    let newData = new Data();
    newData.value = value; //only set the value because the Id is created on the server
    this.httpService.add(newData, this.objectUrl, this.subject, this._data);
  }

  /**
   *
   * @param id
   */
  remove(id:string):void {
    this.httpService.remove(id, this.deleteUrl, this.subject, this._data);
  }
}

