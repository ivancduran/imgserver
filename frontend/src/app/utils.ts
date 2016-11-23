// depending on the env mode, enable prod mode or add debugging modules
if (process.env.ENV === 'build') {
  var URI = 'http://amaterasu.co:8000/';
} else {
  var URI = 'http://localhost:4200/api/';
}

export class Utils {

  public apiUrl: string = URI;

  // public apiUrl: string = 'http://amaterasu.co/';
  // public apiUrl: string = 'http://domain.dev:8080/';

  constructor() {

  }

}
