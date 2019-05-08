import React from 'react'
import { storiesOf } from '@storybook/react'

import Board from '../containers/Board'
import issues from './issues.json'
import pulls from './pulls.json'
import { ArcherContainer } from 'react-archer'
import CardGroup from '../components/CardGroup'
import { filterCards } from '../store/cards/formatter'
import { BaseCard, LinkedCard } from '../store/cards/types'
import Card from '../components/Card'

const spaceData = filterCards(
  issues.map(i => ({ item: i } as BaseCard)),
  pulls.map(i => ({ item: i } as LinkedCard))
)
storiesOf('New Workspace', module)
  .add('Card', () => (
    <Card data={spaceData.baseCards[0].linkedCards[0]} onClick={() => {}} />
  ))
  .add('Card Group Empty', () => (
    <CardGroup data={spaceData.baseCards[5]} onClick={() => {}} />
  ))
  .add('Card Group', () => (
    <CardGroup data={spaceData.baseCards[0]} onClick={() => {}} />
  ))
  .add('Board', () => <Board data={spaceData} />)
