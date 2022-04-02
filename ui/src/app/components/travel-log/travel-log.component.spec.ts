import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TravelLogComponent } from './travel-log.component';

describe('TravelLogComponent', () => {
  let component: TravelLogComponent;
  let fixture: ComponentFixture<TravelLogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TravelLogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TravelLogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
