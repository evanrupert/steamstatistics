import { Action } from '@ngrx/store'
import { TagPlaytime } from '../models/tag-playtime.model'

export const SET_DATA = '[DATA] Set'

export class SetData implements Action {
  readonly type = SET_DATA

  constructor(public payload: TagPlaytime[]) {}
}

export type Actions = SetData
