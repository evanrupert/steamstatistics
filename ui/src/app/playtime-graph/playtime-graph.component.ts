import { Component, OnInit } from '@angular/core'
import { TagPlaytime } from '../models/tag-playtime.model'
import { AppState } from '../app.state'
import { Store } from '@ngrx/store'

@Component({
  selector: 'app-playtime-graph',
  templateUrl: './playtime-graph.component.html',
  styleUrls: ['./playtime-graph.component.scss']
})
export class PlaytimeGraphComponent implements OnInit {

  data: TagPlaytime[]
  barChartLabels: string[] = ['2006', '2007', '2008', '2009', '2010', '2011', '2012']
  barChartData: any = [{
    data: [1, 2, 3, 4, 5, 6, 7],
    label: 'Number'
  }]

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
    this.store.select('data').subscribe(data => {
      this.data = data
      this.setBarChartData()
    })
  }

  setBarChartData(): void {

  }

  setBarChartLables(): void {
    this.barChartLabels = this.data.slice(0, 20).map((tagPlaytime => tagPlaytime.tag))
  }
}
