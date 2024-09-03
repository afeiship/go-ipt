import isPartialEqual from '../src';

describe('api.basic', () => {
  test('01/fast equal', () => {
    const data1 = { a: 1, b: 2, c: 3 };
    const data2 = { a: 1, b: 2, c: 3 };
    expect(isPartialEqual(data1, data2)).toBe(true);
  });
  test('02/field not equal, but result is true', () => {
    const data1 = { a: 1, b: 2, c: 3 };
    const data2 = { a: 1, b: 2 };
    expect(isPartialEqual(data1, data2)).toBe(true);
  });
  test('03/field has ignoreKeys: is_editing', () => {
    const data1 = { a: 1, b: 2, c: 3, is_editing: true };
    const data2 = { a: 1, b: 2, c: 3 };
    const data3 = { a: 1, b: 2, c: 3, is_editing: false };
    expect(isPartialEqual(data1, data2, { ignoreKeys: ['is_editing'] })).toBe(true);
    expect(isPartialEqual(data1, data3, { ignoreKeys: ['is_editing'] })).toBe(true);
  });
});
