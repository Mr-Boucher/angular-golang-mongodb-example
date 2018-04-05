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

export class Page {
  data:Data[] = [];
  state:String = "initial";
  totalCount:number = 0;
  pageNumber:number = 1;
  pageSize:number = 10;
  filter:string = null;
}

/**
 * Supports the CRUD of data objects
 */
@Injectable()
export class DataEditorService {

  objectUrl = "data";
  deleteUrl = this.objectUrl + "/";

  subject:Subject<Page> = new Subject();
  _data:Page = new Page(); //Make sure it is defaulted to an empty array else it will be undefined causing errors

  /**
   *
   * @param _httpService
   * @param _alertService
   */
  constructor(private _httpService:HttpService, private _alertService:AlertService) {
    console.log("DataEditorService Constructor:" );
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Data[]>}
   */
  get data():Observable<Page> {
    return this.subject.asObservable();
  }

  /**
   *
   */
  search( searchCriteria:Page ):Observable<Page> {
    console.log("Search filter:" + searchCriteria.filter );
    console.log("Search pageNumber:" + searchCriteria.pageNumber );
    console.log("Search pageSize:" + searchCriteria.pageSize );
    console.log("Search state:" + searchCriteria.state );
    let url:string = this.objectUrl + "?pageNumber=" + searchCriteria.pageNumber + "&pageSize=" + searchCriteria.pageSize + "&search=";

    //empty search filter is valid so make sure that it does not end up with null as the filter
    if( searchCriteria.filter != null )
      url += searchCriteria.filter;
    this._httpService.load( url, this.searchResult.bind( this ) );

    //
    searchCriteria.state = "Loading";

    //
    return this.subject.asObservable();
  }

  /**
   *
   * @param result
   */
  searchResult( result:any ):void {
    console.log("DataEditorService handleResult total_count:" + result['total_count'] + result['page_size'] );
    this._data.totalCount = result['total_count'];
    this._data.pageSize = result['page_size'];
    this._data.data = [];
    for( let obj of result['data_set'] )
    {
      console.log("DataEditorService handleResult data:" + obj['id'] + "," + obj['value'] + "" );
      let newObj:Data = new Data();
      newObj.id = obj['id'];
      newObj.value = obj['value'];
      this._data.data.push( newObj );
    }

    //Emit the data to the subject so the data will refresh with the new value set
    this.subject.next(this._data);

    //
    this._data.state = "Loaded";
  }

  /**
   *
   * @param value
   */
  add(value:String):void {

    //create the data object
    let newData = new Data();
    newData.value = value; //only set the value because the Id is created on the server
    this._httpService.add(newData, this.objectUrl, this.refreshSearch.bind( this ), this._data.data);
  }

  /**
   *
   * @param id
   */
  remove(id:string):void {
    this._httpService.remove(id, this.deleteUrl, this.refreshSearch.bind( this ), this._data.data);
  }

  /**
   *
   * @param data
   */
  refreshSearch( result:any ): void {
    this.search( this._data );
  }
}

