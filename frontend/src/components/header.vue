<!--
 * @Description: 头部
 * @Author: chay
 * @Date: 2019-12-11 12:32:01
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2021-01-29 10:17:51
 -->
<template>
  <div>
    <header id="head-top">
      <i class="go-back" @click="goBack" v-if="showHandler"></i>
      <span class="title-text">团员信息</span>
    </header>
  </div>
</template>
<script>

export default {
  name: "headerView",
  components: {
  },
  props: {
    path: {
      type: String,
      default: '/',
    }
  },
  data() {
    return {
      showHandler: false,
    };
  },
  created() {
    
  },
  watch: {
    path (val) {
      if(val != '/') {
        this.showHandler = true
      }
      else {
        this.showHandler = false
      }
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    }
  },
  mounted(){

  }
};
</script>
<style>

#head-top {
  position: relative;
  background-color: #fff;
  border-bottom: 1px solid #eee;
  left: 0;
  top: 0;
  right: 0;
  padding: 2rem 0 1rem 0;
  text-align: center;
  font-size: 16px;
  color: #646566;
}
.go-back {
  display: inline-block;
  position: absolute;
  left: 1.5rem;
  width: 1rem;
  height: 1rem;
  border-left: 1px solid#646566;
  border-top: 1px solid#646566;
  transform: rotate(-45deg) translateY(.8rem);
}
</style>
