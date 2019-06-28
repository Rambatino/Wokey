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
  grey = '#6a737d'
  purple = '#6f42c1'

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
    failure: {
      colour: this.red,
      path: (
        <path d="M7.48 8l3.75 3.75-1.48 1.48L6 9.48l-3.75 3.75-1.48-1.48L4.52 8 .77 4.25l1.48-1.48L6 6.52l3.75-3.75 1.48 1.48L7.48 8z" />
      ),
    },
    pending: {
      colour: this.orange,
      path: <path d="M0 8c0-2.2 1.8-4 4-4s4 1.8 4 4-1.8 4-4 4-4-1.8-4-4z" />,
    },
  }
  commentSymbol: { colour: string; path: JSX.Element | undefined } = {
    colour: '#586069',
    path: (
      <path d="M14 1H2c-.55 0-1 .45-1 1v8c0 .55.45 1 1 1h2v3.5L7.5 11H14c.55 0 1-.45 1-1V2c0-.55-.45-1-1-1zm0 9H7l-2 2v-2H2V2h12v8z" />
    ),
  }
  stateSymbol: { [key: string]: JSX.Element | undefined } = {
    APPROVED: (
      <div className="stateSymbol" style={{ backgroundColor: this.green }}>
        <svg className="approvedSVGCircle" width="12" height="16" fill="#FFF">
          <path d="M12 5l-8 8-4-4 1.5-1.5L4 10l6.5-6.5L12 5z" />
        </svg>
      </div>
    ),
    COMMENTED: (
      <div className="stateSymbol" style={{ backgroundColor: this.orange }}>
        <svg className="commentedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M0 1a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7.5L4 15.5V12H1a1 1 0 0 1-1-1V1zm1 0v10h4v2l2-2h8V1H1zm7.5" />
        </svg>
      </div>
    ),
    CHANGES_REQUESTED: (
      <div className="stateSymbol" style={{ backgroundColor: this.red }}>
        <svg className="rejectedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M0 1a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7.5L4 15.5V12H1a1 1 0 0 1-1-1V1zm1 0v10h4v2l2-2h8V1H1zm7.5 3h2v1h-2v2h-1V5h-2V4h2V2h1v2zm2 5h-5V8h5v1z" />
        </svg>
      </div>
    ),
    CONFLICTED: (
      <div className="stateSymbol" style={{ backgroundColor: this.grey }}>
        <svg className="conflictedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z" />
        </svg>
      </div>
    ),
    IS_MERGED: (
      <div className="stateSymbol" style={{ backgroundColor: this.purple }}>
        <svg className="mergedSVGCircle" width="16" height="16" fill="#FFF">
          <path d="M10 7c-.73 0-1.38.41-1.73 1.02V8C7.22 7.98 6 7.64 5.14 6.98c-.75-.58-1.5-1.61-1.89-2.44A1.993 1.993 0 0 0 2 .99C.89.99 0 1.89 0 3a2 2 0 0 0 1 1.72v6.56c-.59.35-1 .99-1 1.72 0 1.11.89 2 2 2a1.993 1.993 0 0 0 1-3.72V7.67c.67.7 1.44 1.27 2.3 1.69.86.42 2.03.63 2.97.64v-.02c.36.61 1 1.02 1.73 1.02 1.11 0 2-.89 2-2 0-1.11-.89-2-2-2zm-6.8 6c0 .66-.55 1.2-1.2 1.2-.65 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2zM2 4.2C1.34 4.2.8 3.65.8 3c0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2zm8 6c-.66 0-1.2-.55-1.2-1.2 0-.65.55-1.2 1.2-1.2.65 0 1.2.55 1.2 1.2 0 .65-.55 1.2-1.2 1.2z" />
        </svg>
      </div>
    ),
  }

  render() {
    return (
      <div className="Card" onClick={this.props.onClick}>
        {this.stateSymbol[this.props.data.approvalState]}
        <div className="title">
          {this.formatTitle(this.props.data.title, this.props.data.link)}
        </div>
        <div className="branch">
          {this.props.data.repo + ':' + this.props.data.branch}
        </div>
        {this.svg(
          (this.statusSymbol[this.props.data.ciStatus.status] || {}).colour,
          (this.statusSymbol[this.props.data.ciStatus.status] || {}).path,
          this.props.data.ciStatus.link
        )}
        {this.comment(
          this.props.data.comments.count,
          this.props.data.comments.lastCommentLink
        )}
      </div>
    )
  }

  formatTitle(title: string, url: string): string {
    const match = url.match(/(\d+)\b/) || []
    return '#' + (match.length > 0 ? match[0] : '') + ' ' + title
  }

  comment(count: number, link: string) {
    if (count === 0) {
      return null
    }
    return (
      <div>
        <svg
          className="commentSymbol"
          width="20"
          height="20"
          fill={this.commentSymbol.colour}
          onClick={() => window.open(link)}
        >
          {this.commentSymbol.path}
        </svg>
        <p
          className="commentCount"
          style={{ color: this.commentSymbol.colour }}
        >
          {count}
        </p>
      </div>
    )
  }

  svg(
    colour: string | undefined,
    path: JSX.Element | undefined,
    onClickUrl: string
  ) {
    return (
      <svg
        className="statusSymbol"
        width="20"
        height="20"
        fill={colour}
        onClick={() => window.open(onClickUrl)}
      >
        {path}
      </svg>
    )
  }
}
