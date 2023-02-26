import {Directive, EventEmitter, Output} from "@angular/core";
import {FormGroup} from "@angular/forms";

@Directive()
export abstract class AbstractAccountType {
  @Output() DoesClientHasAccount = new EventEmitter<boolean>();
  @Output() formSubmit = new EventEmitter<FormGroup>();

  abstract emitClientDataForm(): void;

}
