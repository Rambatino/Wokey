import React, { Component } from 'react'
import './containers.scss'
import { Change } from '../store/cards/types'
import Notification from '../components/Notification'

type Props = {
  changes: Array<Change>
}

type State = { open: boolean }

export default class Notifications extends Component<Props, State> {
  render() {
    return (
      <div className="notifications">
        <p id="notificationsTitle">Recent Activity</p>
        {this.props.changes.length > 0 && <p id="clearAll">Clear All</p>}
        {this.props.changes.map((change, i) => (
          <Notification key={i} change={change} />
        ))}
      </div>
    )
  }
}
