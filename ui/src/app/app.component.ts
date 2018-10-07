import {Component, OnInit} from '@angular/core'
import {TagPlaytime} from './models/tag-playtime.model'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'SteamStatistics'
  data: TagPlaytime[]

  ngOnInit() {

  }
}
