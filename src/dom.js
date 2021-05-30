import {exportTarget, parse} from './_common';

const createElem = ({elem, props, children}) => {
  const $elem = Object.assign(document.createElement(elem), props);
  for (const c of children) {
    // c is either a string or a DOM element (result of calling h).
    $elem.append(c);
  }
  return $elem;
}

exportTarget.h = (...args) => createElem(parse(...args))
