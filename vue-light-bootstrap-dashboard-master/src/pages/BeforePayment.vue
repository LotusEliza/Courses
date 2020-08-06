<template>
  <div>
    <h2>Для того что бы продолжить обучение по курсу </h2>
      <h2>"{{course.Name}}"</h2>
      <h2>необходимо оплатить {{course.Price}} грн.</h2>
<!--    <b-button type="is-info" outlined @click="buyCourse">Оплатить<i class="nc-icon nc-tap-01"></i></b-button>-->
    <form method="POST" action="https://www.liqpay.ua/api/3/checkout" accept-charset="utf-8">
      <input type="hidden" name="data" :value="data.data"/>
      <input type="hidden" name="signature" :value="data.signature"/>
      <input type="image" src="//static.liqpay.ua/buttons/p1ru.radius.png"/>
    </form>
  </div>
</template>

<script>
  // import CryptoJS from '../../crypto-js';

  // import Vue from 'vue'
  // import sha1 from 'sha1';
  // Vue.use(sha1)
  import {mapGetters} from "vuex";

  export default {
    name: "BeforePayment.vue",
    data() {
      return {
        course: [],
        json_string:
          {
            "public_key":"sandbox_i79527556721",
            "version":"3",
            "action":"pay",
            "amount":"1",
            "currency":"UAH",
            "description":"test",
            "order_id":"000001"
          },
        dt: '',
        signature: '',
        privateKey: 'sandbox_9Cis7YaWEp5RHjmmq8c0FvYot35LbiT9kmJuQLBj',
        // data: btoa(json_string),
        // public_key
        // sandbox_i79527556721
        // private_key
        // sandbox_9Cis7YaWEp5RHjmmq8c0FvYot35LbiT9kmJuQLBj
      }
    },
    mounted() {
      this.course = this.$store.getters['courses/courseById'](this.$route.params.id_course)
      this.$store.dispatch('courses/buyCourse');
    },
    computed: {
      ...mapGetters("courses", [
        'data',
      ]),
    },
    methods:{
      // buyCourse(){
      //   this.$store.dispatch('courses/buyCourse');
      // },
      strToUtf8Bytes(str) {
        const utf8 = [];
        for (let ii = 0; ii < str.length; ii++) {
          let charCode = str.charCodeAt(ii);
          if (charCode < 0x80) utf8.push(charCode);
          else if (charCode < 0x800) {
            utf8.push(0xc0 | (charCode >> 6), 0x80 | (charCode & 0x3f));
          } else if (charCode < 0xd800 || charCode >= 0xe000) {
            utf8.push(0xe0 | (charCode >> 12), 0x80 | ((charCode >> 6) & 0x3f), 0x80 | (charCode & 0x3f));
          } else {
            ii++;
            // Surrogate pair:
            // UTF-16 encodes 0x10000-0x10FFFF by subtracting 0x10000 and
            // splitting the 20 bits of 0x0-0xFFFFF into two halves
            charCode = 0x10000 + (((charCode & 0x3ff) << 10) | (str.charCodeAt(ii) & 0x3ff));
            utf8.push(
              0xf0 | (charCode >> 18),
              0x80 | ((charCode >> 12) & 0x3f),
              0x80 | ((charCode >> 6) & 0x3f),
              0x80 | (charCode & 0x3f),
            );
          }
        }
        return utf8;
      },
      unpack(str) {
        var bytes = [];
        for(var i = 0; i < str.length; i++) {
          var char = str.charCodeAt(i);
          bytes.push(char >>> 8);
          bytes.push(char & 0xFF);
        }
        return bytes;
      }
    }
  }
</script>

