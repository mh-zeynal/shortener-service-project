import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserLinksPageComponent } from './user-links-page.component';

describe('UserLinksPageComponent', () => {
  let component: UserLinksPageComponent;
  let fixture: ComponentFixture<UserLinksPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserLinksPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserLinksPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
