import React, { Component } from 'react'
import LinkedCard from './LinkedCard'
import BaseCard from './BaseCard'

import { ArcherContainer } from 'react-archer'

export default class Space extends Component {
  render() {
    return (
      <div>
        <ArcherContainer strokeColor="red">
          <div style={{ display: 'flex' }}>
            <div
              className="baseSpace"
              style={{
                marginTop: 20,
                marginLeft: 20,
                marginRight: 15,
              }}
            >
              <BaseCard />
              <BaseCard archerId="root" targetIds={['a', 'b', 'c']} />
              <BaseCard />
            </div>
            <div
              className="linkedSpace card-deck"
              style={{ flex: 1, marginTop: 20, marginLeft: 15 }}
            >
              <LinkedCard archerId="a" />
              <LinkedCard archerId="b" />
              <LinkedCard archerId="c" />
            </div>
          </div>
        </ArcherContainer>
      </div>
    )
  }
}
