import Vue from 'vue'
import App from './App.vue'
import { Form } from 'vant';
import { Field } from 'vant';
import { Button } from 'vant';
import router from './router/index'
import axios from "axios";

Vue.use(Form);
Vue.use(Field);
Vue.use(Button);
Vue.config.productionTip = false

// const url = "http://127.0.0.1:12555"
const url = "http://info-check.ronething.com"

Vue.prototype.$url = url
Vue.prototype.$axios = axios
new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
