import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UrlGenerationFormComponent } from './url-generation-form.component';

describe('UrlGenerationFormComponent', () => {
  let component: UrlGenerationFormComponent;
  let fixture: ComponentFixture<UrlGenerationFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UrlGenerationFormComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UrlGenerationFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
