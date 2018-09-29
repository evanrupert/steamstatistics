import React, { Component } from 'react';
import './App.css';
import {getTagPlaytimeData} from './DataService.js'
import {Jumbotron} from 'react-bootstrap'
import {VanityUsernameInput} from './VanityUsernameInput.js'
import {PlaytimeGraph} from './PlaytimeGraph.js'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {data: null};

    this.getDataForUsername = this.getDataForUsername.bind(this);
  }

  getDataForUsername(username) {
    getTagPlaytimeData(username, (data) => this.setState({data: data}));
  }

  getMainContent() {
    if (this.state.data === null) {
      return (
        <VanityUsernameInput onSubmit={this.getDataForUsername}/>
      )
    } else {
      return (
        <PlaytimeGraph data={this.state.data} />
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
