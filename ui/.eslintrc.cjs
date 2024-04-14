const prettierConfig = require('./.prettierrc.cjs');

module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
    'plugin:prettier/recommended',
    'plugin:import/typescript',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs'],
  parser: '@typescript-eslint/parser',
  rules: {
    camelcase: 'warn',
    eqeqeq: 'error',
    'no-duplicate-imports': 'error',
    'no-unused-expressions': 'error',
    'no-unused-labels': 'error',
    'prefer-const': 'error',
    'prefer-template': 'error',
    'prettier/prettier': ['error', prettierConfig],
    'sort-imports': 'error',
  },
};
