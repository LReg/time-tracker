import { TestBed } from '@angular/core/testing';

import { TtService } from './tt.service';

describe('TtService', () => {
  let service: TtService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TtService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
