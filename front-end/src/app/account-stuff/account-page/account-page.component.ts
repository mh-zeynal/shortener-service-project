import { Component, OnInit } from '@angular/core';
import {HttpService} from "../../shared/services/http.service";
import {FormGroup} from "@angular/forms";
import {Router} from "@angular/router";
import {NotificationService} from "../../shared/services/notification.service";

@Component({
  selector: 'app-account-page',
  templateUrl: './account-page.component.html',
  styleUrls: ['./account-page.component.scss']
})
export class AccountPageComponent implements OnInit {

  hasAccount = false;

  constructor(private http: HttpService,
              private notification: NotificationService, private router: Router) { }

  ngOnInit(): void {
  }

  submitForm(url: string, event: FormGroup){
    this.http.sendPostRequest(url, event.value).subscribe(response => {
      let msgFlag = response?.body?.isError;
      let message = response?.body?.message;
      this.notification.triggerNotificationOnResponse(response);
      if (!msgFlag)
        this.router.navigateByUrl('/user')
    })
  }

}
