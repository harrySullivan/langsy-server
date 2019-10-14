export interface State {
  count: number;
}

export let actions = () => ({
  increment(state: State) {
    return { count: state.count + 1 }
  },
  decrement(state: State) {
    return { count: state.count - 1 }
  }
})

export const initialCounterState: State = {
	count: 0
}

