import {AbstractControl, AsyncValidator, AsyncValidatorFn, ValidationErrors} from "@angular/forms";
import {Observable, of} from "rxjs";
import {delay, map} from "rxjs/operators";
import {Injectable} from "@angular/core";

Injectable({providedIn: 'root'})
export class PasswordValidator implements AsyncValidator{
  validate(control: AbstractControl): Observable<ValidationErrors | null> {
    let password = control.value;
    return this.checkIfPasswordHasValidCharacters(password).pipe(
      map(res => {
        return res ? {'passwordInvalid': true} : null;
      })
    )
  }

  private checkIfPasswordHasValidCharacters(password: string): Observable<boolean> {
    let passwordRegex = new RegExp('^[a-zA-Z0-9@]');
    return of(passwordRegex.test(password)).pipe(delay(1))
  }

}
