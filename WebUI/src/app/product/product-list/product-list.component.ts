import { Component, OnInit } from '@angular/core';
import { ProductService } from '../product.service';
import { Observable } from 'rxjs';
import { ProductDto } from 'src/app/Interfaces/productDto';
import { Router, RouterLink, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'cz-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.css']
})
export class ProductListComponent implements OnInit {
  products$: Observable<ProductDto[]>;
  
  constructor(private productService: ProductService,
              private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(paraMap => {
      this.products$ = this.productService.GetProductsCat(+paraMap.get("id"));
    })
  }
}
