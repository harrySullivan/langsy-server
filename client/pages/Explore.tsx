import { h } from 'preact'
import { Link } from 'preact-router'
import { connect } from 'unistore/preact'

import { actions } from '$root/state/counter'

import { Sidebar } from '$root/components/explore/sidebar/sidebar'
import { Main } from '$root/components/explore/main/main'

export const Explore = connect('count', actions)(
  ({ count, increment, decrement }: any) => (
    <div class="columns pt2">
    	<div class="column is-one-fifth">{Sidebar()}</div>
    	<div class="column">
    		{Main()}
  		</div>
    	<div class="column is-one-fifth">play queue</div>
    </div>
  )
)