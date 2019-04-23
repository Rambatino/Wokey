import { Item } from '../store/types'

export const SET_ISSUES = 'SET_ISSUES'
export const SET_PULLS = 'SET_PULLS'

interface SetIssuesAction {
  type: typeof SET_ISSUES
  payload: Array<Item>
}

interface SetPullsAction {
  type: typeof SET_PULLS
  payload: Array<Item>
}

export type SetActionTypes = SetIssuesAction | SetPullsAction
