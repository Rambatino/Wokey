import React, { Component } from 'react'
import './App.css'
import Board from './containers/Board'
import { Cards, CardThunkDispatch } from './store/cards/types'
import { fetchPulls, fetchIssues } from './store/cards/actions'
import { AppState } from './store'
import { connect } from 'react-redux'

// https://github.com/reduxjs/redux-thunk/issues/213
interface AppProps {
  fetchPulls: () => void
  fetchIssues: () => void
  cards: Cards
}

class App extends Component<AppProps> {
  componentDidMount() {
    this.props.fetchPulls()
    this.props.fetchIssues()

    const conn = new WebSocket('ws://' + document.location.host + '/ws')

    conn.onopen = () => {
      var item = document.createElement('div')
      item.innerHTML = '<b>OPENED</b>'
      document.body.appendChild(item)
    }

    conn.onclose = () => {
      var item = document.createElement('div')
      item.innerHTML = '<b>Connection closed.</b>'
      document.body.appendChild(item)
    }

    conn.onmessage = evt => {
      var messages = evt.data.split('\n')
      for (var i = 0; i < messages.length; i++) {
        var item = document.createElement('div')
        item.innerText = messages[i]
        document.body.appendChild(item)
      }
    }
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
  fetchPulls: () => dispatch(fetchPulls()),
  fetchIssues: () => dispatch(fetchIssues()),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
