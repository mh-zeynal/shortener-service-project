import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard/dashboard.component';
import {UserMenuRoutingModule} from "./user-account-routing.module";
import {OverlayModule} from "@angular/cdk/overlay";
import {A11yModule} from "@angular/cdk/a11y";
import {SharedModule} from "../shared/shared.module";
import { UrlGenerationFormComponent } from './url-generation-form/url-generation-form.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {ReactiveFormsModule} from "@angular/forms";
import {MatInputModule} from "@angular/material/input";
import {MatButtonModule} from "@angular/material/button";
import {ClipboardModule} from "@angular/cdk/clipboard";
import {RouterModule} from "@angular/router";
import {MatIconModule} from '@angular/material/icon';


@NgModule({
  declarations: [
    DashboardComponent,
    UrlGenerationFormComponent
  ],
    imports: [
        CommonModule,
        UserMenuRoutingModule,
        OverlayModule,
        A11yModule,
        SharedModule,
        MatFormFieldModule,
        MatInputModule,
        MatButtonModule,
        ReactiveFormsModule,
        ClipboardModule,
        RouterModule,
        MatIconModule,
    ]
})
export class UserAccountModule { }
