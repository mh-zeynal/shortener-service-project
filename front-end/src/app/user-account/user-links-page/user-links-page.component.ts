import {Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {UrlType} from "../../shared/interfaces/url-type";
import {Clipboard} from "@angular/cdk/clipboard";
import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {HttpService} from "../../shared/services/http.service";
import {NotificationService} from "../../shared/services/notification.service";
import {HttpResponse} from "@angular/common/http";
import {ResponseMessage} from "../../shared/interfaces/response-message";

@Component({
  selector: 'app-user-links-page',
  templateUrl: './user-links-page.component.html',
  styleUrls: ['./user-links-page.component.scss']
})
export class UserLinksPageComponent implements OnInit {
  links: UrlType[] = [];
  targetUrl !: UrlType;
  dialogRef !: MatDialogRef<any>;
  @ViewChild('editBox') editBox !: TemplateRef<any>;
  @ViewChild('deleteBox') deleteBox !: TemplateRef<any>;

  constructor(private clipboard: Clipboard,
              private notification: NotificationService,
              private dialog: MatDialog, private http: HttpService) { }

  ngOnInit(): void {
    this.http.sendGetRequest('/api/user_urls').subscribe(response => {
      if (!response?.ok || response?.body?.isError)
        return;
      let temp = response?.body?.links;
      if (temp){
        temp.forEach((link) => {
          this.links.push(link);
        })
      }
    })
  }

  copyUrlToClipboard(link: UrlType){
    this.clipboard.copy('http://localhost:9090/api/' + link.shortUrl);
    this.notification.triggerNotification('url copied to clipboard', 'success');
  }

  openUrlEditDialog(link: UrlType) {
    this.targetUrl = link;
    this.dialogRef = this.dialog.open(this.editBox);
  }

  openUrlDeleteDialog(link: UrlType) {
    this.targetUrl = link;
    this.dialogRef = this.dialog.open(this.deleteBox);
  }

  showMessageNotification(response: HttpResponse<ResponseMessage>) {
    this.notification.triggerNotificationOnResponse(response);
    this.dialogRef.close();
  }

  isLinkExpired(link: UrlType) {
    let now = new Date();
    let creationDate = new Date(link.createdAt);
    let dif = (now.getTime() - creationDate.getTime()) / (1000 * 3600 * 24);
    if (dif > 30)
      return true;
    return false;
  }

  sendLinkDeletionRequest() {
    this.http.sendPostRequest('/api/deleteUrl', this.targetUrl).subscribe(response => {
      this.links.forEach((link, index) => {
        if (link === this.targetUrl)
          this.links.splice(index, 1);
      })
      this.notification.triggerNotificationOnResponse(response);
      this.dialogRef.close();
    })
  }

}
