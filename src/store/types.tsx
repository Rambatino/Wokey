export interface Item {
  id: string
  desc?: string
  descHtml?: string
  title: string
  subtitle: string
  url: string
  state: string
}

export interface LinkedCard {
  item: Item
}

export interface BaseCard {
  item: Item
  linkedCards?: Array<LinkedCard>
}

interface Config {}

export interface Store {
  linkedCards: Array<LinkedCard>
  baseCards: Array<BaseCard>
  config?: Config
}
