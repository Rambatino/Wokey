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
        <div className="content">
          <div className="titleAndSubtitle">
            <p className="subtitle">{this.props.data.item.subtitle}</p>
            <p className="title">{this.props.data.item.title}</p>
          </div>
          {this.props.data.linkedCards && (
            <div
              style={{
                width: '0.7px',
                marginLeft: 10,
                marginTop: 10,
                marginBottom: 10,
                height: this.props.data.linkedCards.length * 65,
                backgroundColor: '#D8D8D8',
              }}
            />
          )}
          <div className="cards">
            {this.props.data.linkedCards.map((card, i) => (
              <Card
                key={i}
                data={card}
                onClick={() => window.open(card.item.url)}
              />
            ))}
          </div>
        </div>
      </div>
    )
  }
}
