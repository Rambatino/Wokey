import React, { Component } from 'react'
import logo from './logo.svg'
import './App.css'
import Space from './containers/Space'
import { BaseCard, LinkedCard } from './store/types'

type State = {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
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
        this.setState({})
      })
      .catch(err => console.log(err))
    fetch('/prs')
      .then(res => res.json())
      .then(res => {})
      .catch(err => console.log(err))
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
