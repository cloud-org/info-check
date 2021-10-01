import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

const routes = [
  {
    path: "/",
    name: 'hello',
    component: () => import("../views/HelloWorld.vue"),
  },
  {
    path: "/detail",
    name: 'detail',
    component: () => import("../views/Detail.vue"),
  },
  {
    path: '/404',
    name: '404',
    // meta: {
    //   title: "404",
    //   requireAuth: false
    // },
    // component: () => import("@/views/common/404.vue"),
  }
];

const router = new Router({
  routes,
});


export default router;

