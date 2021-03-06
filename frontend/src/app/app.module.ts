import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MaterialModule } from './material.module';
import { SettingsComponent } from './settings/settings.component';
import { ProductsComponent } from './products/products.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';


@NgModule({
  declarations: [
    AppComponent,
    SettingsComponent,
    ProductsComponent,
    PageNotFoundComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
