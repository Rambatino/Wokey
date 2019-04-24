import { ThunkAction } from 'redux-thunk'
import { Action, AnyAction } from 'redux'
import {
  Item,
  Cards,
  SetActionTypes,
  SET_ISSUES,
  SET_PULLS,
  CardThunkDispatch,
} from './types'

import pulls from './pulls.json'
import issues from './issues.json'

function issuesUrl() {
  return '/issues'
}

function pullsUrl() {
  return '/pulls'
}

function setIssues(issues: Array<Item>): SetActionTypes {
  return {
    type: SET_ISSUES,
    payload: issues,
  }
}

function setPulls(pulls: Array<Item>): SetActionTypes {
  return {
    type: SET_PULLS,
    payload: pulls,
  }
}

export const fetchIssues = (): ThunkAction<
  void,
  Cards,
  void,
  AnyAction
> => async (dispatch: CardThunkDispatch) => {
  // const response = await fetch(issuesUrl())
  // const json = await response.json()
  const json = issues
  return dispatch(setIssues(json))
}

export const fetchPulls = (): ThunkAction<
  void,
  Cards,
  void,
  AnyAction
> => async (dispatch: CardThunkDispatch) => {
  // const response = await fetch(pullsUrl())
  // const json = await response.json()
  const json = issues
  return dispatch(setPulls(json))
}
