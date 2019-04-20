import React, { Component } from 'react'
import { ArcherElement } from 'react-archer'

export default class LinkedCard extends Component {
  render() {
    return (
      <div className="LinkedCard" style={{ marginBottom: 30 }}>
        <div
          className="card"
          style={{
            width: '20rem',
            overflow: 'hidden',
            maxHeight: 300,
          }}
        >
          <ArcherElement id={this.props.archerId}>
            <div className="card-body">
              <h5 className="card-title">Link Card</h5>
              <p className="card-text">
                Lorem ipsum Lorem ipsumLorem ipsumLorem iLorem ipsumLorem
                ipsumLorem ipsumLorem ipsumLorem ipsumpsumLorem ipsumLorem
                ipsumLorem ipsumLorem ipsumLorem ipsumdolor sit amet,
                consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                labore et dolore magna aliqua. Lorem ipsum dolor sit amet,
                consectetur adipiscing elit, sed do eiusmod tempor incididunt ut
                labore et dolore magna aliqua
              </p>
            </div>
          </ArcherElement>
        </div>
      </div>
    )
  }
}
