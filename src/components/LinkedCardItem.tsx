import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'
import './components.scss'

type Props = {
  archerId: string
  title: string
  subtitle: string
  desc: string
  children?: never
}

export default class LinkedCardItem extends Component<Props> {
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
