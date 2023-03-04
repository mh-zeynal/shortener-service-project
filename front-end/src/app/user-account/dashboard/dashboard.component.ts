import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {Router} from "@angular/router";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  @ViewChild('buttonContainer') buttonContainer !: ElementRef;

  isUserDropdownOpen = false;

  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  markAsSelected(){
    let button = document.getElementsByClassName('active-dashboard-button');
    while (button.length)
      button[0].classList.remove('active-dashboard-button');
    let userRoute = this.router.url.replace('/user/', '');
    let activeButton = document.getElementById(userRoute);
    if (activeButton)
      activeButton.classList.add('active-dashboard-button');
  }

}
