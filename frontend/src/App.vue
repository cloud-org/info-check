<template>
  <div id="app">
    <header-top :goBack="true" :path="path"></header-top>
    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  </div>
</template>

<script>
import headerTop from "@/components/header";

export default {
  name: 'App',
  components: {
    headerTop
  },
  data() {
    return {
      path: ''
    }
  },
  beforeUpdate() {
    this.path = this.$route.path
  }
}
</script>

<style>
html {
  font-size: 62.5%;
}
html, body {
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>
