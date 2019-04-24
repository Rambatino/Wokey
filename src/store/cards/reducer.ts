import { SET_ISSUES, SET_PULLS, SetActionTypes } from './types'
import { Cards, BaseCard, LinkedCard } from './types'

const initialState: Cards = {
  linkedCards: [],
  baseCards: [],
}

export default function cardsReducer(
  state: Cards = initialState,
  action: SetActionTypes
): Cards {
  switch (action.type) {
    case SET_ISSUES:
      return {
        linkedCards: [...state.linkedCards],
        baseCards: action.payload.map(item => {
          return { item: item, linkedCards: [] } as BaseCard
        }),
      }
    case SET_PULLS:
      return {
        linkedCards: action.payload.map(item => {
          return { item: item } as LinkedCard
        }),
        baseCards: [...state.baseCards],
      }
    default:
      return state
  }
}
