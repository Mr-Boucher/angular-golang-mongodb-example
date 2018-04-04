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
  changeDetection: ChangeDetectionStrategy.Default
})
export class DataEditorComponent implements OnInit {

  data:DataSet;
  pageSize:number = 10;
  private error;
  private loading:boolean;

  constructor(private _dataEditorService:DataEditorService) {
  }

  ngOnInit() {
    console.log("DataEditorComponent::ngOnInit");
    this._dataEditorService.data.subscribe(
      data => {
        this.data = data;
        console.log("DataEditorComponent::result" + data.data);
      },
      err => {
        this.error = err;
        console.error("DataEditorComponent::error " + err);
      },
      () => {
        console.log('DataEditorComponent::done loading');
      }
    );

    this.getPage(null, 1);
  }

  getDataList():Data[] {
    return this.data.data;
  }

  pageChanged(event):number {
    console.log('DataEditorComponent::pageChanged to ' + event);
    this.data.data = [];
    this.getPage("", event);
    return event;
  }

  getPage(searchCriteria:string, page:number):void {
    this.loading = true;
    this._dataEditorService.search(searchCriteria, page, this.pageSize);
  }

  search(searchCriteria:string, page:number, event):Observable<DataSet> {
    return this._dataEditorService.search(searchCriteria, page, this.pageSize);
  }

  add(data, $event):void {
    this._dataEditorService.add(data)
  }

  remove(id:string, $event):void {
    this._dataEditorService.remove(id);
  }
}
