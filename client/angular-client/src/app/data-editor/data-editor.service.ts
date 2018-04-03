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

export class DataSet {
  data:Data[] = [];
  totalCount:number = 0;
}

/**
 * Supports the CRUD of data objects
 */
@Injectable()
export class DataEditorService {

  objectUrl = "data";
  deleteUrl = this.objectUrl + "/";

  subject:Subject<DataSet> = new Subject();
  _data:DataSet = new DataSet(); //Make sure it is defaulted to an empty array else it will be undefined causing errors

  /**
   *
   * @param httpService
   */
  constructor(private httpService:HttpService, private _alertService:AlertService) {
    console.log("DataEditorService Constructor:" );
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Data[]>}
   */
  get data():Observable<DataSet> {
    return this.subject.asObservable();
  }

  /**
   *
   */
  search( searchCriteria:string, pageNumber:number ):Observable<DataSet> {
    console.log("Search Criteria:" + searchCriteria );
    console.log("Search pageNumber:" + pageNumber );
    let url:string = this.objectUrl + "?pageNumber=" + pageNumber + "&search=";
    if( searchCriteria != null )
      url += searchCriteria;
    this.httpService.load( url, this.handleResult.bind( this ) );
    return this.subject.asObservable();
  }

  /**
   *
   * @param data
   */
  handleResult( data:any ):void {
    console.log("DataEditorService handleResult:" + data['total_count'] );
    this._data.totalCount = data['total_count'];
    this._data.data = [];
    for( let obj of data['data_set'] )
    {
      let newObj:Data = new Data();
      newObj.id = data['id'];
      newObj.value = data['value'];
      this._data.data.push( newObj );
    }
    this._data.data = data['data_set'];

    //Emit the data to the subject so the data will refresh with the new value set
    this.subject.next(this._data);
  }

  /**
   *
   * @param value
   */
  add(value:String):void {

    //create the data object
    let newData = new Data();
    newData.value = value; //only set the value because the Id is created on the server
    this.httpService.add(newData, this.objectUrl, this.subject, this._data.data);
  }

  /**
   *
   * @param id
   */
  remove(id:string):void {
    this.httpService.remove(id, this.deleteUrl, this.subject, this._data.data);
  }
}

