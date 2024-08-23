import { Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { BookListPageComponent } from './book-list-page/book-list-page.component';
import { BookDetailPageComponent } from './book-detail-page/book-detail-page.component';

export const routes: Routes = [
  {
    path: '',
    component: HomePageComponent,
    title: 'Home page',
  },
  {
    path: 'search',
    component: BookListPageComponent,
    title: 'Book search',
  },
  {
    path: 'book/:bookId',
    component: BookDetailPageComponent,
    title: 'Book page',
  },
];
