import React, { Component } from 'react'
import './components.scss'
import { LinkedCard } from '../store/cards/types'

type Props = {
  onClick: () => void
  data: LinkedCard
  children?: never
}
export default class Card extends Component<Props> {
  green = '#28a745'
  orange = '#ffac45'
  red = '#cb2431'

  statusSymbol: {
    [key: string]: {
      colour: string
      path: JSX.Element | undefined
    }
  } = {
    success: {
      colour: this.green,
      path: <path d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5L12 5z" />,
    },
    failed: {
      colour: this.red,
      path: (
        <path d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z" />
      ),
    },
    running: {
      colour: this.orange,
      path: <path d="M0 8c0-2.2 1.8-4 4-4s4 1.8 4 4-1.8 4-4 4-4-1.8-4-4z" />,
    },
  }
  stateSymbol: { [key: string]: JSX.Element | undefined } = {
    approved: (
      <div className="stateSymbol" style={{ backgroundColor: this.green }}>
        <svg className="approvedSVGCircle" width="12" height="16" fill="#FFF">
          <path d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5L12 5z" />
        </svg>
      </div>
    ),
    commented: (
      <div className="stateSymbol" style={{ backgroundColor: this.orange }}>
        <svg className="rejectedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M0 1a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7.5L4 15.5V12H1a1 1 0 0 1-1-1V1zm1 0v10h4v2l2-2h8V1H1zm7.5" />
        </svg>
      </div>
    ),
    changesRequested: (
      <div className="stateSymbol" style={{ backgroundColor: this.red }}>
        <svg className="rejectedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M0 1a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7.5L4 15.5V12H1a1 1 0 0 1-1-1V1zm1 0v10h4v2l2-2h8V1H1zm7.5 3h2v1h-2v2h-1V5h-2V4h2V2h1v2zm2 5h-5V8h5v1z" />
        </svg>
      </div>
    ),
  }

  render() {
    return (
      <div className="Card" onClick={this.props.onClick}>
        {this.stateSymbol[this.props.data.item.state]}
        <div className="title">
          {this.formatTitle(
            this.props.data.item.title,
            this.props.data.item.url
          )}
        </div>
        <div className="branch">{this.props.data.item.repo}</div>
        {this.svg(
          (this.statusSymbol[this.props.data.item.status!] || {}).colour,
          (this.statusSymbol[this.props.data.item.status!] || {}).path
        )}
      </div>
    )
  }

  formatTitle(title: string, url: string): string {
    const match = url.match(/(\d+)\b/) || []
    return '#' + (match.length > 0 ? match[0] : '') + ' ' + title
  }

  svg(colour: string | undefined, path: JSX.Element | undefined) {
    return (
      <svg className="statusSymbol" width="20" height="20" fill={colour}>
        {path}
      </svg>
    )
  }
}
