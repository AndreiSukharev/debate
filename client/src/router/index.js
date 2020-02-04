import Vue from 'vue'
import VueRouter from 'vue-router'
import ManagerComponent from "../components/manager-component.vue"
Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Manager',
    component: ManagerComponent
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router
