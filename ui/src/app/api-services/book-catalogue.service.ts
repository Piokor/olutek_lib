import { Injectable } from '@angular/core';
import { CatalogueBook } from './catalogue-book';

@Injectable({
  providedIn: 'root',
})
export class BookCatalogueService {
  async searchForBooks(searchPhrase: string): Promise<CatalogueBook[]> {
    if (searchPhrase === '') return [];
    const url = `http://localhost:8081/api/books-catalogue/books/${searchPhrase}`;
    const data = await fetch(url);
    return (await data.json()) ?? [];
  }

  async getBook(bookId: string): Promise<CatalogueBook> {
    const url = `http://localhost:8081/api/books-catalogue/book/${bookId}`;
    const data = await fetch(url);
    return (await data.json()) ?? null;
  }

  constructor() {}
}
