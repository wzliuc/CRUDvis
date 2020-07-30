import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductDetailComponent } from './product-detail/product-detail.component';
import { RouterModule, Routes } from '@angular/router';
import { ProductListComponent } from './product-list/product-list.component';
import { ProductEditComponent } from './product-edit/product-edit.component';
import { ReactiveFormsModule } from '@angular/forms';

const routes: Routes = [
  { path:'products',
    children: [
      { path: '', component: ProductListComponent},
      { path: ':id', component: ProductDetailComponent},
      { path: 'edit/:id', component: ProductEditComponent}
    ]
  }
];

@NgModule({
  declarations: [
    ProductDetailComponent, 
    ProductListComponent, 
    ProductEditComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    ReactiveFormsModule
  ]
})
export class ProductModule { }
