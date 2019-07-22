import React, { Component } from 'react'
import { Change } from '../store/cards/types'
import './components.scss'
import moment from 'moment'

type Props = {
  change: Change
}

export default class Notification extends Component<Props> {
  render() {
    return (
      <div className="notification">
        <p className="createdAt">{this.time(this.props.change.createdAt)}</p>
        <p className="message">{this.props.change.message}</p>
      </div>
    )
  }

  time = (createdAt: string) => {
    let time = moment(createdAt).fromNow()
    if (time === 'a few seconds ago') {
      time = 'just now'
    }
    return time
  }
}
