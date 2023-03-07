import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {DashboardComponent} from "./dashboard/dashboard.component";
import {UserLinksPageComponent} from "./user-links-page/user-links-page.component";
import {UrlGeneratorPageComponent} from "./url-generator-page/url-generator-page.component";



const routes: Routes = [
  {path: 'user', component: DashboardComponent,
    children: [
      {path: 'generate-url', component: UrlGeneratorPageComponent},
      {path: 'links', component: UserLinksPageComponent}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserMenuRoutingModule { }
