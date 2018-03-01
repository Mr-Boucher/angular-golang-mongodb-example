import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import "rxjs/add/operator/map";
import {Subject} from "rxjs/Subject";


export interface Data {
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
  courseUrl = "http://localhost:8000/loaddata";

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
    this.httpClient.get<Data[]>(this.courseUrl, httpOptions).subscribe(data => {
      this._data = <Data[]>data; // save your data
      this.subject.next(this._data); // emit your data
    });
  }
}

