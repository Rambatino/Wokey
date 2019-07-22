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
        <p className="createdAt">
          {moment(this.props.change.createdAt).fromNow()}
        </p>
        <p className="message">{this.props.change.message}</p>
      </div>
    )
  }

  time = (createdAt: Date) => {
    let time = moment(createdAt).fromNow()
    if (time === 'a few seconds ago') {
      time = 'just now'
    }
    return time
  }
}
