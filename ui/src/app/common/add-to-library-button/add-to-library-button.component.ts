import { Component, Input } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-add-to-library-button',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './add-to-library-button.component.html',
  styleUrl: './add-to-library-button.component.scss',
})
export class AddToLibraryButtonComponent {
  @Input({ required: true }) isInLibrary!: boolean;
  @Input() onAdding: (() => void) | undefined;

  onClick() {}
}
