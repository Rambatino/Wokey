import React, { Component } from 'react'
import { Change } from '../store/cards/types'
import './components.scss'
import moment from 'moment'

type Props = {
  change: Change
}

export default class Notification extends Component<Props> {
  tick = (
    <svg
      width="10px"
      height="12px"
      viewBox="0 0 10 12"
      stroke-width="1"
      fill="none"
      fill-rule="evenodd"
      stroke="#618FBA"
    >
      <path d="M 0 1 a 1 1 0 0 1 1 -1 h 14 a 1 1 0 0 1 1 1 v 10 a 1 1 0 0 1 -1 1 H 7.5 L 4 15.5 V 12 H 1 a 1 1 0 0 1 -1 -1 V 1 Z m 1 0 v 10 h 4 v 2 l 2 -2 h 8 V 1 H 1 Z m 7.5 3 h 2 v 1 h -2 v 2 h -1 V 5 h -2 V 4 h 2 V 2 h 1 v 2 Z m 2 5 h -5 V 8 h 5 v 1 Z" />
    </svg>
  )
  flag = (fillColour: string) => (
    <svg
      width="10px"
      height="12px"
      viewBox="0 0 10 12"
      stroke-width="1"
      fill={fillColour}
      fill-rule="evenodd"
      stroke="#618FBA"
    >
      <path d="M1,7 L1,11" id="Line-5" fill="#D8D8D8" stroke-linecap="square" />
      <path
        d="M1,0.615322114 L1,6.71010785 C3.49882679,7.2372313 5.35688115,7.5 6.56064229,7.5 C7.64315653,7.5 8.45159281,7.28738766 9,6.87919186 L9,0.90645183 C8.31807098,1.23087118 7.50287034,1.39086885 6.56064229,1.39086885 C5.36422948,1.39086885 3.51568193,1.13193088 1,0.615322114 Z"
        id="Rectangle"
      />
    </svg>
  )
  state = { flagged: false }
  render() {
    return (
      <div className="notification">
        <div>
          <p className="createdAt">{this.time(this.props.change.createdAt)}</p>
          <p className="message">{this.props.change.message}</p>
          <div id="tick">{this.tick}</div>
          <div
            id="flag"
            onClick={() => {
              this.setState({ flagged: !this.state.flagged })
            }}
          >
            {this.flag(this.state.flagged ? '#618FBA' : 'none')}
          </div>
        </div>
        <div id="line" />
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
