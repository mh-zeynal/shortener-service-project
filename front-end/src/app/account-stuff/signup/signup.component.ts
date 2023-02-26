import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AbstractAccountType} from "../abstract-account-type";
import {FormControl, FormGroup} from "@angular/forms";

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent extends AbstractAccountType implements OnInit {
  form = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
    email: new FormControl(''),
    name: new FormControl(''),
  })

  constructor() {
    super();
  }

  ngOnInit(): void {
  }

  emitClientDataForm(): void {
    this.formSubmit.emit(this.form);
  }

}
