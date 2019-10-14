import { h } from 'preact'
import { Link } from 'preact-router'
import { connect } from 'unistore/preact'

import { actions } from '$root/state/counter'

import { Sidebar } from '$root/components/explore/explore.sidebar'

export const Explore = connect('count', actions)(
  ({ count, increment, decrement }: any) => (
    <div class="columns">
    	<div class="column is-one-quarter">{Sidebar()}</div>
    	<div class="column">
    		main {count}
    		<button onclick={increment}>+</button>
    		<button onclick={decrement}>-</button>
  		</div>
    	<div class="column is-one-quarter">play queue</div>
    </div>
  )
)