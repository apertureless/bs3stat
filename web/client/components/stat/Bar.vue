<template>
  <svg :id="chartId" class="chart">
  </svg>
</template>

<script>
  import * as d3 from "d3";

  const zeroMargin = { top: 0, right: 0, bottom: 0, left: 0 }

  export default {
    props: {
      chartId: {
        default: 'barChart'
      },
      width: {
        default: 270,
        type: Number
      },
      height: {
        default: 130,
        type: Number
      },
      axis: {
        default: false,
        type: Boolean
      },
      axisPadding: {
        default: 5,
        type: Number
      },
      xDomain: {
        default: null
      },
      xDomain: {
        default: null
      },
      tickSize: {
        default: 10,
        type: Number
      },
      barPadding: {
        default: 20,
        type: Number
      },
      color: {
        default: ['RGB(0, 177, 240)', 'rgb(243, 43, 101)']
      },
      colorInterpolate: {
        default: d3.interpolateHcl
      },
      ease: {
        default: 'linear'
      },
      type: {
        default: 'rounded'
      },
      data: {
        required: true
      }
    },

    data (){
      return {
        margin: {
          left: 60,
          bottom: 35,
          top: 15,
          right: 0
        },
        dataset: [],
        chart: null,
        x: null,
        y: null,
        xBisect: null,
        xAxis: null,
        yAxis: null
      }
    },
    computed: {
      chartSelector () {
        return '#' + this.chartId
      }
    },
    mounted () {
      // this.dataset = _.map(this.data, 'duration')
      if (!this.axis) this.margin = zeroMargin

      // @TODO: Remove Tmp Data
      //
      for (var i = 10; i; i--) {
        this.dataset.push({
          bin: new Date(Date.now() - (i * 3600000)),
          value: Math.max(250, Math.random() * 3000 | 0),
          status: Math.floor(Math.random() * 3)
        })
      }
      this.init()
      this.renderAxis()
      this.renderBars()
    },

    methods: {
      dimensions () {
        const w = this.width - this.margin.left - this.margin.right
        const h = this.height - this.margin.top - this.margin.bottom
        return [w, h]
      },

      init () {
        const [w, h] = this.dimensions()

        this.chart = d3.select(this.chartSelector)
                      .attr('width', this.width)
                      .attr('height', this.height)
                      .append('g')
                      .attr('transform', `translate(${this.margin.left}, ${this.margin.top})`)

        if (this.color) {
          const color = d3.scaleLinear()
            .interpolate(this.colorInterpolate)
            .range(this.color)
        }

        this.x = d3.scaleTime()
          .range([0, w])
        this.y = d3.scaleLinear()
          .range([h, 0])

        this.xAxis = d3.axisBottom()
          .scale(this.x)
          .ticks(5)
          .tickPadding(8)
          .tickSize(this.tickSize)

        this.yAxis = d3.axisLeft()
          .scale(this.y)
          .ticks(3)
          .tickPadding(8)
          .tickSize(this.tickSize)

        if (this.axis) {
          this.chart.append('g')
            .attr('class', 'x axis')
            .attr('transform', `translate(0, ${h+this.axisPadding})`)
            .call(this.xAxis)

          this.chart.append('g')
            .attr('class', 'y axis')
            .attr('transform', `translate(${-this.axisPadding}, 0)`)
            .call(this.yAxis)
        }

        this.xBisect = d3.bisector(d => d.bin).left
      },

      renderAxis () {
        const yExtent = this.yDomain || d3.extent(this.dataset, d => d.value)
        const xd = this.x.domain(this.xDomain || d3.extent(this.dataset, d => d.bin))
        const yd = this.y.domain(yExtent)

        // if (this.color) this.color.domain(yExtent)

        if (this.nice) {
          xd.nice()
          yd.nice()
        }

        this.chart.select('.x.axis').call(this.xAxis)
        this.chart.select('.y.axis').call(this.yAxis)
      },

      renderBars () {
        const [w, h] = this.dimensions()
        const width = w / this.dataset.length
        const barWidth = width - this.barPadding

        if (barWidth < 1) throw new Error('BarChart is too small for the amount of data provided')

        const column = this.chart.selectAll('.column')
          .data(this.dataset)

        column.enter()
              .append('rect')
              .attr('class', 'column')
              .attr('x', d => this.x(d.bin))
              .attr('rx', this.type == 'rounded' ? barWidth / 2 : 0)
              .attr('ry', this.type == 'rounded' ? barWidth / 2 : 0)
              .attr('width', barWidth)
              .attr('height', h)

        column.exit().remove()

        const bar = this.chart.selectAll('.bar')
          .data(this.dataset)

        // enter
        bar.enter()
          .append('rect')
          .attr('class', d => this.typeClass(d.status))
          .attr('x', d => this.x(d.bin))
          .attr('y', d => this.y(d.value))
          .attr('rx', this.type == 'rounded' ? barWidth / 2 : 0)
          .attr('ry', this.type == 'rounded' ? barWidth / 2 : 0)
          .attr('width', barWidth)
          .attr('height', d => h - this.y(d.value))

        if (this.color) bar.style('fill', d => color(d.value))
        // exit
        bar.exit().remove()
      },

      typeClass (status) {
        let className

        if (status === 0) className = 'bar'
        if (status === 1) className = 'bar--warning'
        if (status === 2) className = 'bar--failure'
        console.log('status', status)
        console.log('className', className)
        return className
      }
    }
  }
</script>

<style>
.chart {
  /*border: 1px dotted #ddd;*/
}

.chart .column {
  fill: rgb(230, 237, 244);
}

.chart .bar {
  fill: rgb(42, 243, 141);
}

.chart .bar--warning {
  fill: rgb(243, 177, 42);
}

.chart .bar--failure {
  fill: rgb(243, 42, 100);
}

.chart .axis text {
  font: 9px sans-serif;
}

.chart .axis path,
.chart .axis line {
  fill: none;
  stroke: rgb(23x0, 237, 244);
  stroke-width: 1.5px;
  shape-rendering: crispEdges;
}

.chart .axis path {
  display: none;
}
</style>
