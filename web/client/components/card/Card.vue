<template>
    <div class="Card">
        <div class="Card__title">{{projectName}}</div>
        <div class="Card__data">
          <bar-chart :data="data" :chartId="projectName"></bar-chart>
        </div>
    </div>
</template>

<script>
  import {first} from 'lodash'
  import BarChart from '../stat/Bar'

  export default {
    components: { BarChart },
    props: {
      data: {
        required: true
      }
    },

    computed: {
      projectName () {
        return first(this.data).Name
      }
    }
  }
</script>

<style>
  .Card {
    border-radius:  5px;
    background:  #fff;
    flex: 0 1 250px;
    padding: 15px;
    margin: 10px;
    min-height: 325px;
    box-shadow: 0px 1px 7px 0px rgba(0, 0, 0, 0.15);
  }

  .Card__title {
    font-size: 13px;
    color: #81879B;
    font-family: 'Helvetica Neue';
    margin-bottom:  20px;
  }
</style>
