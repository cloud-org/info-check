<template>
  <van-form>
    <div class="user-box" v-if="showUser">
      <van-field
        v-model="name"
        name="name"
        label="姓名"
        placeholder="姓名"
        readonly
      ></van-field>
      <van-field
        v-model="gender"
        name="gender"
        label="性别"
        placeholder="性别"
        readonly
      ></van-field>
      <van-field
        v-model="id"
        name="id"
        label="身份证"
        placeholder="身份证"
        readonly
      ></van-field>
      <van-field
        v-model="phone"
        name="phone"
        label="手机号码"
        placeholder="手机号码"
        readonly
      ></van-field>
      <van-field
        v-model="comment"
        name="comment"
        label="备注"
        placeholder="备注"
        readonly
      ></van-field>
      <van-field
        v-model="room"
        name="room"
        label="酒店房间"
        placeholder="酒店房间"
        readonly
      ></van-field>
      <van-field
          v-model="amount"
          name="amount"
          label="团费补交金额"
          placeholder="团费补交金额"
          readonly
      ></van-field>
    </div>

    <div style="margin: 16px;" v-if="family">
      <van-button
        round
        block
        type="info"
        native-type="submit"
        @click="showMoreHanlder"
        >查看家属信息</van-button
      >
    </div>

    <div v-for="val in familyList" :key="val.id" class="padding10">
      <van-field
        v-model="val.name"
        name="name"
        label="姓名"
        placeholder="姓名"
        readonly
      ></van-field>
      <van-field
        v-model="val.gender"
        name="gender"
        label="性别"
        placeholder="性别"
        readonly
      ></van-field>
      <van-field
        v-model="val.id"
        name="id"
        label="身份证"
        placeholder="身份证"
        readonly
      ></van-field>
      <van-field
        v-model="val.phone"
        name="phone"
        label="手机号码"
        placeholder="手机号码"
        readonly
      ></van-field>
      <van-field
        v-model="val.comment"
        name="comment"
        label="备注"
        placeholder="备注"
        readonly
      ></van-field>
      <van-field
        v-model="val.room"
        name="room"
        label="酒店房间"
        placeholder="酒店房间"
        readonly
      ></van-field>
      <van-field
          v-model="val.amount"
          name="amount"
          label="团费补交金额"
          placeholder="团费补交金额"
          readonly
      ></van-field>
    </div>
  </van-form>
</template>

<script>
import "vant/lib/index.css";
import { Toast } from 'vant';
export default {
  name: "HelloWorld",
  props: {},
  data() {
    return {
      id: "",
      gender: "",
      phone: "",
      comment: "",
      room: "",
      name: "",
      amount: "",
      family: false,
      showUser: true,
      familyList: []
    };
  },
  created() {
    const params = this.$route.params;
    if (params && params.data) {
      for (let i in params.data) {
        this[i] = params.data[i];
      }
    }
  },
  methods: {
    showMoreHanlder() {
      Toast.loading({
        message: '加载中...',
        forbidClick: true,
      });
      this.$axios.get(`${this.$url}/api?code=${this.id}`)
      .then(res => {
        Toast.clear()
        const { data } = res;
        if(data.errCode >= 0) {
          this.familyList = data.data;
          this.family = false;
          this.showUser = false;
        }
        else {
          Toast(data.errMsg)
        }
      })
      .catch(e => {
        console.log(e)
        Toast.clear()
      })
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.padding10 {
  margin: 1rem;
  border: 1px solid #e8e6e6;
}
</style>
