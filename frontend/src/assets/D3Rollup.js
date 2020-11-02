import { select, selectAll, event, customEvent } from "d3-selection";
import { zoom } from "d3-zoom";
import { tree, hierarchy } from "d3-hierarchy";
import * as d31 from "d3/dist/d3";
import { drag } from "d3-drag";

export const d3 = {
  select,
  selectAll,
  tree,
  hierarchy,
  zoom,
  drag,
  // event,
  get event() {
    return event;
  },
  customEvent
};
