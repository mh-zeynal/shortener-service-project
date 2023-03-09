import { Component, OnInit } from '@angular/core';
import {ChildrenOutletContexts, Router, RouterOutlet} from "@angular/router";

@Component({
  selector: 'app-lobby',
  templateUrl: './lobby.component.html',
  styleUrls: ['./lobby.component.scss']
})
export class LobbyComponent implements OnInit {

  isUserDropdownOpen = false;

  constructor() { }

  ngOnInit(): void {
  }

  isUserLoggedIn(): boolean {
    let data = localStorage.getItem('data');
    return !data ? false : true;
  }

}
