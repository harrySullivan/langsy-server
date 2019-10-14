import createStore from 'unistore'

import { initialCounterState } from '$root/state/counter';

const initialState = Object.assign(
	{},
	initialCounterState
); 

export default createStore(initialState)