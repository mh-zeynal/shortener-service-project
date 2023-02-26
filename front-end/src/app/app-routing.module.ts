import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {MainPageComponent} from "./main-page/main-page.component";
import {AccountPageComponent} from "./account-stuff/account-page/account-page.component";

const routes: Routes = [
  {path: "main", component: MainPageComponent},
  {path: "account", component: AccountPageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
