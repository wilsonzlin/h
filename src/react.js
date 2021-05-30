import {exportTarget, parse} from './_common';

let reactLib;
try {
  // Easier than figuring out which is global: window, global, this, self.
  reactLib = React;
} catch {
  reactLib = require('react');
}

exportTarget.h = (...args) => {
  const {elem, props, children} = parse(...args);
  return reactLib.createElement(elem, props, ...children);
};
