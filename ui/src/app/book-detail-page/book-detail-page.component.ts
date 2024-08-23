import { Component, inject } from '@angular/core';
import { CatalogueBook } from '../catalogue-book';
import { BookCatalogueService } from '../api-services/book-catalogue.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-book-detail-page',
  standalone: true,
  imports: [],
  templateUrl: './book-detail-page.component.html',
  styleUrl: './book-detail-page.component.scss',
})
export class BookDetailPageComponent {
  book: CatalogueBook | undefined;
  bookCatalogueService: BookCatalogueService = inject(BookCatalogueService);

  constructor(private route: ActivatedRoute) {
    const bookId = route.snapshot.paramMap.get('bookId');
    if (bookId) this.fetchBook(bookId);
  }

  async fetchBook(bookId: string) {
    this.book = await this.bookCatalogueService.getBook(bookId);
  }
}
