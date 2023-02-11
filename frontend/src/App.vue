<script lang="ts" setup>
import { reactive, ref } from 'vue';
import HelloWorld from './components/HelloWorld.vue'
import { ConnectShadowsocks, DumpLogContent } from '../wailsjs/go/main/App';
import { LogInfo } from '../wailsjs/runtime';

// const filePath = reactive({ fullPath: '' });

const logMessage = reactive({ message: '' });
const host = ref('10.36.189.159');
const port = ref('8488');
const cipher = ref('AEAD_CHACHA20_POLY1305');

const connectShadowsocks = async () => {
  setInterval(async () => {
    logMessage.message = await DumpLogContent(false);
  }, 1000);
  LogInfo('Connect with params: ' + [host.value, port.value, cipher.value].join(', '));
  await ConnectShadowsocks(host.value, port.value, cipher.value);
}

const clearLogMessage = async () => {
  logMessage.message = await DumpLogContent(true)
}



</script>

<template>
  <v-app>
    <v-app-bar app :elevation="2">
      <v-app-bar-title>Wails GUI for Shadowsocks</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <v-container grey lighten-5>
        <v-row class="mb-2" align="center" no-gutters>
          <v-col cols="16">
            <v-text-field label="IP Address" v-model="host"></v-text-field>
          </v-col>
          <v-col cols="4">
            <v-text-field label="Port" v-model="port"></v-text-field>
          </v-col>
        </v-row>
        <v-row no-gutters>
          <v-col cols="6">
            <v-select label="Cipher" :items="[
              {
                'title': 'Dummy',
                'value': 'DUMMY'
              },
              {
                'title': 'AES_128_GCM',
                'value': 'AEAD_AES_128_GCM'
              },
              {
                'title': 'AES_256_GCM',
                'value': 'AEAD_AES_256_GCM'
              },
              {
                'title': 'CHACHA20_POLY1305',
                'value': 'AEAD_CHACHA20_POLY1305'
              }
            ]" v-model="cipher"></v-select>
          </v-col>
        </v-row>
        <v-row class="mb-4" align="center" no-gutters>
          <v-col cols="2">
            <v-btn @click="connectShadowsocks">
              Connect
            </v-btn>
          </v-col>
          <v-col cols="2">
            <v-btn @click="clearLogMessage">
              Clear
            </v-btn>
          </v-col>
        </v-row>
        <v-row align="center" no-gutters>
          <v-textarea :value="logMessage.message" readonly rows="20"></v-textarea>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>

</style>
