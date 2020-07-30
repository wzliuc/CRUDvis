import { Component, OnInit } from '@angular/core';
import { ProductService } from '../product.service';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { ProductDto } from 'src/app/Interfaces/productDto';

@Component({
  selector: 'cz-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.css']
})
export class ProductDetailComponent implements OnInit {
  product$: Observable<ProductDto>
  constructor(private productService: ProductService,
              private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(paraMap => {
      this.product$ = this.productService.GetProductById(+paraMap.get('id'))
    })
  }

}
