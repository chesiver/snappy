<script lang="ts" setup>
import { defineComponent, reactive, ref } from 'vue';
import Connection from './components/Connection.vue'
import Cipher from './components/Cipher.vue';
import Client from './components/Client.vue';
import { LogInfo } from '../wailsjs/runtime';
import { state } from './state';

defineComponent({
  components: {
    Connection,
    Cipher,
    Client
  }
});

const tab = ref("connection");
const logMessage = reactive({ message: '' });

</script>

<template>
  <v-app id="shadowsocks">
    <v-app-bar app :elevation="2" class="bg-light-blue">
      <v-app-bar-title>Wails GUI for Shadowsocks</v-app-bar-title>
    </v-app-bar>
    <v-navigation-drawer permanent expand-on-hover rail theme="dark" class="bg-deep-purple">
      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-view-dashboard-edit" title="Connection"
          @click="tab = 'connection'"></v-list-item>
        <v-list-item prepend-icon="mdi-lock" title="Cipher" @click="tab = 'cipher';"></v-list-item>
        <v-list-item prepend-icon="mdi-upload" title="Client" @click="tab = 'client';"></v-list-item>
        <v-list-item prepend-icon="mdi-server" title="Remote" @click="tab = 'remote';"></v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <v-container grey lighten-5>
        <Connection v-if="tab === 'connection'"></Connection>
        <Cipher v-if="tab === 'cipher'"></Cipher>
        <Client v-if="tab === 'client'"></Client>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>

</style>
