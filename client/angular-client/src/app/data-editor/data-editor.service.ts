import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import "rxjs/add/operator/map";
import {Subject} from "rxjs/Subject";
import { HttpService} from "../http.service";


export class Data {
  id: String;
  value: String;
}


@Injectable()
export class DataEditorService {

  objectUrl = "data";
  deleteUrl = this.objectUrl + "/";

  subject:Subject<Data[]> = new Subject();
  _data:Data[] = [];

  get data() {
    return this.subject.asObservable();
  }

  constructor(private httpService: HttpService) {
    this.load();
  }

  //
  load() {
    console.log("load data");
    this.httpService.load( this.objectUrl, this.subject, this._data);
  }

  //
  add(value: String) {

    //create the data object
    let newData = new Data();
    newData.value = value; //only set the value because the Id is created on the server
    this.httpService.add( newData, this.objectUrl, this.subject, this._data );
  }

  //
  remove(id:string) {
    this.httpService.remove( id, this.deleteUrl, this.subject, this._data );
  }
}

