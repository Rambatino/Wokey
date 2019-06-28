import {
  CLOSED_STATE,
  SetConnectionTypes,
  OPEN_STATE,
  CONNECT_STATE,
  MESSAGE_STATE,
} from './types'

const initialState: string = CLOSED_STATE

export default function connectionReducer(
  state: string = initialState,
  action: SetConnectionTypes
): string {
  switch (action.type) {
    case CLOSED_STATE:
      return CLOSED_STATE
    case OPEN_STATE:
      return OPEN_STATE
    case CONNECT_STATE:
      return CONNECT_STATE
    case MESSAGE_STATE:
      return MESSAGE_STATE
  }
  return state
}
