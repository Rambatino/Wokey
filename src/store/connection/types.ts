export const OPEN_STATE = 'REDUX_WEBSOCKET::OPEN'
export const CLOSED_STATE = 'REDUX_WEBSOCKET::CLOSED'
export const MESSAGE_STATE = 'REDUX_WEBSOCKET::MESSAGE'
export const CONNECT_STATE = 'REDUX_WEBSOCKET::CONNECT'
export const ERROR_STATE = 'REDUX_WEBSOCKET::ERROR'

export type ConnectionStateTypes =
  | typeof OPEN_STATE
  | typeof CLOSED_STATE
  | typeof CONNECT_STATE
  | typeof MESSAGE_STATE
  | typeof ERROR_STATE

interface SetConnectionStateAction {
  type: ConnectionStateTypes
}

export type SetConnectionTypes = SetConnectionStateAction
