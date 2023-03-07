import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {Clipboard} from "@angular/cdk/clipboard";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-url-generator-page',
  templateUrl: './url-generator-page.component.html',
  styleUrls: ['./url-generator-page.component.scss']
})
export class UrlGeneratorPageComponent implements OnInit {

  @ViewChild('short_url') urlInput !: ElementRef<HTMLInputElement>;
  shortenedUrlValue !: string;

  constructor(private clipboard: Clipboard, private snackBar: MatSnackBar) { }

  ngOnInit(): void {
  }

  copyUrlToClipboard(){
    if (!this.shortenedUrlValue)
      return;
    this.clipboard.copy(this.shortenedUrlValue);
    this.snackBar.open('✔️' + 'url copied to clipboard', 'ok', {duration: 5000});
  }

  assignEmittedShortUrl(event: string) {
    this.shortenedUrlValue = event;
    this.urlInput.nativeElement.value = this.shortenedUrlValue;
  }

}
