import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProductsComponent} from './products/products.component';
/* Setting a route:
* path: is the path in the URL
* component: which component will be displayed if user navigated to the path */
const routes: Routes = [
  {path: 'products', component: ProductsComponent},
  // {path: '', redirectTo: '/dashboard', pathMatch: 'full'},
  // {path: 'detail/:id', component: HabitDetailComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
