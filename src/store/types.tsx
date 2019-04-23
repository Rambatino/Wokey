import { Item } from '../actions/actionTypes'

export interface LinkedCard {
  item: Item
}

export interface BaseCard {
  item: Item
  linkedCards: Array<LinkedCard>
}

interface Config {}

export interface Store {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
  config: Config
}
