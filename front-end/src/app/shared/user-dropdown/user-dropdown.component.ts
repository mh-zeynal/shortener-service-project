import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";

@Component({
  selector: 'app-user-dropdown',
  templateUrl: './user-dropdown.component.html',
  styleUrls: ['./user-dropdown.component.scss']
})
export class UserDropdownComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  isCurrentRouteOnUserPage(): boolean{
    let currentUrl = this.router.url
    let urlReg = new RegExp('/user')
    return urlReg.test(currentUrl)
  }

}
