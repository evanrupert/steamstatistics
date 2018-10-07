import {Component, OnInit} from '@angular/core'
import {TagPlaytime} from './models/tag-playtime.model'
import {AppState} from './app.state'
import {Store} from '@ngrx/store'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  data: TagPlaytime[]

 constructor(private store: Store<AppState>) {
   this.store.select('data').subscribe(data => {
     this.data = data
   })
 }

  ngOnInit() {
  }

  printData(): void {
    console.log(this.data)
  }
}
