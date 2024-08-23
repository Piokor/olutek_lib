import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddToLibraryButtonComponent } from './add-to-library-button.component';

describe('AddToLibraryButtonComponent', () => {
  let component: AddToLibraryButtonComponent;
  let fixture: ComponentFixture<AddToLibraryButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddToLibraryButtonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddToLibraryButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
