import * as d3 from "d3";
import Block from "./Block";

// eslint-disable-next-line no-unused-vars
export default class Graph {
  // TODO: This needs a singleton
  constructor() {
    this.svg = null;
    this.blocks = {};
    this.connectors = {};
  }

  Init(divName) {
    // Create a svg canvas
    this.svg = d3
      .select(divName)
      .append("svg")
      .attr("width", 700)
      .attr("height", 500);
  }

  CreateBlock(blockName) {
    this.blocks[blockName] = new Block(this.svg);
  }

  CreateConnector(connectorName) {
    this.connectors[connectorName] = this.svg
      .append("line")
      .style("stroke", "black")
      .attr("x1", 150)
      .attr("y1", 100)
      .attr("x2", 250)
      .attr("y2", 300);
  }
}
