import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {LobbyComponent} from "./lobby/lobby.component";
import {HomePageComponent} from "./home-page/home-page.component";
import {AboutPageComponent} from "./about-page/about-page.component";
import {ContactPageComponent} from "./contact-page/contact-page.component";

const routes: Routes = [
  {path: 'lobby', component: LobbyComponent,
    children: [
      {path: 'home', component: HomePageComponent, data: {animation: 'homePage'}},
      {path: 'about', component: AboutPageComponent, data: {animation: 'aboutPage'}},
      {path: 'contact-us', component: ContactPageComponent, data: {animation: 'contactPage'}}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class WebsiteLobbyRoutingModule { }
