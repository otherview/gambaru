<template>
  <div
    id="canvas"
    class="jtk-demo-canvas canvas-wide flowchart-demo jtk-surface jtk-surface-nopan"
  >
    <div
      :id="'flowchartWindow'+item.index"
      class="window jtk-node"
      v-for="item in windows"
      :key="item.index"
    >
      {{ item.index }}
    </div>
    <!--    <div id="flowchartWindow1" class="window jtk-node">1</div>-->
    <!--    <div id="flowchartWindow2" class="window jtk-node">2</div>-->
    <!--    <div id="flowchartWindow3" class="window jtk-node">3</div>-->
    <!--    <div id="flowchartWindow4" class="window jtk-node">4</div>-->
  </div>
</template>

<script>
import JSPlumb from "jsplumb";
export default {
  name: "JSPlumb",
  data: function() {
    return {
      windows: [{ index: 1 }, { index: 2 },{ index: 3 },{ index: 4 }]
    };
  },
  mounted() {
    JSPlumb.jsPlumb.ready(function() {
      const instance = JSPlumb.jsPlumb.getInstance({
        // default drag options
        DragOptions: { cursor: "pointer", zIndex: 2000 },
        // the overlays to decorate each connection with.  note that the label overlay uses a function to generate the label text; in this
        // case it returns the 'labelText' member that we set on each connection in the 'init' method below.
        ConnectionOverlays: [
          [
            "Arrow",
            {
              location: 1,
              visible: true,
              width: 11,
              length: 11,
              id: "ARROW",
              events: {
                click: function() {
                  alert("you clicked on the arrow overlay");
                }
              }
            }
          ],
          [
            "Label",
            {
              location: 0.1,
              id: "label",
              cssClass: "aLabel",
              events: {
                tap: function() {
                  alert("hey");
                },
                drag: function() {
                  alert("hey");
                }
              }
            }
          ]
        ],
        Container: "canvas"
      });
      const basicType = {
        connector: "StateMachine",
        paintStyle: { stroke: "red", strokeWidth: 4 },
        hoverPaintStyle: { stroke: "blue" },
        overlays: ["Arrow"]
      };
      instance.registerConnectionType("basic", basicType);
      // this is the paint style for the connecting lines..
      const connectorPaintStyle = {
          strokeWidth: 2,
          stroke: "#61B7CF",
          joinstyle: "round",
          outlineStroke: "white",
          outlineWidth: 2
        },
        // .. and this is the hover style.
        connectorHoverStyle = {
          strokeWidth: 3,
          stroke: "#216477",
          outlineWidth: 5,
          outlineStroke: "white"
        },
        endpointHoverStyle = {
          fill: "#4157d4",
          stroke: "#216477"
        },
        // the definition of source endpoints (the small blue ones)
        sourceEndpoint = {
          endpoint: "Dot",
          paintStyle: {
            stroke: "#7AB02C",
            fill: "transparent",
            radius: 7,
            strokeWidth: 1
          },
          isSource: true,
          connector: [
            "Flowchart",
            {
              stub: [40, 60],
              gap: 10,
              cornerRadius: 5,
              alwaysRespectStubs: true
            }
          ],
          connectorStyle: connectorPaintStyle,
          hoverPaintStyle: endpointHoverStyle,
          connectorHoverStyle: connectorHoverStyle,
          dragOptions: {},
          overlays: [
            [
              "Label",
              {
                location: [0.5, 1.5],
                label: "Drag",
                cssClass: "endpointSourceLabel",
                visible: false
              }
            ]
          ]
        },
        // the definition of target endpoints (will appear when the user drags a connection)
        targetEndpoint = {
          endpoint: "Dot",
          paintStyle: { fill: "#7AB02C", radius: 7 },
          hoverPaintStyle: endpointHoverStyle,
          maxConnections: -1,
          dropOptions: { hoverClass: "hover", activeClass: "active" },
          isTarget: true,
          overlays: [
            [
              "Label",
              {
                location: [0.5, -0.5],
                label: "Drop",
                cssClass: "endpointTargetLabel",
                visible: false
              }
            ]
          ]
        },
        init = function(connection) {
          connection
            .getOverlay("label")
            .setLabel(
              connection.sourceId.substring(15) +
                "-" +
                connection.targetId.substring(15)
            );
        };

      var _addEndpoints = function(toId, sourceAnchors, targetAnchors) {
        for (var i = 0; i < sourceAnchors.length; i++) {
          var sourceUUID = toId + sourceAnchors[i];
          instance.addEndpoint("flowchart" + toId, sourceEndpoint, {
            anchor: sourceAnchors[i],
            uuid: sourceUUID
          });
        }
        for (var j = 0; j < targetAnchors.length; j++) {
          var targetUUID = toId + targetAnchors[j];
          instance.addEndpoint("flowchart" + toId, targetEndpoint, {
            anchor: targetAnchors[j],
            uuid: targetUUID
          });
        }
      };

      const window = document.createElement("div");
      window.id = "flowchartWindow6";
      window.className =
        "window jtk-node jtk-endpoint-anchor jtk-managed jtk-draggable";
      window.innerHTML = "derp";
      document.getElementById("canvas").appendChild(window);

      //debugger;

      // suspend drawing and initialise.
      instance.batch(function() {
        _addEndpoints(
          "Window4",
          ["TopCenter", "BottomCenter"],
          ["LeftMiddle", "RightMiddle"]
        );
        _addEndpoints(
          "Window2",
          ["LeftMiddle", "BottomCenter"],
          ["TopCenter", "RightMiddle"]
        );
        _addEndpoints(
          "Window3",
          ["RightMiddle", "BottomCenter"],
          ["LeftMiddle", "TopCenter"]
        );
        _addEndpoints(
          "Window1",
          ["LeftMiddle", "RightMiddle"],
          ["TopCenter", "BottomCenter"]
        );

        _addEndpoints(
          "Window6",
          ["LeftMiddle", "RightMiddle"],
          ["TopCenter", "BottomCenter"]
        );

        // listen for new connections; initialise them the same way we initialise the connections at startup.
        instance.bind("connection", function(connInfo) {
          init(connInfo.connection);
        });

        // make all the window divs draggable
        instance.draggable(
          JSPlumb.jsPlumb.getSelector(".flowchart-demo .window"),
          {
            grid: [20, 20]
          }
        );
        // THIS DEMO ONLY USES getSelector FOR CONVENIENCE. Use your library's appropriate selector
        // method, or document.querySelectorAll:
        //JSPlumb.draggable(document.querySelectorAll(".window"), { grid: [20, 20] });

        // connect a few up
        instance.connect({
          uuids: ["Window2BottomCenter", "Window3TopCenter"]
        });
        instance.connect({ uuids: ["Window2LeftMiddle", "Window4LeftMiddle"] });
        instance.connect({ uuids: ["Window4TopCenter", "Window4RightMiddle"] });
        instance.connect({
          uuids: ["Window3RightMiddle", "Window2RightMiddle"]
        });
        instance.connect({
          uuids: ["Window4BottomCenter", "Window1TopCenter"]
        });
        instance.connect({
          uuids: ["Window3BottomCenter", "Window1BottomCenter"]
        });
        //

        //
        // listen for clicks on connections, and offer to delete connections on click.
        //
        instance.bind("click", function(conn) {
          // if (confirm("Delete connection from " + conn.sourceId + " to " + conn.targetId + "?"))
          //   instance.detach(conn);
          conn.toggleType("basic");
        });

        instance.bind("connectionDrag", function(connection) {
          console.log(
            "connection " +
              connection.id +
              " is being dragged. suspendedElement is ",
            connection.suspendedElement,
            " of type ",
            connection.suspendedElementType
          );
        });

        instance.bind("connectionDragStop", function(connection) {
          console.log("connection " + connection.id + " was dragged");
        });

        instance.bind("connectionMoved", function(params) {
          console.log("connection " + params.connection.id + " was moved");
        });
      });

      JSPlumb.jsPlumb.fire("jsPlumbDemoLoaded", instance);
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.item {
  height: 50px;
  width: 50px;
  background-color: red;
  display: inline-block;
}
.demo {
  /* for IE10+ touch devices */
  touch-action: none;
}

.flowchart-demo .window {
  border: 1px solid #346789;
  box-shadow: 2px 2px 19px #aaa;
  -o-box-shadow: 2px 2px 19px #aaa;
  -webkit-box-shadow: 2px 2px 19px #aaa;
  -moz-box-shadow: 2px 2px 19px #aaa;
  -moz-border-radius: 0.5em;
  border-radius: 0.5em;
  opacity: 0.8;
  width: 80px;
  height: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  text-align: center;
  z-index: 20;
  position: absolute;
  background-color: #eeeeef;
  color: black;
  font-family: helvetica, sans-serif;
  padding: 0.5em;
  font-size: 0.9em;
  -webkit-transition: -webkit-box-shadow 0.15s ease-in;
  -moz-transition: -moz-box-shadow 0.15s ease-in;
  -o-transition: -o-box-shadow 0.15s ease-in;
  transition: box-shadow 0.15s ease-in;
}

.flowchart-demo .window:hover {
  box-shadow: 2px 2px 19px #444;
  -o-box-shadow: 2px 2px 19px #444;
  -webkit-box-shadow: 2px 2px 19px #444;
  -moz-box-shadow: 2px 2px 19px #444;
  opacity: 0.6;
}

.flowchart-demo .active {
  border: 1px dotted green;
}

.flowchart-demo .hover {
  border: 1px dotted red;
}

#flowchartWindow1 {
  top: 34em;
  left: 5em;
}

#flowchartWindow2 {
  top: 7em;
  left: 36em;
}

#flowchartWindow3 {
  top: 27em;
  left: 48em;
}

#flowchartWindow4 {
  top: 23em;
  left: 22em;
}

.flowchart-demo .jtk-connector {
  z-index: 4;
}

.flowchart-demo .jtk-endpoint,
.endpointTargetLabel,
.endpointSourceLabel {
  z-index: 21;
  cursor: pointer;
}

.flowchart-demo .aLabel {
  background-color: white;
  padding: 0.4em;
  font: 12px sans-serif;
  color: #444;
  z-index: 21;
  border: 1px dotted gray;
  opacity: 0.8;
  cursor: pointer;
}

.flowchart-demo .aLabel.jtk-hover {
  background-color: #5c96bc;
  color: white;
  border: 1px solid white;
}

.window.jtk-connected {
  border: 1px solid green;
}

.jtk-drag {
  outline: 4px solid pink !important;
}

path,
.jtk-endpoint {
  cursor: pointer;
}

.jtk-overlay {
  background-color: transparent;
}
/* ---------------------------------------------------------------------------------------------------- */
/* --- page structure --------------------------------------------------------------------------------- */
/* ---------------------------------------------------------------------------------------------------- */

body {
  background-color: #fff;
  color: #434343;
  font-family: "Lato", sans-serif;
  font-size: 14px;
  font-weight: 400;
  height: 100%;
  padding: 0;
}

.jtk-bootstrap {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.jtk-bootstrap .jtk-page-container {
  display: flex;
  width: 100vw;
  justify-content: center;
  flex: 1;
}

.jtk-bootstrap .jtk-container {
  width: 60%;
  max-width: 800px;
}

.jtk-bootstrap-wide .jtk-container {
  width: 80%;
  max-width: 1187px;
}

.jtk-demo-main {
  position: relative;
  margin-top: 98px;
  display: flex;
  flex-direction: column;
}

.jtk-demo-main .description {
  font-size: 13px;
  margin-top: 25px;
  padding: 13px;
  margin-bottom: 22px;
  background-color: #f4f5ef;
}

.jtk-demo-main .description li {
  list-style-type: disc !important;
}

.jtk-demo-canvas {
  height: 750px;
  max-height: 700px;
  border: 1px solid #ccc;
  background-color: white;
  display: flex;
  flex-grow: 1;
}

.canvas-wide {
  margin-left: 0;
}

.miniview {
  position: absolute;
  top: 25px;
  right: 25px;
  z-index: 100;
}

.jtk-demo-dataset {
  text-align: left;
  max-height: 600px;
  overflow: auto;
}

.demo-title {
  float: left;
  font-size: 18px;
}

.controls {
  top: 25px;
  color: #fff;
  margin-right: 10px;
  position: absolute;
  left: 25px;
  z-index: 1;
  display: flex;
}

.controls i {
  background-color: #5184a0;
  border-radius: 4px;
  cursor: pointer;
  margin-right: 4px;
  padding: 4px;
}

li {
  list-style-type: none;
}

/* ------------------------ node palette -------------------- */

.sidebar {
  margin: 0;
  padding: 10px 0;
  background-color: white;
  display: flex;
  flex-direction: column;
  border: 1px solid #ccc;
  align-items: center;
}

.sidebar-item {
  background-color: #ccc;
  border-radius: 11px;
  color: #585858;
  cursor: move;
  padding: 8px;
  width: 128px;
  text-align: center;
  margin: 10px;
  outline: none;
}

button.sidebar-item {
  cursor: pointer;
  width: 150px;
}

.sidebar select {
  height: 35px;
  width: 150px;
  outline: none;
}

.sidebar-item.katavorio-clone-drag {
  margin: 0;
  border: 1px solid white;
}

.sidebar-item:hover,
.sidebar-item.katavorio-clone-drag {
  background-color: #5184a0;
  color: white;
}

/*
  .sidebar button {
      background-color: #30686d;
      outline: none;
      border: none;
      margin-left: 25px;
      padding: 7px;
      color: white;
      cursor:pointer;
  }*/

.sidebar i {
  float: left;
}

@media (max-width: 600px) {
  .sidebar {
    float: none;
    height: 55px;
    width: 100%;
    padding-top: 0;
  }

  .sidebar ul li {
    display: inline-block;
    margin-top: 7px;
    width: 67px;
  }
  .jtk-demo-canvas {
    margin-left: 0;
    margin-top: 10px;
    height: 364px;
  }
}

/* ---------------------------------------------------------------------------------------------------- */
/* --- jsPlumb setup ---------------------------------------------------------------------------------- */
/* ---------------------------------------------------------------------------------------------------- */

.jtk-surface-pan {
  display: none;
}

.jtk-connector {
  z-index: 9;
}

.jtk-connector:hover,
.jtk-connector.jtk-hover {
  z-index: 10;
}

.jtk-endpoint {
  z-index: 12;
  opacity: 0.8;
  cursor: pointer;
}

.jtk-overlay {
  background-color: white;
  color: #434343;
  font-weight: 400;
  padding: 4px;
  z-index: 10;
}

.jtk-overlay.jtk-hover {
  color: #434343;
}

path {
  cursor: pointer;
}

.delete {
  padding: 2px;
  cursor: pointer;
  float: left;
  font-size: 10px;
  line-height: 20px;
}

.add,
.edit {
  cursor: pointer;
  float: right;
  font-size: 10px;
  line-height: 20px;
  margin-right: 2px;
  padding: 2px;
}

.edit:hover {
  color: #ff8000;
}

.selected-mode {
  color: #e4f013;
}

.connect {
  width: 10px;
  height: 10px;
  background-color: #f76258;
  position: absolute;
  bottom: 13px;
  right: 5px;
}

/* header styles */

.demo-links {
  position: fixed;
  right: 0;
  top: 57px;
  font-size: 11px;
  background-color: white;
  opacity: 0.8;
  padding-right: 10px;
  padding-left: 5px;
  text-transform: uppercase;
  z-index: 100001;
}

.demo-links div {
  display: inline;
  margin-right: 7px;
  margin-left: 7px;
}

.demo-links i {
  padding: 4px;
}

.jtk-node {
  background-color: #5184a0;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  position: absolute;
  z-index: 11;
  overflow: hidden;
  min-width: 80px;
  min-height: 30px;
  width: auto;
}

.jtk-node .name {
  color: white;
  cursor: move;
  font-size: 13px;
  line-height: 24px;
  padding: 6px;
  text-align: center;
}

.jtk-node .name span {
  cursor: pointer;
}

[undo],
[redo] {
  background-color: darkgray !important;
}
[can-undo="true"] [undo],
[can-redo="true"] [redo] {
  background-color: #3e7e9c !important;
}
</style>
