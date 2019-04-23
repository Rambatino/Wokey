export interface LinkedCardItem {
  id: string
  desc: string
  desc_html: string
  title: string
  subtitle: string
  url: string
  state: string
}

export interface BaseCardItem {
  id: string
  desc: string
  desc_html: string
  title: string
  subtitle: string
  url: string
  state: string
  linkedCards: Array<LinkedCardItem>
}
