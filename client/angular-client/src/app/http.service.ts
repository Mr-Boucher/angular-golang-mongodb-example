import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Data} from "./data-editor/data-editor.service";
import {Subject} from "rxjs/subject";
import {AlertComponent} from "./alert/alert.component";
import {AlertService} from "./alert/alert.service";
import {AppError} from "./alert/alert.service";

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
 * Class for handling the configuration and communication with the REST API Server
 */
@Injectable()
export class HttpService {

  host = "http://localhost:8000/";

  constructor(private _httpClient:HttpClient, private _alertService:AlertService) {
  }

  /**
   * Call the REST API, add the data to the array and update the subject
   *
   * @param objectUrl
   * @param unmarshal
   */
  load(objectUrl:String, handleResult:Function) {
    this._httpClient.get<any>(this.host + objectUrl, httpOptions)
      .subscribe(result =>
      {
        console.log("HttpService::Load Received " + result);

        //invoke the call back/unmarshal method of the specialized service
        handleResult( result)
      },
      err => {
        console.log("HttpService::Loading Error");
        this.handleError(err);
      },
      () => {
        console.log("HttpService::Load Done");
      });
  }

  /**
   * Add a new element to the array and update the server with the new data
   *
   * @param object
   * @param objectUrl
   * @param handleResult
   * @param dataArray
   */
  add(object:any, objectUrl:String, handleResult:Function, dataArray:any[]) {
    console.log("adding data: " + object);
    let json = JSON.stringify(object); //convert object to JSON

    this._httpClient.post<Data>(this.host + objectUrl, json, httpOptions).subscribe(result => {
      handleResult( result );
    });
  }

  /**
   * Update the object by Id
   *
   * @param object
   * @param objectUrl
   * @param subject
   * @param dataArray
     */
  update(object:any, objectUrl:String, handleResult:Function, result:any) {
    this._httpClient.put(this.host + objectUrl + object.id, httpOptions).subscribe(data=> {
      handleResult( result );
    });
  }

  /**
   *
   * @param id
   * @param objectUrl
   * @param subject
   * @param dataArray
   */
  remove(id:string, objectUrl:String,handleResult:Function, result:any) {
    console.log("deleting data(" + id + ")");
    this._httpClient.delete(this.host + objectUrl + id, httpOptions).subscribe(data=> {
      handleResult( result );
    });
  }

  /**
   * Post the error to the alertService to display it over the current component
   *
   * @param err
   */
  private handleError( err:any): void {
    console.log("HttpService::handleError: " + err.error);
    let error = new AppError();
    error.id = 1;
    error.type = "Validation";
    error.message = err.error;
    this._alertService.push( error );
  }
}
