import React from 'react'
import { storiesOf } from '@storybook/react'

import Board from '../containers/Board'
import state from './state.json'
import CardGroup from '../components/CardGroup'
import Card from '../components/Card'
import { parseJSON } from '../store/cards/reducer'
import Banner from '../components/Banner'
import Notifications from '../components/Notifications'

const stateData = parseJSON(state)

storiesOf('New Workstate', module)
  .add('Notifications empty', () => <Notifications changes={[]} />)
  .add('Notifications', () => <Notifications changes={stateData.changes} />)
  .add('Banner OPEN', () => <Banner state="REDUX_WEBSOCKET::OPEN" />)
  .add('Banner CLOSED', () => <Banner state="REDUX_WEBSOCKET::CLOSED" />)
  .add('Banner MESSAGE', () => <Banner state="REDUX_WEBSOCKET::MESSAGE" />)
  .add('Banner CONNECT', () => <Banner state="REDUX_WEBSOCKET::CONNECT" />)
  .add('Card', () => (
    <Card data={stateData.baseCards[0].linkedCards[0]} onClick={() => {}} />
  ))
  .add('Card Group Empty', () => (
    <CardGroup data={stateData.baseCards[2]} onClick={() => {}} />
  ))
  .add('Card Group', () => (
    <CardGroup data={stateData.baseCards[0]} onClick={() => {}} />
  ))
  .add('Board', () => <Board data={stateData} />)
