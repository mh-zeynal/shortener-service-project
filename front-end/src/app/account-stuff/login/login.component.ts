import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AbstractAccountType} from "../abstract-account-type";
import {FormBuilder, FormControl, FormGroup} from "@angular/forms";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent extends AbstractAccountType implements OnInit {

  form = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
  })

  constructor(private fb: FormBuilder) {
    super();
  }

  ngOnInit(): void {
  }

  emitClientDataForm(): void {
    debugger
    this.formSubmit.emit(this.form);
  }

}
