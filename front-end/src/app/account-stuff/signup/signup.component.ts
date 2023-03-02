import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {AbstractAccountType} from "../abstract-account-type";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {PasswordValidator} from "../../shared/validators/password-validator";

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent extends AbstractAccountType implements OnInit {

  constructor(private passwordValidator: PasswordValidator) {
    super();
    this.form = new FormGroup({
      username: new FormControl('', Validators.required),
      password: new FormControl('', [Validators.required,
        this.passwordValidator.validate.bind(this.passwordValidator)]),
      email: new FormControl('', [Validators.required, Validators.email]),
      name: new FormControl('', Validators.required),
    })
  }

  ngOnInit(): void {
  }

  showEmailError(){
    let control = this.form.controls['email'];
    if (control.hasError('required'))
      return 'please enter your email';
    else if (control.hasError('email'))
      return 'the entered value doesn\'t seem like an email';
    return '';
  }

  showPasswordError(){
    let control = this.form.controls['password'];
    if (control.hasError('required'))
      return 'please enter your password';
    return 'use alphabets, numbers and @ in your password';
  }

}
