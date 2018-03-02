import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import "rxjs/add/operator/map";
import {Subject} from "rxjs/Subject";


export class Data {
  id: String;
  value: String;
}

const httpOptions = {
  headers: new HttpHeaders({
    'Accept': 'application/json', //only accept json responses
    'Content-Type': 'application/json', //set the sending data as json
    //'Access-Control-Request-Method': 'GET, POST, PUT, DELETE, OPTIONS',
    //'Access-Control-Request-Origin': '*'
  })
};

@Injectable()
export class JsonViewerService {

  host = "http://localhost:8000/";
  objectUrl = "data";
  deleteUrl = this.objectUrl + "/";

  subject:Subject<Data[]> = new Subject();
  _data:Data[] = [];

  get data() {
    return this.subject.asObservable();
  }

  constructor(private httpClient:HttpClient) {
    this.load()
  }

  //
  load() {
    console.log("load data");
    this.httpClient.get<Data[]>(this.host + this.objectUrl, httpOptions).subscribe(data => {
      this._data = <Data[]>data; // save your data
      this.subject.next(this._data); // emit your data
    });
  }

  //
  add(value: String) {
    let newData = new Data();
    newData.value = value;
    console.log( "adding data: " + newData);
    let json = JSON.stringify(newData);
    console.log( "adding data: " + json);
    this.httpClient.post(this.host + this.objectUrl, json, httpOptions).subscribe(data => {
      //this.data
      this._data.push( data ) // save your data
      this.subject.next(this._data); // emit your data
      console.log( "added data: " + data)
    });
  }

  //
  remove(id:string) {
    console.log("deleting data(" + id + ")");
    this.httpClient.delete(this.host + this.deleteUrl + id, httpOptions).subscribe(data=> {
      for (let index = 0; index < this._data.length; index++) {
        if( this._data[index].id == id ) {
          this._data.splice(index, 1); //remove 1 item
          this.subject.next(this._data); // emit your data
        }
      }
    });
  }
}

