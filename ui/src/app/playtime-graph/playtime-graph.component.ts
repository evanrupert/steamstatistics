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
  barChartLabels: string[] = []
  barChartData: any = []

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
    this.store.select('data').subscribe(data => {
      this.data = data
      this.setBarChartData()
      this.setBarChartLabels()
    })
  }

  setBarChartData(): void {
    this.barChartData = [{
      data: this.data.slice(0, 20).map(tp => tp.playtime),
      label: 'Playtime Hours'
    }]
  }

  setBarChartLabels(): void {
    this.barChartLabels = this.data.slice(0, 20).map(tp => tp.tag)
  }
}
