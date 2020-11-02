import Vue from "vue";
import Vuex from "vuex";
import gofi from "./modules/gofi.js";
Vue.use(Vuex);

export default new Vuex.Store({
  namespaced: true,
  state: {},
  mutations: {},
  actions: {},
  modules: {
    gofi
  }
});
