import { h } from 'preact'

export const SectionHeading = (
  (title: string, description: string) => (
  	<div class="column">
	    <h5 class="has-text-weight-bold has-text-grey-light mb0 is-capitalized">{title}</h5>
	    <h4 class="title is-4 is-capitalized">{description}</h4>
    </div>
  )
)