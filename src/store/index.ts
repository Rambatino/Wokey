import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'

import { combineReducers } from 'redux'
import cardsReducer from './cards/reducer'

import reduxWebsocket from '@giantmachines/redux-websocket'

const rootReducer = combineReducers({
  cards: cardsReducer,
})

// Create the middleware instance.
const reduxWebsocketMiddleware = reduxWebsocket()

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
