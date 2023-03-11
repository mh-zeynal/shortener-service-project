import { Injectable } from '@angular/core';
import {MatSnackBar} from "@angular/material/snack-bar";
import {HttpResponse} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  constructor(private snackBar: MatSnackBar) { }

  public triggerNotification(message: string, state: string, action = 'ok', duration = 5000):void {
    let temp = this.formatMessageByItsState(message, state);
    this.snackBar.open(temp, action, {duration: duration});
  }

  public triggerNotificationOnResponse(response: HttpResponse<any>, action = 'ok', duration = 5000):void {
    let state = this.defineStateByResponseStatus(response);
    let message = response?.body?.message;
    let temp = this.formatMessageByItsState(message, state);
    this.snackBar.open(temp, action, {duration: duration});
  }

  private formatMessageByItsState(message: string, state: string) {
    if (state === 'success')
      return '✔️ ' + message;
    else if (state === 'failure')
      return '❌ ' + message;
    return message;
  }

  private defineStateByResponseStatus(response: HttpResponse<any>) {
    return response?.body?.isError ? 'failure' : 'success';
  }

}
