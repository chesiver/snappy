import { reactive } from "vue";

export const state = reactive({
  host: '139.198.154.48',
  port: '8488',
  authType: 'password',
  publicKeyFilePath: '/Users/bytedance/.ssh/id_rsa',
  username: 'root',
  password: '!Wjlydntms38',
  cipher: 'AEAD_AES_256_GCM',
  logMessage: '',
});
