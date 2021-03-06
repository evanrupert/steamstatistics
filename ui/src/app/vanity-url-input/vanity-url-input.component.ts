import { Component, OnInit } from '@angular/core';
import axios from 'axios'
import { ServerResponse } from '../models/server-response.model'
import {TagPlaytime} from '../models/tag-playtime.model'
import {AppState} from '../app.state'
import {Store} from '@ngrx/store'
import {SetData} from '../actions/data.actions'

@Component({
  selector: 'app-vanity-url-input',
  templateUrl: './vanity-url-input.component.html',
  styleUrls: ['./vanity-url-input.component.scss']
})
export class VanityUrlInputComponent implements OnInit {

  inputError: string = null

  constructor(private store: Store<AppState>) { }

  ngOnInit() {
  }

  getData(vanityUrl: string): void {
    axios.request<ServerResponse<TagPlaytime[]>>({
      url: 'http://localhost:8080/api/data/' + vanityUrl
    }).then((resp) => {
      this.handleData(resp.data)
    })
  }

  handleData(serverResponse: ServerResponse<TagPlaytime[]>): void {
    if (!serverResponse.ok) {
      this.inputError = serverResponse.error
    } else {
      this.store.dispatch(new SetData(serverResponse.data))
    }
  }
}
