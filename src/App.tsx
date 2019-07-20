import React, { Component } from 'react'
import './App.css'
import Board from './containers/Board'
import { Cards, CardThunkDispatch, Change } from './store/cards/types'
import { AppState } from './store'
import { connect } from 'react-redux'
import { connect as websocketConnect } from '@giantmachines/redux-websocket'
import Notifications from './components/Notifications'
import Toolbar from './components/Toolbar'

// https://github.com/reduxjs/redux-thunk/issues/213
interface AppProps {
  cards: Cards
  notifications: Array<Change>
  connect: () => void
  connectedState: string
}

class App extends Component<AppProps> {
  componentDidMount() {
    this.props.connect()
  }

  render() {
    return (
      <div className="App">
        <Toolbar />
        <Notifications changes={this.props.notifications} />
        <Board data={this.props.cards} />
      </div>
    )
  }
}

const mapStateToProps = (state: AppState) => ({
  cards: state.cards,
  notifications: state.cards.changes,
  connectedState: state.connectedState,
})

const mapDispatchToProps = (dispatch: CardThunkDispatch) => ({
  connect: () =>
    dispatch(websocketConnect('ws://' + document.location.host + '/ws')),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
