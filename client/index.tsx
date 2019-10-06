import { h, render } from 'preact'
import { Provider } from 'unistore/preact'
import Router from './router'

import "bulma/css/bulma.min.css"
import "@fortawesome/fontawesome-free/css/all.min.css"
import "spacing/main"

import createStore from 'unistore'

const store = createStore({ count: 55 })

const app = document.getElementById('root')

render(
  <Provider store={store}>
    <Router />
  </Provider>,
  app
)