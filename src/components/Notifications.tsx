import React, { Component } from 'react'
import './notifications.scss'
import { Change } from '../store/cards/types'
import Notification from './Notification'

type Props = {
  changes: Array<Change>
}

export default class Notifications extends Component<Props> {
  render() {
    if (this.props.changes.length == 0) {
      return null
    }
    return (
      <div className="notifications">
        <p>Recent Activity</p>
        {this.props.changes.map((change, i) => (
          <Notification key={i} change={change} />
        ))}
      </div>
    )
  }
}
