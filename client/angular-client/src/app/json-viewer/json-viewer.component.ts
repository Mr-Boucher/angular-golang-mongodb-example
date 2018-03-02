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
  data: Data[];
  error: any;

  constructor(private _jsonViewerService: JsonViewerService) {
  }

  ngOnInit() {
    console.log( "ngOnInit" );
    this._jsonViewerService.data.subscribe(
      data => {
        this.data = data;
        console.log("subscribe result")
      },
      err => console.error(err),
      () => console.log('done loading courses')
    );
  }

  refresh($event) {
    this._jsonViewerService.load();
  }

  add( data, $event ) {
    this._jsonViewerService.add( data )
  }

  remove( id: string, $event ) {
    this._jsonViewerService.remove( id );
  }
}
