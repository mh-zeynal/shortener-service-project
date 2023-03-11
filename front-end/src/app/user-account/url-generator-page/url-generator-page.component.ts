import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {Clipboard} from "@angular/cdk/clipboard";
import {NotificationService} from "../../shared/services/notification.service";

@Component({
  selector: 'app-url-generator-page',
  templateUrl: './url-generator-page.component.html',
  styleUrls: ['./url-generator-page.component.scss']
})
export class UrlGeneratorPageComponent implements OnInit {

  @ViewChild('short_url') urlInput !: ElementRef<HTMLInputElement>;
  shortenedUrlValue !: string;

  constructor(private clipboard: Clipboard, private notification: NotificationService) { }

  ngOnInit(): void {
  }

  copyUrlToClipboard(){
    if (!this.shortenedUrlValue)
      return;
    this.clipboard.copy(this.shortenedUrlValue);
    this.notification.triggerNotification('url copied to clipboard', 'success');
  }

  assignEmittedShortUrl(event: string) {
    this.shortenedUrlValue = event;
    this.urlInput.nativeElement.value = this.shortenedUrlValue;
  }

}
