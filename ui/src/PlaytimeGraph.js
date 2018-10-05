import React, { Component } from 'react';

const ReactHighcharts = require('react-highcharts');

export class PlaytimeGraph extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: props.data,
    };

    this.createHighchartsConfig = this.createHighchartsConfig.bind(this);
    this.parseData = this.parseData.bind(this);
  }

  createHighchartsConfig(data) {
    return {
      chart: {
          type: 'column'
        },
        title: {
          text: "Playtime vs steam game tags"
        },
        xAxis: {
          type: 'category',
          labels: {
            rotation: -45,
            style: {
              fontSize: '13px',
              fontFamily: 'Verdana, sans-serif'
            }
          }
        },
        yAxis: {
          min: 0,
          title: {
            text: 'Playtime (Hours)'
          }
        },
        legend: {
          enabled: false
        },
        tooltip: {
          pointFormat: 'Playtime: <b>{point.y:.1f} hours</b>'
        },
        series: [{
          name: 'Playtime',
          data: data,
          dataLabels: {
            enabled: true,
            rotation: -90,
            color: '#FFFFFF',
            align: 'right',
            format: '{point.y:.1f}',
            y: 10, 
            style: {
              fontSize: '13px',
              fontFamily: 'Verdana, sans-serif'
            }
          }
        }]
      };
  }

  parseData(rawData) {
    return rawData.data.slice(0, 20).map((x) => [x.tag, x.playtime]);
  }

  render() {
    return (
      <div>
        <ReactHighcharts config={this.createHighchartsConfig(this.parseData(this.state.data))} />
      </div>
    )
  }
}