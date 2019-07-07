import React, { Component } from 'react'
import './notifications.scss'
import { Change } from '../store/cards/types'
import Notification from './Notification'

type Props = {
  changes: Array<Change>
}

export default class Notifications extends Component<Props> {
  render() {
    console.log(this.props)
    return (
      <div className="notifications">
        <p>Recent Activity</p>

        {this.props.changes.length > 0 ? (
          this.props.changes.map((change, i) => (
            <Notification key={i} change={change} />
          ))
        ) : (
          <p> No new alerts. You're all set </p>
        )}
      </div>
    )
  }
}
