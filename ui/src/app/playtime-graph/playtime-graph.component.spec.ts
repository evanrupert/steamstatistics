import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlaytimeGraphComponent } from './playtime-graph.component';

describe('PlaytimeGraphComponent', () => {
  let component: PlaytimeGraphComponent;
  let fixture: ComponentFixture<PlaytimeGraphComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlaytimeGraphComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlaytimeGraphComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
