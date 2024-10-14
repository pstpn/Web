module.exports = {
  'env': {
    'browser': true,
    'es2021': true,
  },
  'extends': [
    'plugin:vue/essential',
    'google',
  ],
  'parserOptions': {
    'ecmaVersion': 12,
    'sourceType': 'module',
  },
  'plugins': [
    'vue',
  ],
  'rules': {
    'require-jsdoc': 'off',
    'max-len': ['error', {code: 130}],
  },
};
