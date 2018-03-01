import { TestBed, inject } from '@angular/core/testing';

import { JsonViewerService } from './json-viewer.service';

describe('JsonViewerService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [JsonViewerService]
    });
  });

  it('should be created', inject([JsonViewerService], (service: JsonViewerService) => {
    expect(service).toBeTruthy();
  }));
});
