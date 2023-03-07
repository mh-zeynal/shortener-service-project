import {Component, ElementRef, EventEmitter, Input, OnInit, Output, ViewChild} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {HttpService} from "../../shared/services/http.service";
import {Clipboard} from "@angular/cdk/clipboard";
import {UrlType} from "../../shared/interfaces/url-type";

@Component({
  selector: 'app-url-generation-form',
  templateUrl: './url-generation-form.component.html',
  styleUrls: ['./url-generation-form.component.scss']
})
export class UrlGenerationFormComponent implements OnInit {

  @Input() isEditing = false;
  @Input() inputUrlObject !: UrlType;
  @Output() shortUrl = new EventEmitter<string>();
  @Output() responseMessage = new EventEmitter<string>();

  form = new FormGroup({
    title: new FormControl('', [Validators.required]),
    originalUrl: new FormControl('', [
      Validators.pattern('(https?://)?([\\da-z.-]+)\\.([a-z.]{2,6})[/\\w .-]*/?'), Validators.required
    ]),
    description: new FormControl('', [Validators.maxLength(200)])
  })

  constructor(private http: HttpService) { }

  ngOnInit(): void {
    if (this.inputUrlObject) {
      this.form.patchValue({
        title: this.inputUrlObject?.title,
        originalUrl: this.inputUrlObject?.originalUrl,
        description: this.inputUrlObject?.description.toString()
      })
    }
  }

  submitForm(){
    debugger
    if (this.form.invalid)
      return;
    let path = this.defineRequestUrl();
    let requestBody = this.defineRequestBody();
    this.http.sendPostRequest(path, requestBody)
      .subscribe( response => {
        debugger
      if (!response.ok || response?.body?.isError)
        return;
      if (!this.isEditing)
        this.shortUrl.emit(response?.body?.link);
      else
        this.responseMessage.emit(response?.body?.message);

    })
  }

  showUrlFieldError() {
    let urlControl = this.form.controls.originalUrl;
    if (urlControl.hasError('required'))
      return 'please enter your intended url';
    else if (urlControl.hasError('pattern'))
      return 'please enter a valid url';
    return ''
  }

  defineRequestUrl() {
    return this.isEditing ? '/api/editUrl' : '/api/shorten';
  }

  defineRequestBody() {
    if (!this.isEditing)
      return this.form.value;
    this.inputUrlObject.title = this.form.controls['title'].value;
    this.inputUrlObject.originalUrl = this.form.controls['originalUrl'].value;
    this.inputUrlObject.description = this.form.controls['description'].value;
    return this.inputUrlObject;
  }

}
