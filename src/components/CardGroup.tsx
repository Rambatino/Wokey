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
    console.log(this.props.data.linkedCards)
    return (
      <div className="CardGroup" onClick={this.props.onClick}>
        <div className="cards">
          <div className="content">
            <p className="title">{this.props.data.item.title}</p>
            <p className="subtitle">{this.props.data.item.subtitle}</p>
            {this.props.data.item.descHtml ? (
              <p
                className="desc"
                dangerouslySetInnerHTML={{
                  __html: this.props.data.item.descHtml,
                }}
              />
            ) : (
              <p className="desc"> {this.props.data.item.desc}</p>
            )}
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
