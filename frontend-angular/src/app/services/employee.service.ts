import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { Employee } from '../models/employee.model';

const baseUrl = 'http://localhost:8080/employees';

@Injectable({
  providedIn: 'root'
})
export class EmployeeService {

  constructor(private http: HttpClient) {}

  getAll(): Observable<Employee[]> {
    return this.http.get<{ data: { data: Employee[] } }>(baseUrl).pipe(
      map(response => response.data.data)  
    );
  }

  get(id: any): Observable<Employee> {
    return this.http.get<Employee>(`${baseUrl}/${id}`);
  }

  create(data: any): Observable<any> {
    return this.http.post(baseUrl, data);
  }

  update(id: any, data: any): Observable<any> {
    return this.http.put(`${baseUrl}/${id}`, data);
  }

  delete(id: any): Observable<any> {
    return this.http.delete(`${baseUrl}/${id}`);
  }
}
