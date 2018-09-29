import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import {getTestData} from './DataService.js'

class App extends Component {
  printDataFromService() {
    getTestData((data) => console.log(data))
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
        <button onClick={this.printDataFromService}>Click me for data</button>
      </div>
    );
  }
}

export default App;
