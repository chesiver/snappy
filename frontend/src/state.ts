import { reactive } from "vue";

export const state = reactive({
  host: '10.36.189.159',
  port: '8488',
  authType: 'password',
  publicKeyFilePath: '/Users/bytedance/.ssh/id_rsa',
  username: 'adam.liu',
  password: ',.!Wjlydntms38',
  cipher: 'AEAD_CHACHA20_POLY1305',
  logMessage: '',
});
