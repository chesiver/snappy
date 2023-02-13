<script lang="ts" setup>
import { state } from '../state';
import { Connect, ListDirectoryEntries, ListDirectoryEntriesForPath } from '../../wailsjs/go/main/App';
import { ssh } from '../../wailsjs/go/models';
import { LogFatal, LogInfo } from '../../wailsjs/runtime/runtime';
import { ref, defineComponent } from 'vue';
import { VDataTable } from 'vuetify/labs/VDataTable'

defineComponent({
    components: {
        VDataTable,
    }
});

const isLoading = ref(false);
const entries = ref<ssh.FileInfo[]>([]);
const curDir = ref('')

const testConnection = async () => {
    LogInfo('Test Connection');
    isLoading.value = true;
    if (state.authType === 'publicKey') {
        await Connect(state.host, 'publicKey', [state.publicKeyFilePath]);
    } else if (state.authType === 'password') {
        await Connect(state.host, 'password', [state.username, state.password]);
    } else {
        LogFatal('Auth type not supported');
    }
    const { curDir: _curDir, fileInfos } = await ListDirectoryEntries();
    curDir.value = _curDir;
    entries.value = fileInfos
    isLoading.value = false;
}

const clickDirectoryName = async (item: ssh.FileInfo) => {
    if (item.isDir) {
        isLoading.value = true;
        const path = curDir.value + '/' + item.name
        LogInfo('clickDirectoryName path: ' + path);
        const { curDir: _curDir, fileInfos} = await ListDirectoryEntriesForPath(path);
        curDir.value = _curDir;
        entries.value = fileInfos
        isLoading.value = false;
    }
}

const clickBack = async () => {
    const path = curDir.value;
    const parentPath = path.slice(0, path.lastIndexOf('/'));
    isLoading.value = true;
    const { curDir: _curDir, fileInfos} = await ListDirectoryEntriesForPath(parentPath);
    curDir.value = _curDir;
    entries.value = fileInfos
    isLoading.value = false;
    isLoading.value = false;
}

</script>

<template>
    <v-row class="mb-2" align="center" no-gutters>
        <v-col cols="16">
            <v-text-field hide-details label="IP Address" v-model="state.host"></v-text-field>
        </v-col>
        <v-col cols="4">
            <v-text-field hide-details label="Port" v-model="state.port"></v-text-field>
        </v-col>
    </v-row>
    <v-row class="mb-2" align="center" no-gutters>
        <v-col class="mr-2" cols="3">
            <v-select hide-details label="Authentication Type" :items="[
                {
                    'title': 'Password',
                    'value': 'password'
                },
                {
                    'title': 'Public Key',
                    'value': 'publicKey'
                },
            ]" v-model="state.authType"></v-select>
        </v-col>
        <v-col class="mr-2" v-if="state.authType === 'publicKey'" cols="4">
            <v-text-field hide-details label="Public Key File Path" v-model="state.publicKeyFilePath"></v-text-field>
        </v-col>
        <v-col class="mr-2" v-if="state.authType === 'password'" cols="4">
            <v-text-field hide-details label="Username" v-model="state.username"></v-text-field>
            <v-text-field hide-details label="Password" v-model="state.password"></v-text-field>
        </v-col>
        <v-col cols="2">
            <v-btn :disabled="isLoading" @click="testConnection">Connect</v-btn>
        </v-col>
        <v-col cols="2">
            <v-progress-circular v-if="isLoading" indeterminate :size="30" color="lime">
            </v-progress-circular>
        </v-col>
    </v-row>
    <v-row class="mb-2" align="center" no-gutters>
        <v-col cols="6">
            <v-text-field hide-details justify="center" label="Current Directory" :modelValue="curDir"
                readonly></v-text-field>
        </v-col>
        <v-col class="ml-2" align="left" cols="4">
            <v-btn :disabled="isLoading" v-on:click="clickBack">Back</v-btn>
        </v-col>
    </v-row>
    <v-row align="center" no-gutters>
        <v-data-table v-if="!isLoading" :headers="[
            {
                title: 'Type',
                align: 'start',
                key: 'isDir',
            },
            {
                title: 'Name',
                align: 'start',
                key: 'name',
            },
            {
                title: 'Modified Time',
                align: 'start',
                key: 'modTime',
            },
        ]" :items="entries" item-value="name" :dense="true" items-per-page="10" :height="400" hide-no-data>
            <template v-slot:item="{ item }">
                <tr>
                    <td><v-icon :icon="item.columns.isDir ? 'mdi-folder':'mdi-file-account'"></v-icon></td>
                    <td><v-btn class="no-uppercase" v-on:click="clickDirectoryName(item.columns)">{{ item.columns.name }}</v-btn> </td>
                    <td>{{ item.columns.modTime }}</td>
                </tr>
            </template>
        </v-data-table>
    </v-row>


</template>

<style scoped>
.no-uppercase {
     text-transform: unset !important;
}
</style>