import { Component } from '@angular/core';
import {DataEditorComponent} from "./data-editor/data-editor.component";
import {ConfigurationComponent} from "./configuration/configuration.component";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Angular Data Editor using (Go)lang and MongoDB';
}
