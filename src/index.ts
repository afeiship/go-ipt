import fde from 'fast-deep-equal';
import filterKeysDeep from '@jswork/filter-keys-deep';

export interface IsDataEqualOptions {
  ignoreKeys?: string[];
}

function filterByCommonKeys(obj1: any, obj2: any) {
  const commonKeys = Object.keys(obj1).filter((key) =>
    obj2.hasOwnProperty(key),
  );
  const filterKeys = (obj, keys) => {
    if (Array.isArray(obj)) {
      return obj.map((item) => filterKeys(item, keys));
    } else if (obj !== null && typeof obj === 'object') {
      return keys.reduce((acc, key) => {
        if (obj.hasOwnProperty(key)) {
          acc[key] = filterKeys(obj[key], keys);
        }
        return acc;
      }, {});
    }
    return obj;
  };

  return [filterKeys(obj1, commonKeys), filterKeys(obj2, commonKeys)];
}

const isDataEqual = (oldData: any, newData: any, options?: IsDataEqualOptions): boolean => {
  const { ignoreKeys = [] } = options || {};
  const filteredObj1 = filterKeysDeep(oldData, ignoreKeys);
  const filteredObj2 = filterKeysDeep(newData, ignoreKeys);

  const [commonObj1, commonObj2] = filterByCommonKeys(
    filteredObj1,
    filteredObj2,
  );
  return fde(commonObj1, commonObj2);
};

export default isDataEqual;
