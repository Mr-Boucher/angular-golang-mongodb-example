import {Component, OnInit} from '@angular/core';

import {Data, DataEditorService} from "./data-editor.service";

@Component({
  selector: 'app-data-editor',
  templateUrl: './data-editor.component.html',
  styleUrls: ['./data-editor.component.css'],
  providers: [DataEditorService]
})
export class DataEditorComponent implements OnInit {

  json: any;
  data: Data[];
  error: any;

  constructor(private _dataEditorService: DataEditorService) {
  }

  ngOnInit() {
    console.log( "ngOnInit" );
    this._dataEditorService.data.subscribe(
      data => {
        this.data = data;
        console.log("subscribe result")
      },
      err => console.error(err),
      () => console.log('done loading courses')
    );
  }

  refresh($event) {
    this._dataEditorService.load();
  }

  add( data, $event ) {
    this._dataEditorService.add( data )
  }

  remove( id: string, $event ) {
    this._dataEditorService.remove( id );
  }
}