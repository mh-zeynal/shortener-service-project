import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard/dashboard.component';
import {UserMenuRoutingModule} from "./user-account-routing.module";



@NgModule({
  declarations: [
    DashboardComponent
  ],
  imports: [
    CommonModule,
    UserMenuRoutingModule
  ]
})
export class UserAccountModule { }
