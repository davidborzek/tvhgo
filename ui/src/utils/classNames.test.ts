import { expect, test } from 'vitest';
import { c } from './classNames';

test('should return correct class name', () => {
  const className = c('first', undefined, 'second');
  expect(className).toEqual('first second');
});

test('should return empty class name', () => {
  const className = c();
  expect(className).toEqual('');
});
