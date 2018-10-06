import React, { Component } from 'react';
import './App.css';
import {getTagPlaytimeData} from './DataService.js'
import {VanityUsernameInput} from './VanityUsernameInput/VanityUsernameInput.js'
import {PlaytimeGraph} from './PlaytimeGraph.js'
import {MuiThemeProvider, createMuiTheme} from '@material-ui/core/styles'

const theme = createMuiTheme({
  typography: {
    fontSize: 20
  }
});

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {data: null, error: null};

    this.getDataForUsername = this.getDataForUsername.bind(this);
  }

  getDataForUsername(username) {
    getTagPlaytimeData(username, (data) => {
      if (!data.ok) {
        this.setState({data: null, error: data.error});
      } else {
        this.setState({data: data, error: null});
      }
    });

    return this.state.error;
  }

  getMainContent() {
    if (this.state.data === null) {
      return (
        <VanityUsernameInput onSubmit={this.getDataForUsername} />
      )
    } else {
      return (
        <PlaytimeGraph data={this.state.data} />
      )
    }
  }

  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <div className="App container">
          <h1 className="app-title">Steam Tag Playtime Calculator</h1>
          <div className="main-content">
            {this.getMainContent()}
          </div>
        </div>
      </MuiThemeProvider>
    );
  }
}

export default App;
