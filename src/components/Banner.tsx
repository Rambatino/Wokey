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
type State = {
  open: boolean
}

export default class Banner extends Component<Props, State> {
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
  state: State = { open: false }
  render() {
    const { state } = this.props
    if (state === MESSAGE_STATE) {
      this.setState({ open: false })

      return null
    }
    this.setState({ open: true })
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
