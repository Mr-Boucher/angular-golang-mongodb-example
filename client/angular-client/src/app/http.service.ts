import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Data} from "./data-editor/data-editor.service";
import {Subject} from "rxjs/Subject";

//HttpOptions are needed to make sure that all REST API pass basic security as well as browser CORS
const httpOptions = {
  headers: new HttpHeaders({
    'Accept': 'application/json', //only accept json responses
    'Content-Type': 'application/json', //set the sending data as json
    //'Access-Control-Request-Method': 'GET, POST, PUT, DELETE, OPTIONS',
    //'Access-Control-Request-Origin': '*'
  })
};

/**
 *
 */
@Injectable()
export class HttpService {

  host = "http://localhost:8000/";

  constructor(private httpClient:HttpClient) { }

  //Call the REST API, add the data to the array and update the subject
  load( objectUrl:String, subject: Subject<any>, dataArray:any[] ) {
    this.httpClient.get<Data[]>(this.host + objectUrl, httpOptions).subscribe(data => {
      console.log( "Received " + data );
      // dataArray.splice(0, dataArray.length -1);
      dataArray.concat( data );
      // data.forEach(function (value) {
      //   console.log(value);
      //   dataArray.concat();
      // });

      subject.next(dataArray); // emit your data
    });
  }

  //Add a new element to the array and update the server with the new data
  add(object: any, objectUrl:String, subject: Subject<any>, dataArray:any[] ) {
    console.log( "adding data: " + object);
    let json = JSON.stringify(object);
    console.log( "adding data: " + json );
    this.httpClient.post<Data>(this.host + objectUrl, json, httpOptions).subscribe(data => {
      console.log( dataArray);
      dataArray.push( data ); // save your data
      subject.next(dataArray); // emit your data
      console.log( "added data: " + data + " to " + dataArray);
    });
  }
}
