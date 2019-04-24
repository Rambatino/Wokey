import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'
import './components.scss'
import { Item } from '../store/cards/types'

type Props = {
  onClick: () => void
  targetIds?: Array<string>
  archerId: string
  item: Item
  children?: never
}
export default class BaseCardItem extends Component<Props> {
  render() {
    return (
      <div className="BaseCard" onClick={this.props.onClick}>
        <ArcherElement
          id={this.props.archerId}
          relations={(this.props.targetIds || []).map(id => {
            return {
              targetId: id,
              targetAnchor: 'left',
              sourceAnchor: 'right',
              style: { strokeColor: '#6E6E6E', strokeWidth: 1 },
            }
          })}
        >
          <div className="card">
            <div className="card-body">
              <p className="card-title">{this.props.item.title}</p>
              <p className="card-key">{this.props.item.subtitle}</p>
              {this.props.item.desc && (
                <div className="card-text">{this.props.item.desc}</div>
              )}
              {this.props.item.descHtml && (
                <div
                  className="card-text"
                  dangerouslySetInnerHTML={{ __html: this.props.item.descHtml }}
                />
              )}
            </div>
          </div>
        </ArcherElement>
      </div>
    )
  }
}
