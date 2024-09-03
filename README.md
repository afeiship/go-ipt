# is-partial-equal
> Fast data equal for business data.

[![version][version-image]][version-url]
[![license][license-image]][license-url]
[![size][size-image]][size-url]
[![download][download-image]][download-url]

## installation
```shell
npm install @jswork/is-partial-equal
```

## usage
```js
import isDataEqual from '@jswork/is-partial-equal';

const data1 = { name: 'afei', age: 25 };
const data2 = { name: 'afei' };

console.log(isDataEqual(data1, data2)); // true
```

## license
Code released under [the MIT license](https://github.com/afeiship/is-partial-equal/blob/master/LICENSE.txt).

[version-image]: https://img.shields.io/npm/v/@jswork/is-partial-equal
[version-url]: https://npmjs.org/package/@jswork/is-partial-equal

[license-image]: https://img.shields.io/npm/l/@jswork/is-partial-equal
[license-url]: https://github.com/afeiship/is-partial-equal/blob/master/LICENSE.txt

[size-image]: https://img.shields.io/bundlephobia/minzip/@jswork/is-partial-equal
[size-url]: https://github.com/afeiship/is-partial-equal/blob/master/dist/index.min.js

[download-image]: https://img.shields.io/npm/dm/@jswork/is-partial-equal
[download-url]: https://www.npmjs.com/package/@jswork/is-partial-equal
