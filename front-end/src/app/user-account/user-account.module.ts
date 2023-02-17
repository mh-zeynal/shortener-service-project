import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard/dashboard.component';
import {UserMenuRoutingModule} from "./user-account-routing.module";
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";
import {SharedModule} from "../shared/shared.module";



@NgModule({
  declarations: [
    DashboardComponent
  ],
    imports: [
      CommonModule,
      UserMenuRoutingModule,
      OverlayModule,
      A11yModule,
      SharedModule
    ]
})
export class UserAccountModule { }
