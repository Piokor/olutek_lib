import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CatalogueBookComponent } from './catalogue-book.component';

describe('CatalogueBookComponent', () => {
  let component: CatalogueBookComponent;
  let fixture: ComponentFixture<CatalogueBookComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CatalogueBookComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CatalogueBookComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
