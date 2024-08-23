import { Component, inject } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BookCatalogueService } from '../api-services/book-catalogue.service';
import { CatalogueBook } from '../catalogue-book';
import { CatalogueBookComponent } from './catalogue-book/catalogue-book.component';

@Component({
  selector: 'app-book-list-page',
  standalone: true,
  imports: [CatalogueBookComponent],
  templateUrl: './book-list-page.component.html',
  styleUrl: './book-list-page.component.scss',
})
export class BookListPageComponent {
  bookCatalogueService = inject(BookCatalogueService);
  searchPhrase = '';
  books: CatalogueBook[] = [];
  params: string = '';

  constructor(private router: Router, private route: ActivatedRoute) {}
  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      if (params['s']) this.searchForBooks(params['s']);
      this.params = params['s'];
    });
  }

  async searchForBooks(searchPhrase: string) {
    this.books = await this.bookCatalogueService.searchForBooks(searchPhrase);
  }
}
