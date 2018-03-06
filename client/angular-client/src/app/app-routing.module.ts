import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {DataEditorComponent} from "./data-editor-viewer/data-editor.component";
import {ConfigurationComponent} from "./configuration/configuration.component";

const routes: Routes = [
  { path: '', redirectTo: '/data-editor', pathMatch: 'full' },
  { path: 'data-editor', component: DataEditorComponent },
  { path: 'configuration', component: ConfigurationComponent }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ],
})
export class AppRoutingModule { }
