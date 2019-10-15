import { h } from 'preact'

export const SidebarItem = (
  (description: string) => (
  	<div>
	    <a class="button is-white">
	    	{description}
	    </a>
    </div>
  )
)