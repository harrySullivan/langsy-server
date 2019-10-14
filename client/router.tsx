import { h } from 'preact'
import Router from 'preact-router'

import { Explore } from "$root/pages/Explore";
import { About } from "$root/pages/About";

export default () => (
  <Router>
    <Explore path="/" />
    <About path="/about" />
  </Router>
)