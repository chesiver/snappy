<script lang="ts" setup>
import { reactive } from 'vue';
import HelloWorld from './components/HelloWorld.vue'

const message = reactive({ text: 'test' });

</script>

<template>
  <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png" />
  <p>Using text interpolation: {{ message.text }}</p>
  <v-btn>
    Button
  </v-btn>
  <HelloWorld />
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
