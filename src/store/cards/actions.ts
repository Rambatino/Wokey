import { ThunkAction } from 'redux-thunk'
import { AnyAction } from 'redux'
import {
  Item,
  Cards,
  SetActionTypes,
  SET_ISSUES,
  SET_PULLS,
  CardThunkDispatch,
} from './types'

// import pulls from '../../stories/pulls.json'
// import issues from '../../stories/issues.json'

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
  const response = await fetch(issuesUrl())
  const json = await response.json()
  console.log(json)
  // const json = await Promise.resolve(issues)
  return dispatch(setIssues(json))
}

export const fetchPulls = (): ThunkAction<
  void,
  Cards,
  void,
  AnyAction
> => async (dispatch: CardThunkDispatch) => {
  const response = await fetch(pullsUrl())
  const json = await response.json()
  console.log(json)
  // const json = await Promise.resolve(pulls)
  return dispatch(setPulls(json))
}
