<template>
	<div class="page">
		<h1>bs3stat</h1>
		<div class="Cardlist">
			<card v-for="project in projects" :data="project">
			</card>
		</div>
	</div>
</template>

<script>
import _ from 'lodash'
import Card from '../components/card/Card'

export default {
	components: {
		Card
	},

	data () {
		return {
			backups: [],
			projects: []
		}
	},
	created () {
		this.$http.get('/projects').then((response) => {
			this.backups = response.data ? response.data : []
			this.getProjects()
		})
	},

	methods: {
		getProjects () {
			this.projects = _.groupBy(this.backups, 'Name')
		}
	}
}
</script>

<style>
	body {
		background: #F5F6F9;
	}

  h1 {
    font-family: 'Helvetica Neue';
    font-size: 28px;
    font-weight: 100;
    margin: 20px 0;
    color: #81879b;
  }

	.Cardlist {
		display: flex;
		flex-flow: row wrap;
		margin:  0 -10px;
	}

  .page {
    max-width: 1200px;
    margin: 0 auto;
    padding-left: 15px;
    padding-right: 15px;
    width: 100%;
  }
</style>
