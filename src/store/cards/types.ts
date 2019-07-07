import { ThunkDispatch } from 'redux-thunk'
import { AnyAction } from 'redux'

export interface Comment {
  count: number
  lastCommentLink: string
}

export interface Status {
  status: string
  link: string
}

export interface LinkedCard {
  id: string
  number: number
  approvalState: string
  ciStatus: Status
  title: string
  link: string
  branch: string
  repo: string

  comments: Comment
}

export interface BaseCard {
  id: string
  key: string
  title: string
  link: string
  state: string

  comments: Comment
  linkedCards: Array<LinkedCard>
}

export interface Change {
  type: string
  read: boolean
  message: string
  pullRequestID: string
  issueID: string
  createdAt: string
  seenAt: string
}

export interface Cards {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
  changes: Array<Change>
}

export const SET_STATE = 'REDUX_WEBSOCKET::MESSAGE'

interface SetStateAction {
  type: typeof SET_STATE
  payload: { event: MessageEvent; message: string; origin: string }
}

export type SetActionTypes = SetStateAction

export type CardThunkDispatch = ThunkDispatch<Cards, void, AnyAction>
