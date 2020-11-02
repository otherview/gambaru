<template>
  <div>
    <p>
      JSPLUMB
    </p>
    <JSPlumb></JSPlumb>
    <p>
      Derp
    </p>
    <D3Canvas style="height: 400px; width: 400px"></D3Canvas>
    <p>
      Derp end
    </p>
  </div>
</template>

<script>
import D3Canvas from "../components/D3Canvas";
import JSPlumb from "../components/JSPlumb";

export default {
  name: "Home",
  components: {
    JSPlumb,
    D3Canvas
  }
};
</script>
