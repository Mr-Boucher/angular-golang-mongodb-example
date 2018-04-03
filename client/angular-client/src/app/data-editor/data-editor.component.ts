import {Component, OnInit, Input, ChangeDetectionStrategy} from '@angular/core';

import {Data, DataEditorService} from "./data-editor.service";
import {AlertService} from "../alert/alert.service";
import {Observable} from "rxjs/Observable";
import {DataSet} from "./data-editor.service";

@Component({
  selector: 'app-data-editor',
  templateUrl: './data-editor.component.html',
  styleUrls: ['./data-editor.component.css'],
  providers: [],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DataEditorComponent implements OnInit {

  @Input('data') data: DataSet;
  private asyncData: Observable<DataSet>;
  private error;
  private page: number = 1;
  private total: number;
  private loading: boolean;

  constructor(private _dataEditorService: DataEditorService) {
  }

  ngOnInit() {
    console.log( "ngOnInit" );
    //this._dataEditorService.data.subscribe(
    //  data => {
    //    this.data = data;
    //    console.log("DataEditorComponent::result");
    //  },
    //  err => {
    //    this.error = err;
    //    console.error(err);
    //  },
    //  () => {
    //    console.log('done loading');
    //  }
    //);

    this.getPage(null, 1);
  }

  getObservable() :Observable<DataSet> {
    return this.asyncData;
  }

  getPage(data, page: number) {
    this.loading = true;
    this.asyncData = this._dataEditorService.search( data, page );
  }

  search( searchCriteria:string, page:number, $event):Observable<DataSet> {
    return this._dataEditorService.search( searchCriteria, page );
  }

  add( data, $event ):void {
    this._dataEditorService.add( data )
  }

  remove( id: string, $event ):void {
    this._dataEditorService.remove( id );
  }
}
