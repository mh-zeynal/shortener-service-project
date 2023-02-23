import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {UrlType} from "../interfaces/url-type";
import {ResponseMessage} from "../interfaces/response-message";

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  public sendGetRequest(url: string): Observable<UrlType>{
    return this.http.get<UrlType>(url);
  }

  public sendPostRequest(url: string, body: any): Observable<ResponseMessage> {
    return this.http.post<ResponseMessage>(url, body);
  }
}
