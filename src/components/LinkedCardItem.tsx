import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'
import './components.scss'
import { Item } from '../store/cards/types'

type Props = {
  archerId: string
  item: Item
  children?: never
}

export default class LinkedCardItem extends Component<Props> {
  render() {
    return (
      <ArcherElement id={this.props.archerId}>
        <div className="LinkedCard">
          <div className="card">
            <div className="card-body">
              <p className="card-title">{this.props.item.title}</p>
              <p className="card-subtitle">{this.props.item.subtitle}</p>
              <div className="card-text">{this.props.item.desc}</div>
            </div>
          </div>
        </div>
      </ArcherElement>
    )
  }
}
