import React from 'react'
import { storiesOf } from '@storybook/react'

import Space from '../containers/Space'
import spaceData from './space'
import BaseCardItem from '../components/BaseCardItem'
import LinkedCardItem from '../components/LinkedCardItem'
import { ArcherContainer } from 'react-archer'

storiesOf('Workspace', module)
  .add('Base Card', () => (
    <ArcherContainer>
      <BaseCardItem
        archerId=""
        onClick={() => {}}
        item={{
          id: '1',
          state: 'Ready for Code Review',
          url: 'https://zigroup.atlassian.net/browse/STATS-302',
          title: 'CTA incorrect shows question text instead of measure name',
          descHtml:
            '\u003cp\u003eCTA now shows question text instead of measure names.\u003c/p\u003e\n\n\u003cp\u003e \u003cspan class="error"\u003eUnable to render embedded object: File (image-2019-02-01-15-18-49-806.png) not found.\u003c/span\u003e \u003c/p\u003e\n\n\u003cp\u003e\u003ca href="https://www.zappistore.com/app/analysis/video_creative_test/explore/44917,44913,44894,44897,44898/stats/view/Content%20Theme%20Analysis/pages/6377587" class="external-link" rel="nofollow"\u003ehttps:/,/www.zappistore.com/app/analysis/video_creative_test/explore/44917,44913,44894,44897,44898/stats/view/Content%20Theme%20Analysis/pages/6377587\u003c/a\u003e\u003c/p\u003e\n\n\u003cp\u003eUAC:\u003c/p\u003e\n\n\u003cp\u003eCTA shows measure names instead of question text. \u003c/p\u003e\n',
          subtitle: 'STATS-302',
        }}
      />
    </ArcherContainer>
  ))
  .add('Link Card', () => (
    <ArcherContainer>
      <LinkedCardItem
        archerId=""
        item={{
          id: '1',
          title: 'Link Card',
          descHtml:
            '<p> Ambitioni dedisse scripsisse iudicaretur. Cras mattis iudicium purus sit amet fermentum. Donec sed odio operae, eu vulputate felis rhoncus. Praeterea iter est quasdam res quas ex communi. At nos hinc posthac, sitientis piros Afros. Petierunt uti sibi concilium totius Galliae in diem certam indicere. Cras mattis iudicium purus sit amet fermentum. Ambitioni dedisse scripsisse iudicaretur. Cras mattis iudicium purus sit amet fermentum. Donec sed odio operae, eu vulputate felis rhoncus. Praeterea iter est quasdam res quas ex communi. At nos hinc posthac, sitientis piros Afros. Petierunt uti sibi concilium totius Galliae in diem certam indicere. Cras mattis iudicium purus sit amet fermentum.</p>',
          subtitle: 'Subtitle 1',
          url: 'www.lolcat.com',
          state: 'APPROVED',
        }}
      />
      <LinkedCardItem
        archerId=""
        item={{
          id: '1',
          title: 'Link Card Short Desc',
          descHtml:
            '<p>Ambitioni dedisse scripsisse iudicaretur. Cras mattis iudicium purus</p>',
          subtitle: '',
          state: '',
          url: '',
        }}
      />
    </ArcherContainer>
  ))
  .add('Space', () => <Space data={spaceData} />)
