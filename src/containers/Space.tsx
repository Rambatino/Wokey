import React, { Component } from 'react'
import LinkedCardItem from '../components/LinkedCardItem'
import BaseCardItem from '../components/BaseCardItem'
import './containers.scss'

import { ArcherContainer } from 'react-archer'

import { BaseCard, LinkedCard } from '../store/types'

type Props = {
  data: {
    linkedCards: Array<LinkedCard>
    baseCards: Array<BaseCard>
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
                <BaseCardItem
                  key={'base-' + card.item.id}
                  archerId={'root-' + card.item.id}
                  targetIds={
                    i === this.state.currentSelectionIdx
                      ? card.linkedCards && card.linkedCards.map(c => c.item.id)
                      : undefined
                  }
                  title={card.item.title}
                  desc={card.item.desc_html}
                  subtitle={card.item.subtitle}
                  onClick={() => this.setState({ currentSelectionIdx: i })}
                />
              ))}
            </div>
            <div className="linkedSpace">{this.linkedCards()}</div>
          </div>
        </ArcherContainer>
      </div>
    )
  }

  linkedCards = () => {
    return (
      this.props.data.baseCards.length > this.state.currentSelectionIdx &&
      this.props.data.baseCards[this.state.currentSelectionIdx].linkedCards &&
      this.props.data.baseCards[this.state.currentSelectionIdx].linkedCards.map(
        c => (
          <LinkedCardItem
            key={'linked-' + c.item.id}
            archerId={c.item.id}
            title={c.item.title}
            subtitle={c.item.subtitle}
            desc={c.item.desc}
          />
        )
      )
    )
  }
}
