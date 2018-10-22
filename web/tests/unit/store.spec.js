import { createLocalVue } from '@vue/test-utils'
import Vuex from 'vuex'
import store from '@/store'

describe('Vuex store', () => {
  beforeEach(() => {
    const localVue = createLocalVue()
    localVue.use(Vuex)
  })
  it('should update current number of available coupons when commited', () => {
    expect(store.state.available).toBe(0)
    store.commit('SetAvailable', { count: 100 })
    expect(store.state.available).toBe(100)
  })
})