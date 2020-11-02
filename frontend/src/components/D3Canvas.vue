<template>
  <div id="drawArea">
    <button v-on:click="CreateNewProcessor()">Create new Processor</button>
    <select v-model="selected">
      <option disabled value="">Please select one</option>
      <option>ProcessorA</option>
      <option>ProcessorB</option>
      <option>ProcessorC</option>
    </select>
    <p>
      <span>Selected: {{ selected }}</span>
    </p>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "D3Canvas",
  data: function() {
    return {
      selected: "",
      graph: {}
    };
  },
  mounted() {
    this.InitGraph("#drawArea");
  },
  methods: {
    ...mapActions("gofi", ["InitGraph", "CreateBlock"]),
    CreateNewProcessor: function() {
      this.CreateBlock({
        type: this.selected
      });
      console.log("creating new processor " + this.selected);
    }
  }
};
</script>

<style scoped></style>
