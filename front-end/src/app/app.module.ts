import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {RouterModule} from "@angular/router";

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { MainPageComponent } from './main-page/main-page.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";
import {SharedModule} from "./shared/shared.module";
import {UserAccountModule} from "./user-account/user-account.module";
import { AccountPageComponent } from './account-stuff/account-page/account-page.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {AccountStuffModule} from "./account-stuff/account-stuff.module";

@NgModule({
  declarations: [
    AppComponent,
    MainPageComponent,
    AccountPageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    RouterModule,
    BrowserAnimationsModule,
    OverlayModule,
    A11yModule,
    SharedModule,
    UserAccountModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    AccountStuffModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
