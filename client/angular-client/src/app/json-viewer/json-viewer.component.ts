import {Component, OnInit} from '@angular/core';

import {Data, JsonViewerService} from "./json-viewer.service";

@Component({
  selector: 'app-json-viewer',
  templateUrl: './json-viewer.component.html',
  styleUrls: ['./json-viewer.component.css'],
  providers: [JsonViewerService]
})
export class JsonViewerComponent implements OnInit {

  json: any;
  courses: Data[];
  error: any;

  constructor(private _jsonViewerService: JsonViewerService) {
  }

  ngOnInit() {
    console.log( "ngOnInit" );
    this._jsonViewerService.data.subscribe(
      data => {
        this.courses = data;
        console.log("subscribe result")
      },
      err => console.error(err),
      () => console.log('done loading courses')
    );
  }

  refresh($event) {
    this._jsonViewerService.updateData();
  }

  remove( id: string, $event ) {
    this._jsonViewerService.removeData( id );
  }
}
