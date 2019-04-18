import 'bootstrap/dist/css/bootstrap.css';

import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import { linkTo } from '@storybook/addon-links';

import { Button, Welcome } from '@storybook/react/demo';

import LinkedCard from '../components/LinkedCard'
import BaseCard from '../components/BaseCard'

storiesOf('Welcome', module).add('to Storybook', () => <Welcome showApp={linkTo('Button')} />);

storiesOf('Workspace', module)
  .add('can handle cards with movement', () => (
      <BaseCard>
        <div>
          <LinkedCard />
          <LinkedCard />
          <LinkedCard />
        </div>
      </BaseCard>
  ));
