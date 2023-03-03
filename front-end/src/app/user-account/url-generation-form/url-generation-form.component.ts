import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {validateUrlFormat} from "../../shared/validators/url-validator";
import {HttpService} from "../../shared/services/http.service";
import {Clipboard} from "@angular/cdk/clipboard";

@Component({
  selector: 'app-url-generation-form',
  templateUrl: './url-generation-form.component.html',
  styleUrls: ['./url-generation-form.component.scss']
})
export class UrlGenerationFormComponent implements OnInit {
  isResponseOk = false;

  @ViewChild('short_url') urlInput !: ElementRef<HTMLInputElement>;

  form = new FormGroup({
    title: new FormControl('', [Validators.required]),
    originalUrl: new FormControl('', [
      Validators.pattern('(https?://)?([\\da-z.-]+)\\.([a-z.]{2,6})[/\\w .-]*/?'), Validators.required
    ]),
    description: new FormControl('', [Validators.maxLength(200)])
  })

  constructor(private http: HttpService, private clipboard: Clipboard) { }

  ngOnInit(): void {
  }

  submitForm(){
    console.log(this.form)
    if (this.form.invalid)
      return;
    this.http.sendPostRequest('/api/shorten', this.form.value).subscribe( response => {
      debugger
      if (response.ok) {
        this.isResponseOk = true;
        this.urlInput.nativeElement.value = response?.body?.link ?? '';
      }
    })
  }

  copyUrlToClipboard(){
    if (!this.urlInput)
      return;
    this.clipboard.copy(this.urlInput.nativeElement.value);
  }

  showUrlFieldError() {
    let urlControl = this.form.controls.originalUrl;
    if (urlControl.hasError('required'))
      return 'please enter your intended url';
    else if (urlControl.hasError('pattern'))
      return 'please enter a valid url';
    return ''
  }

}
