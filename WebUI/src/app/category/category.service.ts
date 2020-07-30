import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CategoryDto } from '../Interfaces/categoryDto';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor(private http: HttpClient) { }

  GetAll(): Observable<CategoryDto[]> {
    return this.http.get<CategoryDto[]>("http://localhost:7070/categories")
  }
}
