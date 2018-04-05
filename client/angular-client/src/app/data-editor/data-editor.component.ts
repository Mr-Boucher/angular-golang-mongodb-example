import {Component, OnInit, Input, ChangeDetectionStrategy} from '@angular/core';

import {Data, DataEditorService} from "./data-editor.service";
import {AlertService} from "../alert/alert.service";
import {Observable} from "rxjs/Observable";
import {Page} from "./data-editor.service";

@Component({
  selector: 'app-data-editor',
  templateUrl: './data-editor.component.html',
  styleUrls: ['./data-editor.component.css'],
  providers: [],
  changeDetection: ChangeDetectionStrategy.Default
})
export class DataEditorComponent implements OnInit {

  data:Page = new Page();
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

    this.getPage(this.data);
  }

  getDataList():Data[] {
    return this.data.data;
  }

  pageChanged(event):number {
    console.log('DataEditorComponent::pageChanged to ' + event);
    this.data.data = [];
    this.data.pageNumber = event;
    this.getPage(this.data);
    return event;
  }

  getPage(page:Page):void {
    this.loading = true;
    this._dataEditorService.search(this.data);
  }

  search(searchCriteria:string, page:number, event):Observable<Page> {
    this.data.filter = searchCriteria;
    this.data.pageNumber = page;
    return this._dataEditorService.search(this.data);
  }

  add(data, $event):void {
    this._dataEditorService.add(data)
  }

  remove(id:string, $event):void {
    this._dataEditorService.remove(id);
  }
}
