import { Component } from '@angular/core';

import { Router } from '@angular/router';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {

  loginForm: any;
  error: boolean = false;
  errorMessage: string;

  constructor(
    public auth: AuthService,
    public router: Router
  ) {

    this.loginForm = {
      username: '',
      password: ''
    }

  }

  onSubmit(value: any) {

    this.auth.login(value.username, value.password)
      .subscribe(data => {
        if (data.response) {

          this.auth.loginSetData(data);
          this.router.navigate(['/home']);

        } else {
          this.error = true;
          this.errorMessage = data.error;
        }
      });

  }
}
