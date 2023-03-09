import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {RouterModule} from "@angular/router";

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";
import {SharedModule} from "./shared/shared.module";
import {UserAccountModule} from "./user-account/user-account.module";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {AccountStuffModule} from "./account-stuff/account-stuff.module";
import {WebsiteLobbyModule} from "./website-lobby/website-lobby.module";

@NgModule({
  declarations: [
    AppComponent,
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
    AccountStuffModule,
    WebsiteLobbyModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
