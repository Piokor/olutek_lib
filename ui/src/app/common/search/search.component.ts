import { Component } from '@angular/core';
import { SearchBooksService } from '../../search-books.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'search',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.scss',
})
export class SearchComponent {
  searchPhrase = '';

  constructor(private searchService: SearchBooksService) {}

  onSearch() {
    this.searchService.onSearch(this.searchPhrase);
  }
}
