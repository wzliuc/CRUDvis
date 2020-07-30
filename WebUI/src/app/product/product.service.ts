import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ProductDto } from '../Interfaces/productDto';
import { $ } from 'protractor';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  constructor(private http: HttpClient) { }

  GetProductsCat(id: number): Observable<ProductDto[]> {
    return this.http.get<ProductDto[]>(`http://localhost:7070/categories/${id}`).pipe();
  }

  GetProductById(id: number): Observable<ProductDto> {
    return this.http.get<ProductDto>(`http://localhost:7070/products/${id}`).pipe();
  }

  AddProduct(p:ProductDto): void {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };
    this.http.post<ProductDto>("http://localhost:7070/products", p, httpOptions).subscribe();
  }

  UpdateProduct(p:ProductDto): void {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };
    this.http.put<ProductDto>(`http://localhost:7070/products/${p.id}`, p, httpOptions).subscribe();
  }

  DeleteProduct(id: number): void {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };
    this.http.delete(`http://localhost:7070/products/${id}`, httpOptions).subscribe();
  }
}
