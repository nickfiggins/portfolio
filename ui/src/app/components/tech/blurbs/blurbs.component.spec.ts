import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BlurbsComponent } from './blurbs.component';

describe('BlurbsComponent', () => {
  let component: BlurbsComponent;
  let fixture: ComponentFixture<BlurbsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BlurbsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BlurbsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
