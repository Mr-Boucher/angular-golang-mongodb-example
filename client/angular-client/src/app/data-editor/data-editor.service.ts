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
  deleteUrl = this.objectUrl + "/";

  subject:Subject<Data[]> = new Subject();
  _data:Data[] = []; //Make sure it is defaulted to an empty array else it will be undefined causing errors

  /**
   *
   * @param httpService
   */
  constructor(private httpService:HttpService, private _alertService:AlertService) {
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Data[]>}
   */
  get data():Observable<Data[]> {
    return this.subject.asObservable();
  }

  /**
   *
   */
  search( searchCriteria:string, pageNumber:number ):Observable<Data[]> {
    console.log("Search Criteria:" + searchCriteria );
    console.log("Search pageNumber:" + pageNumber );
    let url:string = this.objectUrl + "?pageNumber=" + pageNumber + "&search=";
    if( searchCriteria != null )
      url += searchCriteria;
    this.httpService.load( url, this.subject, this._alertService, this._data );
    return this.subject.asObservable();
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

