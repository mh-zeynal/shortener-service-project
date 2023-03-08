import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {UserDropdownComponent} from "./components/user-dropdown/user-dropdown.component";
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";
import {HttpClientModule} from "@angular/common/http";
import {MatSnackBarModule} from "@angular/material/snack-bar";



@NgModule({
  declarations: [
    UserDropdownComponent
  ],
  imports: [
    CommonModule,
    OverlayModule,
    A11yModule,
    HttpClientModule,
    MatSnackBarModule
  ],
  exports: [
    UserDropdownComponent
  ]
})
export class SharedModule { }
