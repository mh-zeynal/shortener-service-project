import { Component, OnInit } from '@angular/core';
import {HttpService} from "../../shared/services/http.service";
import {FormGroup} from "@angular/forms";
import {MatSnackBar} from "@angular/material/snack-bar";
import {Router} from "@angular/router";

@Component({
  selector: 'app-account-page',
  templateUrl: './account-page.component.html',
  styleUrls: ['./account-page.component.scss']
})
export class AccountPageComponent implements OnInit {

  hasAccount = false;

  constructor(private http: HttpService,
              private snackBar: MatSnackBar, private router: Router) { }

  ngOnInit(): void {
  }

  submitForm(url: string, event: FormGroup){
    this.http.sendPostRequest(url, event.value).subscribe(response => {
      let msgFlag = response?.body?.isError;
      let message = response?.body?.message;
      if (msgFlag == false)
        message = '✔️' + message;
      else
        message = '❌' + message;
      this.snackBar.open(message, 'ok', {duration: 5000})
      if (!msgFlag)
        this.router.navigateByUrl('/user')
    })
  }

}
