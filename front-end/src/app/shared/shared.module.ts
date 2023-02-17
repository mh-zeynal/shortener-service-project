import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {UserDropdownComponent} from "./user-dropdown/user-dropdown.component";
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";



@NgModule({
  declarations: [
    UserDropdownComponent
  ],
  imports: [
    CommonModule,
    OverlayModule,
    A11yModule
  ],
  exports: [
    UserDropdownComponent
  ]
})
export class SharedModule { }
