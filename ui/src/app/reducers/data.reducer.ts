import { TagPlaytime } from '../models/tag-playtime.model'
import * as DataStore from '../actions/data.actions'

const initialState = null

export function dataReducer(state: TagPlaytime[] = initialState, action: DataStore.Actions) {
  switch (action.type) {
    case DataStore.SET_DATA:
      return action.payload
    default:
      return state
  }
}
