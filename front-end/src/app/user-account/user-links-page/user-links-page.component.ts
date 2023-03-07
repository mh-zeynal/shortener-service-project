import {Component, OnInit, TemplateRef, ViewChild} from '@angular/core';
import {UrlType} from "../../shared/interfaces/url-type";
import {Clipboard} from "@angular/cdk/clipboard";
import {MatSnackBar} from "@angular/material/snack-bar";
import {MatDialog, MatDialogRef} from "@angular/material/dialog";
import {delay} from "rxjs/operators";
import {timer} from "rxjs";
import {HttpService} from "../../shared/services/http.service";

@Component({
  selector: 'app-user-links-page',
  templateUrl: './user-links-page.component.html',
  styleUrls: ['./user-links-page.component.scss']
})
export class UserLinksPageComponent implements OnInit {
  links: UrlType[] = [/*{userId: 1, title: 'css tricks', createdAt: new Date('2023-02-11T18:01:26'),
    originalUrl: 'https://css-tricks.com',
    shortUrl: 'Ty00R9sW', description: ''}*/];
  targetUrl !: UrlType;
  dialogRef !: MatDialogRef<any>;
  @ViewChild('editBox') editBox !: TemplateRef<any>;
  @ViewChild('deleteBox') deleteBox !: TemplateRef<any>;

  constructor(private clipboard: Clipboard,
              private snackBar: MatSnackBar,
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
    this.clipboard.copy('http://localhost:9090/' + link.shortUrl);
    this.snackBar.open('✔️' + 'url copied to clipboard', 'ok', {duration: 5000})
  }

  openUrlEditDialog(link: UrlType) {
    this.targetUrl = link;
    this.dialogRef = this.dialog.open(this.editBox);
  }

  openUrlDeleteDialog(link: UrlType) {
    this.targetUrl = link;
    this.dialogRef = this.dialog.open(this.deleteBox);
  }

  showMessageNotification(message: string) {
    this.snackBar.open('✔️' + message, 'ok', {duration: 5000});
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
      this.snackBar.open('✔️' + response?.body?.message, 'ok', {duration: 5000});
      this.dialogRef.close();
    })
  }

}
