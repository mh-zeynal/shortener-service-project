import {AbstractControl, ValidationErrors, ValidatorFn} from "@angular/forms";

export function validateUrlFormat(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const urlRegex = /^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;
    const value = control.value;
    if (!value)
      return null;
    return !new RegExp(urlRegex).test(value) ? {invalidUrl: true} : null
  }
}
