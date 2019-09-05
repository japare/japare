import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ProductsComponent } from './products/products.component';
import { SettingsComponent} from './settings/settings.component';
import {PageNotFoundComponent} from './page-not-found/page-not-found.component';

/* Setting a route:
* path: is the path in the URL
* component: which component will be displayed if user navigated to the path */
const routes: Routes = [
  {path: '', component: ProductsComponent, data: {title: 'Products'}},
  {path: 'settings', component: SettingsComponent, data: {title: 'Settings'}},
  {path: '**', component: PageNotFoundComponent, data: {title: 'Settings'}},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
