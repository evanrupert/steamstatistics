import React, { Component } from 'react';

export class PlaytimeGraph extends Component {
  constructor(props) {
    super(props);
    this.state = {data: props.data};

    this.printData = this.printData.bind(this);
  }

  printData() {
    console.log(this.state.data);
  }

  render() {
    return (
      <div>
        <p>Placeholder for graph</p>
        <button onClick={this.printData}>Click to print data</button>
      </div>
    )
  }
}