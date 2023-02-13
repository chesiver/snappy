<script lang="ts" setup>
import { ConnectShadowsocks, DumpLogContent } from '../../wailsjs/go/main/App';
import { LogInfo } from '../../wailsjs/runtime/runtime';
import { state } from '../state';

const connectShadowsocks = async () => {
    setInterval(async () => {
        state.logMessage = await DumpLogContent(false);
    }, 1000);
    await ConnectShadowsocks(state.host, state.port, state.cipher);
}

const clearLogMessage = async () => {
    state.logMessage = await DumpLogContent(true)
}


</script>

<template>
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
        <v-textarea label="Socks Log" :value="state.logMessage" readonly rows="20"></v-textarea>
    </v-row>
</template>