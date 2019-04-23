export const SET_ISSUES = 'SET_ISSUES'
export const SET_PULLS = 'SET_PULLS'

export interface Item {
  id: string
  desc: string
  desc_html: string
  title: string
  subtitle: string
  url: string
  state: string
}

interface SetIssuesAction {
  type: typeof SET_ISSUES
  payload: Array<Item>
}

interface SetPullsAction {
  type: typeof SET_PULLS
  payload: Array<Item>
}

export type SetActionTypes = SetIssuesAction | SetPullsAction
