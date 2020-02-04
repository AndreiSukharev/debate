import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loggedUser: {}
  },
  mutations: {
    loginUser(state, user) {
      state.loggedUser = user
    },
    logoutUser(state) {
      state.loggedUser = {}
    },
  },
  actions: {
  },
  modules: {
  }
})
