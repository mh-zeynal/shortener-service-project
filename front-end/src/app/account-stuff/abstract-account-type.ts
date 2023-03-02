import {Directive, EventEmitter, Output} from "@angular/core";
import {FormGroup} from "@angular/forms";

@Directive()
export abstract class AbstractAccountType {
  form !: FormGroup;
  @Output() DoesClientHasAccount = new EventEmitter<boolean>();
  @Output() formSubmit = new EventEmitter<FormGroup>();

  emitClientDataForm(): void {
    if (this.form.invalid)
      return;
    this.formSubmit.emit(this.form);
  }

}
