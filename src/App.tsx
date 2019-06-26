import React, { Component } from 'react'
import './App.css'
import Board from './containers/Board'
import { Cards, CardThunkDispatch } from './store/cards/types'
import { AppState } from './store'
import { connect } from 'react-redux'
import { connect as websocketConnect } from '@giantmachines/redux-websocket'

// https://github.com/reduxjs/redux-thunk/issues/213
interface AppProps {
  cards: Cards
  connect: () => void
}

class App extends Component<AppProps> {
  componentDidMount() {
    this.props.connect()
  }

  render() {
    return (
      <div className="App">
        <Board data={this.props.cards} />
      </div>
    )
  }
}

const mapStateToProps = (state: AppState) => ({
  cards: state.cards,
})

const mapDispatchToProps = (dispatch: CardThunkDispatch) => ({
  connect: () =>
    dispatch(websocketConnect('ws://' + document.location.host + '/ws')),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
