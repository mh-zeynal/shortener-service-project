import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AbstractAccountType} from "../abstract-account-type";
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent extends AbstractAccountType implements OnInit {

  constructor(private fb: FormBuilder) {
    super();
    this.form = new FormGroup({
      username: new FormControl('', Validators.required),
      password: new FormControl('', Validators.required),
    })
  }

  ngOnInit(): void {
  }

}
