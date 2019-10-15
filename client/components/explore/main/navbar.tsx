import { h } from 'preact'
import { Link } from "preact-router/match";
import logo from "$root/img/cropped-and-shrunk.png"

export const Navbar = (
  () => (
  	<nav class="level">
			<div class="level-left">
				<Link class="navbar-item" href="/">
		      <img src={logo} />
		      <h5 class="title is-5 has-text-weight-bold is-uppercase ml1">musicsy</h5>
		    </Link>
				<input class="input is-rounded is-medium ml2" placeholder="Search..." />
			</div>
			<div class="level-right">
				<button class="button is-white is-large">
					John Doe
					<img class="ml2" src={logo} />
				</button>
			</div>
  	</nav>
  )
)