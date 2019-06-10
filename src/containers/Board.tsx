import React, { Component } from 'react'
import CardGroup from '../components/CardGroup'
import Card from '../components/Card'
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
    console.log(this.props)
    const allStates = this.getAllStates()
    return (
      <div className="Board">
        {allStates.map((state, i) => (
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
        {this.props.data.linkedCards && (
          <div className="boardRow">
            <div
              className="statusTitle"
              style={{
                backgroundColor: allStates.length % 2 ? '#F8F8F8' : '#FFF',
              }}
            >
              <p> Pull Requests </p>
            </div>
            <div className="verticalLine" />
            <div
              className={'pullRequests'}
              style={{
                width: '100%',
                paddingLeft: '7px',
                backgroundColor: allStates.length % 2 ? '#F8F8F8' : '#FFF',
              }}
            >
              {this.props.data.linkedCards.map((card, i) => (
                <Card
                  key={i}
                  data={card}
                  onClick={() => window.open(card.item.url)}
                />
              ))}
            </div>
          </div>
        )}
      </div>
    )
  }

  getAllStates() {
    return [...new Set(this.props.data.baseCards.map(c => c.item.state))]
  }
}
