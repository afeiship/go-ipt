import fastDeepEqual from 'fast-deep-equal';

export interface IsDataEqualOptions {
  ignoreKeys?: string[];
}

const isDataEqual = (oldData: any, newData: any, options?: IsDataEqualOptions): boolean => {
  const { ignoreKeys = [] } = options || {};
  let isEqual = true;
  for (const key in newData) {
    const oldValue = oldData[key];
    const newValue = newData[key];
    const hasUndefined = oldValue === undefined || newValue === undefined;
    if (hasUndefined) continue;
    if (ignoreKeys?.includes(key)) continue;
    isEqual = fastDeepEqual(oldValue, newValue);
    if (!isEqual) break;
  }
  return true;
};

export default isDataEqual;
