import encodeForHtml from "extlib/js/encodeForHtml";
import { exportTarget, parse } from "./_common";

// Prevent accidental XSS.
const assertValidName = (raw) => {
  if (!/^[a-zA-Z0-9-:]+$/.test(raw)) {
    throw new Error(`Invalid tag/attribute name: ${raw}`);
  }
  return raw;
};

const gen = ({ elem, props, children }) =>
  [
    "<",
    assertValidName(elem),
    ...Object.entries(props)
      .map(([n, v]) => {
        if (typeof v == "boolean") {
          if (v) {
            return " " + assertValidName(n);
          }
        } else {
          return [" ", assertValidName(n), '="', encodeForHtml(v), '"'];
        }
      })
      .filter((a) => a != undefined),
    ">",
    ...children.map((c) => (typeof c == "string" ? encodeForHtml(c) : gen(c))),
    "</",
    elem,
    ">",
  ].join("");

exportTarget.h = (...args) => gen(parse(...args));
