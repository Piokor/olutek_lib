import { Component, Input } from '@angular/core';
import { CatalogueBook } from '../../catalogue-book';
import { NgOptimizedImage } from '@angular/common';

@Component({
  selector: 'app-catalogue-book',
  standalone: true,
  imports: [NgOptimizedImage],
  templateUrl: './catalogue-book.component.html',
  styleUrl: './catalogue-book.component.scss',
})
export class CatalogueBookComponent {
  @Input({ required: true }) book!: CatalogueBook;
}
