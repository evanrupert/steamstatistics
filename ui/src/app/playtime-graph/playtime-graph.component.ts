import { Component, OnInit } from '@angular/core'
import * as CanvasJS from './../../assets/canvasjs.min.js'
import { TagPlaytime } from '../models/tag-playtime.model'
import { AppState } from '../app.state'
import { Store } from '@ngrx/store'
import { SetData } from '../actions/data.actions.js';

@Component({
  selector: 'app-playtime-graph',
  templateUrl: './playtime-graph.component.html',
  styleUrls: ['./playtime-graph.component.scss']
})
export class PlaytimeGraphComponent implements OnInit {

  data: TagPlaytime[]

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
    this.store.select('data').subscribe(data => {
      this.data = data

      const chart = new CanvasJS.Chart('chartContainer', this.getChartConfig())
      chart.render()
    })
  }

  getChartData() {
    return this.data.map(tp => {
      return { label: tp.tag, y: tp.playtime }
    }).slice(0, 20)
  }

  getChartConfig(): any {
    return {
      animationEnabled: true,
      exportEnabled: true,
      title: {
        text: 'Steam playtime vs game tags'
      },
      data: [{
        type: 'column',
        dataPoints: this.getChartData()
      }]
    }
  }
}
