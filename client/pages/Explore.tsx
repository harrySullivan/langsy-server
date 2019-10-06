import { h } from 'preact'
import { Link } from 'preact-router'
import { connect } from 'unistore/preact'

import { actions } from '../state/counter'

import { Sidebar } from '../components/explore/explore.sidebar'

export const Explore = connect('count', actions)(
  ({ count, increment, decrement }: any) => (
    <div class="columns">
    	<div class="column is-one-quarter">{Sidebar()}</div>
    	<div class="column">main</div>
    	<div class="column is-one-quarter">play queue</div>
    </div>
  )
)