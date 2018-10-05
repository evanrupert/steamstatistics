import React, { Component } from 'react';
import './App.css';
import {getTagPlaytimeData} from './DataService.js'
import {VanityUsernameInput} from './VanityUsernameInput/VanityUsernameInput.js'
import {PlaytimeGraph} from './PlaytimeGraph.js'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {data: null, error: null};

    this.getDataForUsername = this.getDataForUsername.bind(this);
  }

  getDataForUsername(username) {
    getTagPlaytimeData(username, (data) => {
      if (!data.ok) {
        this.setState({data: null, error: data.error})
      } else {
        this.setState({data: data, error: null})
      }
    });
  }

  getMainContent() {
    if (this.state.data === null) {
      return (
        <VanityUsernameInput
          error={this.state.error}
          onSubmit={this.getDataForUsername}
        />
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
        <h2 className="app-title">Steam Tag Playtime Calculator</h2>
        <div className="main-content">
          {this.getMainContent()}
        </div>
      </div>
    );
  }
}

export default App;
