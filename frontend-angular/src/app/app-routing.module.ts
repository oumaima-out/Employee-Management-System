import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from "./features/home/home.component";
import { LoginComponent } from "./features/login/login.component";
import { AdminTemplateComponent } from "./features/admin-template/admin-template.component";
import { AuthGuard } from "./guards/auth.guard";
import { AuthorizationGuard } from "./guards/authorization.guard";
import { EmployeesComponent } from './features/manageEmployee/list-employees/list-employees.component';

const routes: Routes = [
  {path : "", component : LoginComponent},
  {path : "login", component : LoginComponent},
  {path : "admin", component : AdminTemplateComponent,
    canActivate : [AuthGuard],
    children : [
      {path : "home", component : HomeComponent},
      {
        path : "employees", component : EmployeesComponent,
        canActivate : [AuthorizationGuard], data : {roles : ['ADMIN']}
      }
    ]},

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

