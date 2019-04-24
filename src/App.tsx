import React, { Component } from 'react'
import './App.css'
import Space from './containers/Space'
import { Cards, CardThunkDispatch } from './store/cards/types'
import { fetchPulls, fetchIssues } from './store/cards/actions'
import { AppState } from './store'
import { connect } from 'react-redux'

// https://github.com/reduxjs/redux-thunk/issues/213
interface AppProps {
  fetchPulls: typeof fetchPulls
  fetchIssues: typeof fetchIssues
  cards: Cards
}

class App extends Component<AppProps> {
  componentDidMount() {
    this.props.fetchPulls()
    this.props.fetchIssues()
  }

  render() {
    return (
      <div className="App">
        <Space data={this.props.cards} />
      </div>
    )
  }
}

const mapStateToProps = (state: AppState) => ({
  cards: state.cards,
})

const mapDispatchToProps = (dispatch: CardThunkDispatch) => ({
  fetchPulls: () => dispatch(fetchPulls),
  fetchIssues: () => dispatch(fetchIssues),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
