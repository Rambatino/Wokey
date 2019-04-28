import { ThunkDispatch } from 'redux-thunk'
import { AnyAction } from 'redux'

export interface Item {
  id: string
  desc?: string
  descHtml?: string
  descMarkdown?: string
  title: string
  subtitle: string
  url: string
  state: string
  branch?: string
  repo?: string
}

export interface LinkedCard {
  item: Item
}

export interface BaseCard {
  item: Item
  linkedCards: Array<LinkedCard>
}

export interface Cards {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
}

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

export type CardThunkDispatch = ThunkDispatch<Cards, void, AnyAction>
