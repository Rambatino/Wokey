import React, { Component } from 'react'
import LinkedCard from '../components/LinkedCard'
import BaseCard from '../components/BaseCard'
import './containers.scss'

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
          <div className="space">
            <div className="baseSpace">
              {this.props.data.baseCards.map((card, i) => (
                <BaseCard
                  key={'base-' + card.id}
                  archerId={'root-' + card.id}
                  targetIds={
                    i === this.state.currentSelectionIdx
                      ? card.linkedCards && card.linkedCards.map(c => c.id)
                      : undefined
                  }
                  title={card.title}
                  desc={card.desc_html}
                  subtitle={card.subtitle}
                  onClick={() => this.setState({ currentSelectionIdx: i })}
                />
              ))}
            </div>
            <div className="linkedSpace">
              {this.props.data.baseCards.length >
                this.state.currentSelectionIdx &&
                this.props.data.baseCards[this.state.currentSelectionIdx]
                  .linkedCards &&
                this.props.data.baseCards[
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
