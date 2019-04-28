import { SET_ISSUES, SET_PULLS, SetActionTypes } from './types'
import { Cards, BaseCard, LinkedCard } from './types'
import { filterCards } from './formatter'

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
      console.log(
        filterCards(
          action.payload.map(item => {
            return { item, linkedCards: [] } as BaseCard
          }),
          [...state.linkedCards]
        )
      )
      return filterCards(
        action.payload.map(item => {
          return { item, linkedCards: [] } as BaseCard
        }),
        [...state.linkedCards]
      )
    case SET_PULLS:
      console.log(
        filterCards(
          action.payload.map(item => {
            return { item, linkedCards: [] } as BaseCard
          }),
          [...state.linkedCards]
        )
      )
      return filterCards(
        [...state.baseCards],
        action.payload.map(item => {
          return { item } as LinkedCard
        })
      )
    default:
      return state
  }
}
