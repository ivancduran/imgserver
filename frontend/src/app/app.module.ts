import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';

import { AuthService } from './auth.service';


const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'home', component: HomeComponent },
  // {
  //   path: 'channel', component: ChannelLayoutComponent, canActivate: [AuthGuard],
  //   children: [
  //     { path: '', redirectTo: 'videos', pathMatch: 'full' },
  //     { path: 'videos', component: ChannelVideoLayoutComponent },
  //     { path: 'upload', component: ChannelUploadLayoutComponent }
  //   ]
  // },
  {
    path: '**',
    redirectTo: '/login',
    pathMatch: 'full'
  }
];

const routing = RouterModule.forRoot(routes, { useHash: true });

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    routing
  ],
  providers: [
    AuthService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
