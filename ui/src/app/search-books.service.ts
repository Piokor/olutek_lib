import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class SearchBooksService {
  onSearch: (searchPhrase: string) => void;
  private defaultOnSearch = (searchPhrase: string) =>
    this.router.navigate(['/search'], { queryParams: { s: searchPhrase } });

  setDefaultHandler() {
    this.onSearch = this.defaultOnSearch;
  }
  setHandler(handler: (searchPhrase: string) => void) {
    this.onSearch = handler;
  }

  constructor(private router: Router) {
    this.onSearch = this.defaultOnSearch;
  }
}
