import React, { Component } from 'react'
import CardGroup from '../components/CardGroup'
// import BaseCardItem from '../components/BaseCardItem'
import './containers.scss'

import { BaseCard, LinkedCard } from '../store/cards/types'

type Props = {
  data: {
    linkedCards: Array<LinkedCard>
    baseCards: Array<BaseCard>
  }
}

// Board will organise the flow of the board, how
// all the cards are layed out
export default class Board extends Component<Props> {
  render() {
    return (
      <div className="Board">
        {this.props.data.baseCards.map((card, i) => (
          <CardGroup key={i} data={card} onClick={() => window.open(card.item.url)} />
        ))}
      </div>
    )
  }
}
