import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'

import { combineReducers } from 'redux'
import cardsReducer from './cards/reducer'
import connectionReducer from './connection/reducer'

import reduxWebsocket from '@giantmachines/redux-websocket'

const rootReducer = combineReducers({
  cards: cardsReducer,
  connectedState: connectionReducer,
})

// Create the middleware instance.
const reduxWebsocketMiddleware = reduxWebsocket({ reconnectInterval: 12 })

export type AppState = ReturnType<typeof rootReducer>

export default function configureStore() {
  return createStore(
    rootReducer,
    compose(
      applyMiddleware(thunk, reduxWebsocketMiddleware),
      (window as any).__REDUX_DEVTOOLS_EXTENSION__
        ? (window as any).__REDUX_DEVTOOLS_EXTENSION__()
        : (f: any) => f
    )
  )
}
