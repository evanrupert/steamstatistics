import React, { Component } from 'react';
import './App.css';
import {getTestData} from './DataService.js'
import {Jumbotron} from 'react-bootstrap'
import {VanityUsernameInput} from './VanityUsernameInput.js'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {data: null};

    this.getDataForUsername = this.getDataForUsername.bind(this);
  }

  getDataForUsername(_username) {
    getTestData((resp) => this.setState({data: resp}));
  }

  getMainContent() {
    if (this.state.data === null) {
      return (
        <VanityUsernameInput onSubmit={this.getDataForUsername}/>
      )
    } else {
      return (
        <p>Placeholder for graph</p>
      )
    }
  }

  render() {
    return (
      <div className="App container">
        <Jumbotron>
          <h1 className="app-title">Steam Tag Playtime Calculator</h1>
          <div className="main-content">
            {this.getMainContent()}
          </div>
        </Jumbotron>
      </div>
    );
  }
}

export default App;
