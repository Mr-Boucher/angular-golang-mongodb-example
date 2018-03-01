import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import "rxjs/add/operator/map";
import {Subject} from "rxjs/Subject";


export interface Data {
  id: string;
  value: string;
}

const httpOptions = {
  headers: new HttpHeaders({
    // 'Content-Type': 'application/json',
    // "Access-Control-Allow-Origin": "*",
    // "Access-Control-Allow-Methods": "GET",
    // "Access-Control-Allow-Headers": "Content-Type"
  })
};

@Injectable()
export class JsonViewerService {

  // courseUrl = "https://angular-http-guide.firebaseio.com/courses.json";
  host = "http://localhost:8000/";
  getDataUrl = "loaddata";
  deleteDataUrl = "deletedata";

  subject: Subject<Data[]> = new Subject();
  _data: Data[] = [];

  get data() {
    return this.subject.asObservable();
  }

  constructor(private httpClient: HttpClient) {
    this.updateData()
  }

  updateData() {
    console.log( "updateData" );
    this.httpClient.get<Data[]>(this.host + this.getDataUrl, httpOptions).subscribe(data => {
      this._data = <Data[]>data; // save your data
      this.subject.next(this._data); // emit your data
    });
  }

  removeData(id: string) {
    console.log( "removing: " + id );
    this.httpClient.delete( this.host + this.deleteDataUrl ).subscribe( data=>{
      for ( let index = 0; index < this._data.length; index++ ) {
        this._data.splice(index, 1); //remove 1 item
      }
      this.subject.next(this._data); // emit your data
    });
  }
}

