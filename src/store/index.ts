import { createStore, applyMiddleware, compose } from 'redux'
import thunk from 'redux-thunk'

import { combineReducers } from 'redux'
import cardsReducer from './cards/reducer'

const rootReducer = combineReducers({
  cards: cardsReducer,
})

export type AppState = ReturnType<typeof rootReducer>

export default function configureStore() {
  return createStore(
    rootReducer,
    compose(
      applyMiddleware(thunk),
      (window as any).__REDUX_DEVTOOLS_EXTENSION__
        ? (window as any).__REDUX_DEVTOOLS_EXTENSION__()
        : (f: any) => f
    )
  )
}
