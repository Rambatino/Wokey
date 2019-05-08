import React, { Component } from 'react'
import './components.scss'
import { BaseCard } from '../store/cards/types'
import Card from './Card'

type Props = {
  onClick: () => void
  data: BaseCard
  children?: never
}

export default class CardGroup extends Component<Props> {
  render() {
    return (
      <div className="CardGroup" onClick={this.props.onClick}>
        {this.props.data.linkedCards.length === 0 ? (
          <div className="contentTitle" style={{ width: 200 }}>
            {this.props.data.item.title}
          </div>
        ) : (
          <div className="contentTitle">{this.props.data.item.title}</div>
        )}
        <hr />
        <div className="cards">
          <div className="content">
            <p className="subtitle">{this.props.data.item.subtitle}</p>
            <div className="state">{this.props.data.item.state}</div>
          </div>
          {this.props.data.linkedCards.map((card, i) => (
            <Card key={i} data={card} onClick={() => {}} />
          ))}
        </div>
      </div>
    )
  }
}
