import {Injectable} from '@angular/core';
import {Subject} from "rxjs/Subject";
import {Observable} from "rxjs/Observable";

/**
 * Supports the CRUD of data objects
 */
@Injectable()
export class AlertService {

  subject:Subject<String> = new Subject();
  _errors: String = "";

  /**
   *
   */
  constructor() {
  }

  /**
   * getter that converts the data into an observable
   *
   * @returns {Observable<Data[]>}
   */
  get errors(): Observable {
    return this.subject.asObservable();
  }

  hasErrors(): boolean {
    console.log( "AlertService::hasErrors: " + this._errors );
    return this._errors.length > 0;
  }

  clearErrors():void {
    console.log( "AlertService::clearErrors: " + this._errors );
    this._errors = "";
    console.log( "AlertService::clearErrors: " + this._errors );
    this.subject.next( this._errors );
  }

  push(daError:string): void {
    console.log( "AlertService::push: " + this._errors );
    this._errors = daError;
    console.log( "AlertService::push: " + this._errors );
    this.subject.next( this._errors );
  }
}

