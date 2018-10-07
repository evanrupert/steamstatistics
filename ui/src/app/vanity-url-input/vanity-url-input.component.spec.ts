import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { VanityUrlInputComponent } from './vanity-url-input.component';

describe('VanityUrlInputComponent', () => {
  let component: VanityUrlInputComponent;
  let fixture: ComponentFixture<VanityUrlInputComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ VanityUrlInputComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VanityUrlInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
