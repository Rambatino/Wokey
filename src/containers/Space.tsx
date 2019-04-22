import React, { Component } from 'react'
import LinkedCard from '../components/LinkedCard'
import BaseCard from '../components/BaseCard'
import './containers.css'

import { ArcherContainer } from 'react-archer'

import { BaseCardItem, LinkedCardItem } from '../data/types'

type Props = {
  data: {
    linkedCards: Array<LinkedCardItem>
    baseCards: Array<BaseCardItem>
  }
}
export default class Space extends Component<Props> {
  state = {
    currentSelectionIdx: 0,
  }

  render() {
    return (
      <div>
        <ArcherContainer strokeColor="red">
          <div style={{ display: 'flex' }}>
            <div
              className="baseSpace"
              style={{
                marginTop: 20,
                marginLeft: 20,
                marginRight: 15,
              }}
            >
              {this.props.data.baseCards.map((card, i) => (
                <BaseCard
                  key={'base-' + card.id}
                  archerId={'root-' + card.id}
                  targetIds={
                    i === this.state.currentSelectionIdx
                      ? card.linkedCards.map(c => c.id)
                      : undefined
                  }
                  title={card.title}
                  desc={card.desc}
                  subtitle={card.subtitle}
                  onClick={() => this.setState({ currentSelectionIdx: i })}
                />
              ))}
            </div>
            <div
              className="linkedSpace"
              style={{ flex: 1, marginTop: 20, marginLeft: 15 }}
            >
              {this.props.data.baseCards[
                this.state.currentSelectionIdx
              ].linkedCards.map(c => (
                <LinkedCard
                  key={'linked-' + c.id}
                  archerId={c.id}
                  title={c.title}
                  subtitle={c.subtitle}
                  desc={c.desc}
                />
              ))}
            </div>
          </div>
        </ArcherContainer>
      </div>
    )
  }
}
