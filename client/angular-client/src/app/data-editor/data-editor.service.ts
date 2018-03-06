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
    let newData = new Data();
    newData.value = value;
    this.httpService.add( newData, this.objectUrl, this.subject, this._data );
  }

  //
  remove(id:string) {
    // console.log("deleting data(" + id + ")");
    // this.httpClient.delete(this.host + this.deleteUrl + id, httpOptions).subscribe(data=> {
    //   for (let index = 0; index < this._data.length; index++) {
    //     if( this._data[index].id == id ) {
    //       this._data.splice(index, 1); //remove 1 item
    //       this.subject.next(this._data); // emit your data
    //     }
    //   }
    // });
  }
}

