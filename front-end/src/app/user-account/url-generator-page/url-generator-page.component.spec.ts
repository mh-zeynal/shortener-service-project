import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UrlGeneratorPageComponent } from './url-generator-page.component';

describe('UrlGeneratorPageComponent', () => {
  let component: UrlGeneratorPageComponent;
  let fixture: ComponentFixture<UrlGeneratorPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UrlGeneratorPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UrlGeneratorPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
