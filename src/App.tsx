import React, { Component } from 'react'
import logo from './logo.svg'
import './App.css'
import Space from './containers/Space'
import { BaseCardItem, LinkedCardItem } from './data/types'

type State = {
  linkedCards: Array<LinkedCardItem>
  baseCards: Array<BaseCardItem>
}

class App extends Component<{}, State> {
  state = {
    linkedCards: [],
    baseCards: [],
  }
  componentDidMount() {
    fetch('/issues')
      .then(res => res.json())
      .then(res => {
        this.setState({
          baseCards: res as Array<BaseCardItem>,
        })
      })
    fetch('/prs')
      .then(res => res.json())
      .then(res => {})
  }

  render() {
    return (
      <div className="App">
        <Space data={this.state} />
      </div>
    )
  }
}

export default App
