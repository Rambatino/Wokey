import { LinkedCard, Cards, BaseCard } from './types'

export const filterCards = (
  baseCards: Array<BaseCard>,
  linkedCards: Array<LinkedCard>
): Cards => {
  if (baseCards.length * linkedCards.length === 0) {
    return {
      baseCards: baseCards,
      linkedCards: linkedCards,
    }
  }
  // if can match them by branch name or by desc or title then add to
  // base card linked cards and remove from linkedCards set
  let found: Array<boolean> = linkedCards.map(() => false)
  return {
    baseCards: baseCards.map(baseCard => {
      const key = baseCard.item.subtitle
      let linkedStore: Array<LinkedCard> = []
      linkedCards.forEach((linkedCard, i) => {
        const { descMarkdown, title, branch } = linkedCard.item
        if (
          (descMarkdown && descMarkdown!.includes(key)) ||
          title.includes(key) ||
          (branch && branch!.includes(key))
        ) {
          found[i] = true
          linkedStore = [...linkedStore, linkedCard]
        }
      })
      baseCard.linkedCards = linkedStore
      return baseCard
    }),
    linkedCards: linkedCards.filter((_, i) => !found[i]),
  }
}
