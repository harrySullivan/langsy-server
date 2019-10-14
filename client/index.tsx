import { h, render } from 'preact'
import { Provider } from 'unistore/preact'
import Router from '$root/router'

import "bulma/css/bulma.min.css"
import "@fortawesome/fontawesome-free/css/all.min.css"
import "spacing/main"

import store from "./store"

const app = document.getElementById('root')

render(
  <Provider store={store}>
    <Router />
  </Provider>,
  app
)