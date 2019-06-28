import React, { Component } from 'react'
import './banner.scss'
import {
  MESSAGE_STATE,
  CONNECT_STATE,
  CLOSED_STATE,
  OPEN_STATE,
  ERROR_STATE,
} from '../store/connection/types'

type Props = {
  state: string
}
export default class Banner extends Component<Props> {
  config: { [key: string]: { colour: string; text: string } } = {
    [CONNECT_STATE]: {
      colour: '#FFCC00',
      text: 'Connecting...',
    },
    [CLOSED_STATE]: {
      colour: '#FF3B30',
      text: 'Connection lost',
    },
    [OPEN_STATE]: {
      colour: '#4CD964',
      text: 'Connected',
    },
    [ERROR_STATE]: {
      colour: '#FF3B30',
      text: 'Error',
    },
  }
  currentState: string | undefined

  render() {
    const { state } = this.props
    if (state === MESSAGE_STATE) {
      return null
    }
    return (
      <div
        className="banner"
        style={{ backgroundColor: this.config[state].colour }}
      >
        <p>{this.config[state].text}</p>
      </div>
    )
  }
}
