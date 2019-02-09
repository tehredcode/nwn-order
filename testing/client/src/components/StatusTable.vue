<template>
  <v-data-table :headers="headers" :items="serverStats" class="elevation-1">
    <template slot="items" slot-scope="props">
      <td>{{ props.item.stat }}</td>
      <td class="text-xs-left">{{ props.item.info }}</td>
    </template>
  </v-data-table>
</template>
<script>
/* eslint-disable */
import axios from "axios";
export default {
  data() {
    return {
      info: {},
      errored: false,
      loading: true,

      headers: [
        { text: "Stat", align: "left", value: "stat" },
        { text: "Info", value: "info" }
      ],
      serverStats: []
    };
  },

  mounted() {
    axios
      .get("http://127.0.0.1:5750/api/stats")
      .then(response => {
        this.info = response.data;
        for (let prop in this.info) {
          this.serverStats.push({ stat: prop, info: this.info[prop] });
        }
      })
      .catch(error => {
        console.log(error + " did not get stats");
        this.errored = true;
      })
      .finally(() => (this.loading = false));
  }
};
</script>
<style>
</style>
