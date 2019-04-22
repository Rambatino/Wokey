export interface LinkedCardItem {
  id: string
  title: string
  desc: string
  subtitle: string
}

export interface BaseCardItem {
  id: string
  title: string
  desc: string
  subtitle: string
  linkedCards: Array<LinkedCardItem>
}
