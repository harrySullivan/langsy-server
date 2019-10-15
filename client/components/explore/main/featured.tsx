import { h } from 'preact'
import "./featured.css"
import { SectionHeading } from './section-heading'

const Item = (
	(primaryInfo: string, secondaryInfo: string, thumbnail: string) => (
		<div class="column is-4">
			<img class="featured-thumbnail" src={thumbnail} />
			<h5 class="has-text-weight-bold ml1 mt1">{primaryInfo}</h5>
			<h6 class="has-text-grey-light ml1">{secondaryInfo}</h6>
		</div>
	)
)

export const Featured = (
  () => (
  	<div class="column">
	    {SectionHeading("featured", "albums")}
	    <div class="columns pl2 results-wrapper">
	    	{Item("Blue Moon", "Hayke", "https://www.theaudiodb.com/images/media/artist/thumb/uxrqxy1347913147.jpg")}
	    	{Item("Blue Moon", "Hayke", "https://www.theaudiodb.com/images/media/artist/thumb/wtxxru1527589862.jpg")}
	    	{Item("Blue Moon", "Hayke", "https://www.theaudiodb.com/images/media/artist/thumb/uxrqxy1347913147.jpg")}
	    	{Item("Blue Moon", "Hayke", "https://www.theaudiodb.com/images/media/artist/thumb/wtxxru1527589862.jpg")}
	    </div>
    </div>
  )
)