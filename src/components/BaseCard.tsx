import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'

export default class BaseCard extends Component {
  render() {
    return (
      <div className="BaseCard" style={{ marginBottom: 20 }}>
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
          <div className="card" style={{ width: '20rem' }}>
            <div className="card-body">
              <h5 className="card-title">Base Card</h5>
              <p className="card-text">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
                eiusmod tempor incididunt ut labore et dolore magna aliqua
              </p>
            </div>
          </div>
        </ArcherElement>
      </div>
    )
  }
}
