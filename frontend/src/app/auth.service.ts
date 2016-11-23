import { Injectable } from '@angular/core';
import { Http, Headers } from '@angular/http';
// import { Observable } from 'rxjs/Observable';
import { Utils } from './utils';

@Injectable()
export class AuthService {

  token: string;
  errorMessage: string;
  public http;
  private u;

  constructor(http: Http) {
    this.http = http;
    this.u = new Utils();
    this.token = localStorage.getItem('token');
  }

  login(username: String, password: String): any {

    let req = this.http.post(this.u.apiUrl + 'v1/auth/login', JSON.stringify({
      email: username,
      pass: password
    }), {
        headers: new Headers({
          'Content-Type': 'text/plain'
        })
      })
      .map(res => res.json());

    return req;

  }

  loginSetData(data) {
    console.log('enter login data');
    console.log(data);

    if (data.response === true) {
      this.token = data.token;
      // localStorage.setItem('api_key', data.api_key);
      localStorage.setItem('token', data.token);
      localStorage.setItem('status', data.status);
      localStorage.setItem('role', data.role);
      localStorage.setItem('username', data.username);
    } else {
      this.errorMessage = data.error;
    }

  }

  logout(token: String) {
    let req = this.http.get(this.u.apiUrl + 'v1/auth/logout')
      .map(res => res.json());

    return req;
  }

  logoutData(data) {
    console.log(data.response);
    if (data.response) {
      this.token = undefined;
      localStorage.setItem('token', '');
      localStorage.setItem('status', '');
      localStorage.setItem('role', '');
      localStorage.setItem('username', '');
      localStorage.removeItem('token');
      localStorage.removeItem('status');
      localStorage.removeItem('role');
      localStorage.removeItem('username');
    }
  }

  register(value: any): any {

    if (value.password === value.repassword) {
      let req = this.http.post(this.u.apiUrl + 'v1/auth/register', JSON.stringify({
        username: value.username,
        email: value.email,
        pass: value.password,
        role: 0,
        active: 1
      }), {
          headers: new Headers({
            'Content-Type': 'text/plain'
          })
        })
        .map(res => res.json());

      return req;

    }


  }


}
