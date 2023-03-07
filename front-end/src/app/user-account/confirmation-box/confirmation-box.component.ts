import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-confirmation-box',
  templateUrl: './confirmation-box.component.html',
  styleUrls: ['./confirmation-box.component.scss']
})
export class ConfirmationBoxComponent implements OnInit {

  @Input() parentComponent!: any;

  constructor() { }

  ngOnInit(): void {
  }

}
