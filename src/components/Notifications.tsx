import React, { Component } from 'react'
import './components.scss'
import { Change } from '../store/cards/types'
import Notification from './Notification'

type Props = {
  changes: Array<Change>
}

type State = { open: boolean }

export default class Notifications extends Component<Props, State> {
  state: State = { open: false }
  width = 300
  margin = 50

  slide = () => {
    this.setState({ open: !this.state.open })
  }

  render() {
    const { open } = this.state
    return (
      <div
        className="notifications"
        style={{
          width: this.width,
          right: open ? '0px' : this.margin - this.width,
        }}
      >
        <p>Recent Activity</p>
        <button onClick={this.slide}> {open ? '>' : '<'} </button>
        {this.props.changes.map((change, i) => (
          <Notification key={i} change={change} />
        ))}
      </div>
    )
  }
}
