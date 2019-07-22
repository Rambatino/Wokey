import React, { Component } from 'react'
import './components.scss'
import { Change } from '../store/cards/types'
import Notification from './Notification'

type Props = {
  changes: Array<Change>
}

type State = { open: boolean }

export default class Notifications extends Component<Props, State> {
  render() {
    if (this.props.changes === undefined || !this.props.changes.length) {
      return null
    }
    return (
      <div className="notifications">
        <p className="notificationsTitle">Recent Activity</p>
        {this.props.changes.map((change, i) => (
          <Notification key={i} change={change} />
        ))}
      </div>
    )
  }
}
