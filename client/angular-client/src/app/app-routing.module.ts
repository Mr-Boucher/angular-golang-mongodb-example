import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {DataEditorComponent} from "./data-editor/data-editor.component";
import {ConfigurationComponent} from "./configuration/configuration.component";
import {TestComponent} from "./test/test.component";

const routes: Routes = [
  { path: '', redirectTo: '/data-editor', pathMatch: 'full' },
  { path: 'data-editor', component: DataEditorComponent },
  { path: 'configuration', component: ConfigurationComponent },
  { path: 'test', component: TestComponent}
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ],
})
export class AppRoutingModule { }
