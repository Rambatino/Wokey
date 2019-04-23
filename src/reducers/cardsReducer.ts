import initialState from './initialState'
import { SET_ISSUES, SET_PULLS, SetActionTypes } from '../actions/actionTypes'
import { Store, BaseCard, LinkedCard } from '../store/types'

export default function cards(
  state: Store = initialState,
  action: SetActionTypes
): Store {
  switch (action.type) {
    case SET_ISSUES:
      return {
        linkedCards: [...state.linkedCards],
        baseCards: action.payload.map(item => {
          return { item: item, linkedCards: [] } as BaseCard
        }),
        config: state.config,
      }
    case SET_PULLS:
      return {
        linkedCards: action.payload.map(item => {
          return { item: item } as LinkedCard
        }),
        baseCards: [...state.baseCards],
        config: state.config,
      }
    default:
      return state
  }
}
