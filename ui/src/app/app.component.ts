import {Component, OnInit} from '@angular/core'
import {TagPlaytime} from './models/tag-playtime.model'
import {AppState} from './app.state'
import {Store} from '@ngrx/store'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  data: TagPlaytime[]
  displayGraph = false

 constructor(private store: Store<AppState>) {}

  ngOnInit() {
    this.store.select('data').subscribe(data => {
      this.data = data
      if (data) {
        this.displayGraph = true
      }
    })
  }

  resetData() {
    this.displayGraph = false
  }
}
