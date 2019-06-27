import { ThunkDispatch } from 'redux-thunk'
import { AnyAction } from 'redux'

export interface Comment {
  id: string
  comment: string
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

  comments: Array<Comment>
}

export interface BaseCard {
  id: string
  key: string
  title: string
  link: string
  state: string

  comments: Array<Comment>
  linkedCards: Array<LinkedCard>
}

export interface Cards {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
}

export const SET_STATE = 'REDUX_WEBSOCKET::MESSAGE'

interface SetStateAction {
  type: typeof SET_STATE
  payload: { event: MessageEvent; message: string; origin: string }
}

export type SetActionTypes = SetStateAction

export type CardThunkDispatch = ThunkDispatch<Cards, void, AnyAction>
