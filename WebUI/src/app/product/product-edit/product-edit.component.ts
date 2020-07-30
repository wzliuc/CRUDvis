import { Component, OnInit } from '@angular/core';
import { ProductService } from '../product.service';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductDto } from 'src/app/Interfaces/productDto';
import { switchMap } from 'rxjs/operators';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { of } from 'rxjs';

@Component({
  selector: 'cz-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.css']
})
export class ProductEditComponent implements OnInit {
  productToEdit: ProductDto;
  productForm = this.fb.group({
    productId: this.fb.control({value:'', disabled: true}),
    productName: ['', Validators.required],
    productPrice: ['', Validators.required],
    productCategoryId: ['', Validators.required]
  });;

  constructor(private productService: ProductService,
              private route: ActivatedRoute,
              private fb: FormBuilder,
              private router: Router) { }

  ngOnInit(): void {
    this.route.paramMap
      .pipe(
        switchMap(param => {
          const id = +param.get('id');
          if(id == 0) {
            var newProduct: ProductDto = {
              id: 0,
              name: null,
              price: null,
              categoryId: null
            };
            return of(newProduct);
          } else {
            return this.productService.GetProductById(id);
          }
        })
      )
      .subscribe(p => {
        this.productToEdit = p;
        this.productForm.controls.productId.setValue(p.id);
        this.productForm.controls.productName.setValue(p.name);
        this.productForm.controls.productPrice.setValue(p.price);
        this.productForm.controls.productCategoryId.setValue(p.categoryId);
      });
  }

  checkValidity(control: string): boolean {
    const form = this.productForm;
    return form.get(control).invalid && 
            (form.get(control).dirty || 
            form.get(control).touched);
  }

  onSubmit(): void {
    const formValue = this.productForm.value;

    const productToSubmit: ProductDto = {
      id: this.productToEdit.id,
      name: formValue.productName,
      price: Number(formValue.productPrice),
      categoryId: Number(formValue.productCategoryId)
    }

    if(this.productToEdit.id == 0) {
      this.productService.AddProduct(productToSubmit);
    } else {
      this.productService.UpdateProduct(productToSubmit);
    }
    
    this.router.navigate(['/categories', productToSubmit.categoryId]);
  }

  deleteProduct(): void {
    this.productService.DeleteProduct(this.productToEdit.id);
    this.router.navigate(['/categories', this.productToEdit.categoryId]);
  }
}
