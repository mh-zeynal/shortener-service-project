import { Component, OnInit } from '@angular/core';
import {HttpService} from "../../shared/services/http.service";
import {ResponseMessage} from "../../shared/interfaces/response-message";
import {FormGroup} from "@angular/forms";

@Component({
  selector: 'app-account-page',
  templateUrl: './account-page.component.html',
  styleUrls: ['./account-page.component.scss']
})
export class AccountPageComponent implements OnInit {

  hasAccount = false;

  constructor(private http: HttpService) { }

  ngOnInit(): void {
  }

  submitForm(url: string, event: FormGroup){
    debugger
    this.http.sendPostRequest(url, event.value).subscribe(value => {
      console.log(value);
    })
  }

}
