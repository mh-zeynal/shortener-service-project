import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {DashboardComponent} from "./dashboard/dashboard.component";
import {UrlGenerationFormComponent} from "./url-generation-form/url-generation-form.component";



const routes: Routes = [
  {path: 'user', component: DashboardComponent,
    children: [
      {path: 'generate-url', component: UrlGenerationFormComponent},
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserMenuRoutingModule { }
