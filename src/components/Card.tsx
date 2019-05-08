import React, { Component } from 'react'
import './components.scss'
import { LinkedCard } from '../store/cards/types'
import ReactMarkdown from 'react-markdown'

type Props = {
  onClick: () => void
  data: LinkedCard
  children?: never
}
export default class Card extends Component<Props> {
  render() {
    return (
      <div className="Card" onClick={this.props.onClick}>
        <div className="state" style={{ backgroundColor: 'green' }}>
          {this.props.data.item.state}
        </div>
        <div className="title">{this.props.data.item.title}</div>
        <div className="branch">{this.props.data.item.branch}</div>
        {this.desc()}
      </div>
    )
  }

  desc() {
    return this.props.data.item.desc ? (
      <p className="desc"> {this.props.data.item.desc}</p>
    ) : this.props.data.item.descHtml ? (
      <div
        className="desc"
        dangerouslySetInnerHTML={{ __html: this.props.data.item.descHtml }}
      />
    ) : (
      <ReactMarkdown
        className="desc"
        source={this.props.data.item.descMarkdown}
        escapeHtml={false}
      />
    )
  }
}
