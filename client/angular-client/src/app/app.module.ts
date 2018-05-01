import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {HttpClientModule} from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { NgxPaginationModule } from 'ngx-pagination';


import { AppComponent } from './app.component';
import { DataEditorComponent } from './data-editor/data-editor.component';
import { ConfigurationComponent } from './configuration/configuration.component';
import { AppRoutingModule } from './app-routing.module';
import {HttpService} from "./http.service";
import {AlertComponent} from "./alert/alert.component";
import {AlertService} from "./alert/alert.service";
import {DataEditorService} from "./data-editor/data-editor.service";
import {ConfigurationService} from "./configuration/configuration.service";
import {TestService} from "./test/test.service";
import {TestComponent} from "./test/test.component";


@NgModule({
  declarations: [
    AppComponent,
    DataEditorComponent,
    ConfigurationComponent,
    TestComponent,
    AlertComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule,
    NgxPaginationModule
  ],
  providers: [
    HttpService,
    AlertService,
    DataEditorService,
    ConfigurationService,
    TestService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
