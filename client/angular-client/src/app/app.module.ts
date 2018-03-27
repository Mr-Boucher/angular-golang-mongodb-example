import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {HttpClientModule} from '@angular/common/http';
import { FormsModule } from '@angular/forms';


import { AppComponent } from './app.component';
import { DataEditorComponent } from './data-editor/data-editor.component';
import { ConfigurationComponent } from './configuration/configuration.component';
import { AppRoutingModule } from './app-routing.module';
import {HttpService} from "./http.service";
import {AlertComponent} from "./alert/alert.component";
import {AlertService} from "./alert/alert.service";


@NgModule({
  declarations: [
    AppComponent,
    DataEditorComponent,
    ConfigurationComponent,
    AlertComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule
  ],
  providers: [
    HttpService,
    AlertService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
