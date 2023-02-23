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
  @ViewChild('short_url') urlInput !: ElementRef<HTMLInputElement>;
  form = new FormGroup({
    title: new FormControl('', [Validators.required]),
    originalUrl: new FormControl('', [validateUrlFormat, Validators.required]),
    description: new FormControl('')
  })

  constructor(private http: HttpService, private clipboard: Clipboard) { }

  ngOnInit(): void {
  }

  submitForm(){
    this.http.sendPostRequest('/api/shorten', this.convertToFormData()).subscribe( value => {
      console.log(value)
    })
  }

  copyUrlToClipbaord(){
    if (!this.urlInput)
      return;
    this.clipboard.copy(this.urlInput.nativeElement.value);
  }

  private convertToFormData(): FormData {
    let formData = new FormData();
    formData.append('title', this.form.get('title')?.value);
    formData.append('url', this.form.get('originalUrl')?.value);
    formData.append('description', this.form.get('description')?.value);
    return formData;
  }

}
