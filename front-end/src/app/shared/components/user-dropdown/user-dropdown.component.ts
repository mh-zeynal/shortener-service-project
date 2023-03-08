import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {HttpService} from "../../services/http.service";
import {UserBriefData} from "../../interfaces/user-brief-data";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-user-dropdown',
  templateUrl: './user-dropdown.component.html',
  styleUrls: ['./user-dropdown.component.scss']
})
export class UserDropdownComponent implements OnInit {
  briefData: UserBriefData | null | undefined;

  constructor(private router: Router,
              private http: HttpService,
              private snackBar: MatSnackBar) { }

  ngOnInit(): void {
    let temp = localStorage.getItem('data');
    if (temp) {
      this.briefData = JSON.parse(temp);
      return;
    }
    this.http.getBriefData().subscribe(response => {
      if (!response.ok)
        return;
      this.briefData = response?.body;
      localStorage.setItem('data', JSON.stringify(this.briefData));
    })
  }

  isCurrentRouteOnUserPage(): boolean {
    let currentUrl = this.router.url
    let urlReg = new RegExp('/user')
    return urlReg.test(currentUrl)
  }

  goToUserDashboard() {
    this.router.navigateByUrl('/user/');
  }

  logoutUser() {
    this.http.logoutUser().subscribe(response => {
      debugger
      if (!response?.ok)
        return;
      localStorage.removeItem('data');
      this.snackBar.open('✔️' + response?.body?.message,
        'ok', {duration: 5000});
      this.router.navigateByUrl('/account');
    })
  }

}
