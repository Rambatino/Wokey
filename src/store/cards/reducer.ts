import { SET_STATE, SetActionTypes } from './types'
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
    case SET_STATE:
      return parseJSON(JSON.parse(action.payload.message))
    default:
      return state
  }
}

export const parseJSON = (state: any): Cards => {
  return {
    baseCards: state.Issues.map(
      (item: any) =>
        ({
          ...item,
          linkedCards: item.pullRequests || [],
        } as BaseCard)
    ),
    linkedCards: state.pullRequests || [],
  }
}
