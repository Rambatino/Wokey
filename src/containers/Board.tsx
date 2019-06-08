import React, { Component } from 'react'
import CardGroup from '../components/CardGroup'
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
        {this.getAllStates().map((state, i) => (
          <div className="boardRow">
            <div
              className="statusTitle"
              style={{ backgroundColor: i % 2 ? '#F8F8F8' : '#FFF' }}
            >
              <p> {state} </p>
            </div>
            <div className="verticalLine" />
            <div
              className={state}
              key={state}
              style={{
                width: '100%',
                paddingLeft: '7px',
                backgroundColor: i % 2 ? '#F8F8F8' : '#FFF',
              }}
            >
              {this.props.data.baseCards
                .filter(card => card.item.state === state)
                .map((card, i) => (
                  <CardGroup
                    key={i}
                    data={card}
                    onClick={() => window.open(card.item.url)}
                  />
                ))}
            </div>
          </div>
        ))}
      </div>
    )
  }

  getAllStates() {
    return [...new Set(this.props.data.baseCards.map(c => c.item.state))]
  }
}
