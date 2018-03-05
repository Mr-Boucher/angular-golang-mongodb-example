import { Component } from '@angular/core';
import { ConfigurationComponent } from 'configuration/Configuration.component'
import { DataEditorComponent } from 'data-editor-viewer/data-editor.component'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Angular Data Editor using (Go)lang and MongoDB';
}
