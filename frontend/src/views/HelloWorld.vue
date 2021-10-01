<template>
  <van-form @submit="onSubmit">
    <van-field
      v-model="username"
      name="name"
      label="姓名"
      placeholder="姓名"
      :rules="[{ required: true, message: '请填写姓名' }]"
    ></van-field>
    <van-field
      v-model="idNumber"
      name="code"
      label="身份证后四位"
      placeholder="身份证后四位"
      :rules="[{ required: true, message: '请填写身份证后四位' }]"
    ></van-field>
    <div style="margin: 16px;">
      <van-button round block type="info" native-type="submit">提交确认</van-button>
    </div>
  </van-form>
</template>

<script>
import 'vant/lib/index.css'
import { Toast } from 'vant';

export default {
  name: 'HelloWorld',
  props: {
  },
  data() {
    return {
      username: '',
      idNumber: '',
    };
  },
  methods: {
    onSubmit(values) {
      Toast.loading({
        message: '加载中...',
        forbidClick: true,
      });
      this.$axios.get(`${this.$url}/api?code=${values.code}&name=${values.name}`)
      .then(res => {
        Toast.clear()
        const { data } = res
        if(data.errCode >= 0) {
          this.$router.push({name: 'detail', params: { data: data.data }})
        }
        else {
          Toast(data.errMsg)
        }
      })
      .catch(e => {
        console.log(e)
        Toast.clear()
      })
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
