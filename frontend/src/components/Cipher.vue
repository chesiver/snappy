<script lang="ts" setup>
import { state } from '../state';

</script>

<template>
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
            ]" v-model="state.cipher"></v-select>
        </v-col>
    </v-row>
</template>