import { TestBed } from '@angular/core/testing';

import { TravelHelperService } from './travelHelper.service';

describe('TravelHelperService', () => {
  let service: TravelHelperService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TravelHelperService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
