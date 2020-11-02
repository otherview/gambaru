// eslint-disable-next-line no-unused-vars
import * as d3 from "d3";

export default class Block {
  constructor(svg) {
    const closure = this;
    svg
      .append("svg:defs")
      .append("svg:marker")
      .attr("id", "triangle")
      .attr("refX", 6)
      .attr("refY", 6)
      .attr("markerWidth", 30)
      .attr("markerHeight", 30)
      .attr("orient", "auto")
      .append("path")
      .attr("d", "M 0 0 12 6 0 12 3 6")
      .style("fill", "black");

    this.blockGroup = svg
      .append("g")
      .call(d3.drag().on("start", this.DragBlock));

    this.blockGroup
      .append("rect")
      .attr("width", 120)
      .attr("height", 100)
      .attr("fill", "#69a3b2");
    this.blockGroup
      .append("rect")
      .attr("x", 10)
      .attr("y", 10)
      .attr("width", 30)
      .attr("height", 30)
      .attr("fill", "#f10663")
      .call(
        d3.drag().on("start", function(event) {
          closure.DragLine(event, closure.blockGroup);
        })
      );
  }

  DragBlock(event) {
    let rectangle = d3.select(this).classed("dragging", true);

    event.on("drag", dragged).on("end", ended);

    function dragged(event) {
      console.log(event);
      console.log(event.subject);
      rectangle
        .raise()
        .attr("transform", "translate(" + event.x + "," + event.y + ")");
    }

    function ended() {
      rectangle.classed("dragging", false);
    }
  }

  DragLine(event, blockGroup) {
    let line = blockGroup
      .append("line")
      .style("stroke", "lightgreen")
      .attr("stroke-width", 6)
      .attr("marker-end", "url(#triangle)")
      .raise();
    event.on("drag", dragged).on("end", ended);

    function dragged(event) {
      line
        .attr("x1", 1)
        .attr("y1", 1)
        .attr("x2", event.x)
        .attr("y2", event.y);

      // console.log(event.subject);
      // // Update the line properties
      // let attributes = {
      //   x1: 10,
      //   y1: 10,
      //   x2: 100,
      //   y2: 100
      // };
      // line.raise().attr(attributes);
    }

    function ended(event) {
      console.log(event);
      console.log(event.sourceEvent.target);
      const derp = d3.select(event.sourceEvent.target);
      console.log(derp);
      debugger;
      line.attr("x2", event.sourceEvent.target.x);
      line.attr("y2", event.sourceEvent.target.y);
      line.remove();
      //rectangle.classed("dragging", false);
    }
  }
}
