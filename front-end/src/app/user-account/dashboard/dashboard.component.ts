import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  @ViewChild('buttonContainer') buttonContainer !: ElementRef;

  isUserDropdownOpen = false;

  constructor() { }

  ngOnInit(): void {
  }

  markAsSelected(targetButton: Element){
    let button = document.getElementsByClassName('active-dashboard-button');
    while (button.length)
      button[0].classList.remove('active-dashboard-button');
    targetButton.classList.add('active-dashboard-button');
  }

}
