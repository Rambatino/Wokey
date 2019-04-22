import React from 'react'

import { storiesOf } from '@storybook/react'
import { action } from '@storybook/addon-actions'
import { linkTo } from '@storybook/addon-links'

import { Button, Welcome } from '@storybook/react/demo'

import Space from '../containers/Space'
import spaceData from './space'
import BaseCard from '../components/BaseCard'
import LinkedCard from '../components/LinkedCard'
import { ArcherContainer } from 'react-archer'

storiesOf('Workspace', module)
  .add('Base Card', () => (
    <ArcherContainer>
      <BaseCard title="title" desc="desc" subtitle="subtitle" />
    </ArcherContainer>
  ))
  .add('Link Card', () => (
    <ArcherContainer>
      <LinkedCard
        title="Link Card"
        desc={
          <p>
            Ambitioni dedisse scripsisse iudicaretur. Cras mattis iudicium purus
            sit amet fermentum. Donec sed odio operae, eu vulputate felis
            rhoncus. Praeterea iter est quasdam res quas ex communi. At nos hinc
            posthac, sitientis piros Afros. Petierunt uti sibi concilium totius
            Galliae in diem certam indicere. Cras mattis iudicium purus sit amet
            fermentum. Ambitioni dedisse scripsisse iudicaretur. Cras mattis
            iudicium purus sit amet fermentum. Donec sed odio operae, eu
            vulputate felis rhoncus. Praeterea iter est quasdam res quas ex
            communi. At nos hinc posthac, sitientis piros Afros. Petierunt uti
            sibi concilium totius Galliae in diem certam indicere. Cras mattis
            iudicium purus sit amet fermentum.
          </p>
        }
        subtitle="Subtitle 1"
      />
      <LinkedCard
        title="Link Card Short Desc"
        desc={
          <p>
            Ambitioni dedisse scripsisse iudicaretur. Cras mattis iudicium purus
          </p>
        }
        subtitle="Subtitle 2"
      />
    </ArcherContainer>
  ))
  .add('Space', () => <Space data={spaceData} />)
