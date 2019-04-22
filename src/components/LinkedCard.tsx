import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'
import './components.scss'

export default class LinkedCard extends Component {
  render() {
    return (
      <ArcherElement id={this.props.archerId}>
        <div className="LinkedCard">
          <div className="card">
            <div className="card-body">
              <p className="card-title">{this.props.title}</p>
              <p className="card-subtitle">{this.props.subtitle}</p>
              <div className="card-text">{this.props.desc}</div>
            </div>
          </div>
        </div>
      </ArcherElement>
    )
  }
}
