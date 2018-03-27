import {Injectable} from '@angular/core';
import {Subject} from "rxjs/Subject";
import {Observable} from "rxjs/Observable";

/**
 * Supports the CRUD of data objects
 */
@Injectable()
export class AlertService {

  subject:Subject<String> = new Subject();
  error: String = "";

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
    return this.error.length > 0;
  }

  push(error:String): void {
    this.error = error;
    this.subject.next( error );
  }
}

