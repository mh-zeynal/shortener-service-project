import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {Observable} from "rxjs";
import {UrlType} from "../interfaces/url-type";
import {ResponseMessage} from "../interfaces/response-message";
import {UserLinks} from "../interfaces/user-links";
import {UserBriefData} from "../interfaces/user-brief-data";

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  public logoutUser(): Observable<HttpResponse<ResponseMessage>> {
    return this.http.get<ResponseMessage>('/api/logout', {observe: 'response'});
  }

  public getBriefData(): Observable<HttpResponse<UserBriefData>> {
    return this.http.get<UserBriefData>('/api/briefData', {observe: 'response'});
  }

  public sendGetRequest(url: string): Observable<HttpResponse<UserLinks>> {
    return this.http.get<UserLinks>(url, {observe: 'response'});
  }

  public sendPostRequest(url: string, body: any): Observable<HttpResponse<ResponseMessage>> {
    return this.http.post<ResponseMessage>(url, body, {observe: 'response'});
  }
}
