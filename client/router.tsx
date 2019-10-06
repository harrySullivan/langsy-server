import { h } from 'preact'
import Router from 'preact-router'

import { Explore } from "./pages/Explore";
import { About } from "./pages/About";

export default () => (
  <Router>
    <Explore path="/" />
    <About path="/about" />
  </Router>
)