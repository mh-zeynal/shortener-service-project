import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {Observable} from "rxjs";
import {UrlType} from "../interfaces/url-type";
import {ResponseMessage} from "../interfaces/response-message";
import {UserLinks} from "../interfaces/user-links";

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  public sendGetRequest(url: string): Observable<HttpResponse<UserLinks>>{
    return this.http.get<UserLinks>(url, {observe: 'response'});
  }

  public sendPostRequest(url: string, body: any): Observable<HttpResponse<ResponseMessage>> {
    return this.http.post<ResponseMessage>(url, body, {observe: 'response'});
  }
}
