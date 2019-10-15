import { h } from 'preact'
import { Navbar } from './navbar'
import { Featured } from './featured'

export const Main = (
  () => (
  	<div class="column">
	    {Navbar()}
	    {Featured()}
    </div>
  )
)