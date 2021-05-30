import {exportTarget, parse} from './_common';

const createElem = ({elem, props, children}) => {
  const $elem = Object.assign(document.createElement(elem), props);
  for (const c of children) {
    $elem.append(typeof c == 'string' ? c : createElem(c));
  }
  return $elem;
}

exportTarget.h = (...args) => createElem(parse(...args))
