import { Component } from '@angular/core';
import {DataEditorComponent} from "./data-editor/data-editor.component";
import {ConfigurationComponent} from "./configuration/configuration.component";
import {AlertComponent} from "./alert/alert.component";
import {AlertService} from "./alert/alert.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Angular Data Editor using (Go)lang and MongoDB';

  constructor(private _alertService:AlertService) {
  }

  hasErrors(): boolean  {
    return this._alertService.hasErrors();
  }
}
