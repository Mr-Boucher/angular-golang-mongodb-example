import {Component, OnInit} from '@angular/core';

import {Data, DataEditorService} from "./data-editor.service";

@Component({
  selector: 'app-data-editor',
  templateUrl: './data-editor.component.html',
  styleUrls: ['./data-editor.component.css'],
  providers: [DataEditorService]
})
export class DataEditorComponent implements OnInit {

  data: Data[];
  private error;

  constructor(private _dataEditorService: DataEditorService) {
  }

  ngOnInit() {
    console.log( "ngOnInit" );
    this._dataEditorService.data.subscribe(
      data => {
        this.data = data;
        console.log("subscribe result");
      },
      err => {
        this.error = err;
        console.error(err);
      },
      () => {
        console.log('done loading');
      }

    );
  }

  load( ): void {
    this._dataEditorService.load();
  }

  search( data, $event):void {
    this._dataEditorService.search( data );
  }

  add( data, $event ):void {
    this._dataEditorService.add( data )
  }

  remove( id: string, $event ):void {
    this._dataEditorService.remove( id );
  }
}
