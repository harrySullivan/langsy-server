import createStore from 'unistore'

import { initialCounterState } from './state/counter';

const initialState = Object.assign(
	{},
	initialCounterState
); 

export default createStore(initialState)