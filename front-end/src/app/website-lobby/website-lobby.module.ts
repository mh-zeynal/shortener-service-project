import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {WebsiteLobbyRoutingModule} from "./website-lobby-routing.module";
import { LobbyComponent } from './lobby/lobby.component';
import {MatButtonModule} from "@angular/material/button";
import {RouterModule} from "@angular/router";
import { HomePageComponent } from './home-page/home-page.component';
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {BrowserModule} from "@angular/platform-browser";
import { AboutPageComponent } from './about-page/about-page.component';
import { ContactPageComponent } from './contact-page/contact-page.component';
import {SharedModule} from "../shared/shared.module";
import {A11yModule} from "@angular/cdk/a11y";
import {OverlayModule} from "@angular/cdk/overlay";



@NgModule({
  declarations: [
    LobbyComponent,
    HomePageComponent,
    AboutPageComponent,
    ContactPageComponent
  ],
  imports: [
    CommonModule,
    WebsiteLobbyRoutingModule,
    MatButtonModule,
    RouterModule,
    BrowserModule,
    BrowserAnimationsModule,
    SharedModule,
    A11yModule,
    OverlayModule,
  ]
})
export class WebsiteLobbyModule { }
