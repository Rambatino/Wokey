import { ThunkDispatch } from 'redux-thunk'
import { Action } from 'redux'
import { Store } from '../store/types'

import { Item, SetActionTypes, SET_ISSUES, SET_PULLS } from './actionTypes'

function issuesUrl() {
  return '/issues'
}

function pullsUrl() {
  return '/pulls'
}

export function setIssues(issues: Array<Item>): SetActionTypes {
  return {
    type: SET_ISSUES,
    payload: issues,
  }
}

export function setPulls(pulls: Array<Item>): SetActionTypes {
  return {
    type: SET_PULLS,
    payload: pulls,
  }
}

export function fetchPulls() {
  return async (dispatch: ThunkDispatch<Store, void, Action>) => {
    const response = await fetch(pullsUrl())
    const json = await response.json()
    return dispatch(setPulls(json))
  }
}

export function fetchIssues() {
  return async (dispatch: ThunkDispatch<Store, void, Action>) => {
    const response = await fetch(issuesUrl())
    const json = await response.json()
    return dispatch(setIssues(json))
  }
}
