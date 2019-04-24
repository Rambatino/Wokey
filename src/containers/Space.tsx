import React, { Component } from 'react'
import LinkedCardItem from '../components/LinkedCardItem'
import BaseCardItem from '../components/BaseCardItem'
import './containers.scss'

import { ArcherContainer, ArcherContainerProps } from 'react-archer'

import { BaseCard, LinkedCard } from '../store/types'

type Props = {
  data: {
    linkedCards: Array<LinkedCard>
    baseCards: Array<BaseCard>
  }
}
type State = {
  currentSelectionIdx: number
}
export default class Space extends Component<Props, State> {
  state = {
    currentSelectionIdx: 0,
  }
  baseSpaceRef = React.createRef<HTMLDivElement>()
  linkSpaceRef = React.createRef<HTMLDivElement>()
  archerContainerRef = React.createRef<ArcherContainerProps>()

  componentDidMount = () => {
    this.baseSpaceRef.current &&
      this.baseSpaceRef.current.addEventListener('scroll', this.isScrolling)
    this.linkSpaceRef.current &&
      this.linkSpaceRef.current.addEventListener('scroll', this.isScrolling)
  }

  componentWillUnmount = () => {
    this.baseSpaceRef.current &&
      this.baseSpaceRef.current.removeEventListener('scroll', this.isScrolling)
    this.linkSpaceRef.current &&
      this.linkSpaceRef.current.removeEventListener('scroll', this.isScrolling)
  }

  isScrolling = () => {
    this.archerContainerRef!.current!.forceUpdate()
  }

  render() {
    return (
      <div>
        <ArcherContainer ref={this.archerContainerRef}>
          <div className="space">
            <div className="baseSpace" ref={this.baseSpaceRef}>
              {this.baseCards()}
            </div>
            <div className="linkedSpace" ref={this.linkSpaceRef}>
              {this.linkedCards()}
            </div>
          </div>
        </ArcherContainer>
      </div>
    )
  }

  baseCards = () => {
    return this.props.data.baseCards.map((card, i) => (
      <BaseCardItem
        key={'base-' + card.item.id}
        archerId={'root-' + card.item.id}
        targetIds={
          i === this.state.currentSelectionIdx
            ? card.linkedCards && card.linkedCards.map(c => c.item.id)
            : undefined
        }
        item={card.item}
        onClick={() => this.setState({ currentSelectionIdx: i })}
      />
    ))
  }

  linkedCards = () => {
    return (
      this.props.data.baseCards[this.state.currentSelectionIdx] &&
      this.props.data.baseCards[this.state.currentSelectionIdx].linkedCards &&
      this.props.data.baseCards[
        this.state.currentSelectionIdx
      ].linkedCards!.map(c => (
        <LinkedCardItem
          key={'linked-' + c.item.id}
          archerId={c.item.id}
          item={c.item}
        />
      ))
    )
  }
}
