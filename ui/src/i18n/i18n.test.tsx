import { expect, test } from 'vitest';
import i18n, { t } from 'i18next';

test('can initialize i18n', async () => {
  // set prod to true, to reduce logs.
  import.meta.env.PROD = true;

  await import('./i18n');

  expect(i18n.isInitialized).toBeTruthy();
  expect(t('username')).toEqual('Username');
});
