import React, { Component } from 'react'
import { Change } from '../store/cards/types'
import './components.scss'

type Props = {
  change: Change
}

export default class Notification extends Component<Props> {
  render() {
    return <div className="notification"> {this.props.change.message}</div>
  }
}
