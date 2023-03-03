import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatInputModule} from "@angular/material/input";
import {MatFormFieldModule} from "@angular/material/form-field";
import {SharedModule} from "../shared/shared.module";
import {AccountStuffRoutingModule} from "./account-stuff-routing.module";
import { SignupComponent } from './signup/signup.component';
import {MatButtonModule} from "@angular/material/button";
import { LoginComponent } from './login/login.component';
import {ReactiveFormsModule} from "@angular/forms";
import {MatSnackBarModule} from '@angular/material/snack-bar';



@NgModule({
  declarations: [
    SignupComponent,
    LoginComponent
  ],
  exports: [
    SignupComponent,
    LoginComponent
  ],
  imports: [
    CommonModule,
    MatInputModule,
    MatFormFieldModule,
    SharedModule,
    MatButtonModule,
    ReactiveFormsModule,
    MatSnackBarModule,
    AccountStuffRoutingModule
  ]
})
export class AccountStuffModule { }
