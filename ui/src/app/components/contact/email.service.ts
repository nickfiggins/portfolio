import { Injectable } from '@angular/core';
import {
  HttpClient, HttpHeaders
} from "@angular/common/http";
import { environment } from "../../../environments/environment.prod";

@Injectable({
  providedIn: 'root'
})
export class EmailService {

  constructor(private http : HttpClient) { }

  sendEmail(emailVM: any){
    return this.http.post('/api/contact', emailVM);
  }

  sendEmailLambda(emailVM: any) {
    const headers = new HttpHeaders();
    return this.http.post(environment.apiGatewayUrl, emailVM);
  }

}
