import { h } from 'preact'

export const LensItem = (
  (description: string, icon: string) => (
  	<div>
	    <a class="button is-white mt1">
	    	<i class={`fas fa-${icon} mr1`}></i>
	    	{description}
	    </a>
    </div>
  )
)