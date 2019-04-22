import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'
import './components.scss'

export default class BaseCard extends Component {
  render() {
    return (
      <div
        className="BaseCard"
        style={{ marginBottom: 20 }}
        onClick={this.props.onClick}
      >
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
              <h5 className="card-title">{this.props.title}</h5>
              <p className="card-key">{this.props.subtitle}</p>
              <div className="card-text">{this.props.desc}</div>
            </div>
          </div>
        </ArcherElement>
      </div>
    )
  }
}
