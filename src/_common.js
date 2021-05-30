export const parse = (selOrComp, props, childrenArr) => {
  let tagName = undefined;
  let id = undefined;
  const classes = [];
  if (typeof selOrComp == "string") {
    for (const p of selOrComp.split(/(?=[#.])/)) {
      if (p[0] == "#") {
        id = p.slice(1);
      } else if (p[0] == ".") {
        classes.push(p.slice(1));
      } else {
        tagName = p;
      }
    }
  }
  if (Array.isArray(props) || typeof props == "string") {
    childrenArr = props;
    props = undefined;
  }
  if (typeof childrenArr == "string") {
    childrenArr = [childrenArr];
  }
  props ??= {};
  if (id != undefined) {
    props.id = id;
  }
  if (classes.length) {
    props.className = classes.join(" ");
  }
  return {
    elem: tagName ?? selOrComp,
    props,
    children: childrenArr ?? [],
  };
};

export const exportTarget = typeof exports == 'object' ? exports : window;
