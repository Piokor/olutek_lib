import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class BookLibraryService {
  async searchForBooks(searchPhrase: string): Promise<CatalogueBook[]> {
    if (searchPhrase === '') return [];
    const url = `http://localhost:8081/api/books-catalogue/books/${searchPhrase}`;
    const data = await fetch(url);
    return (await data.json()) ?? [];
  }
}
