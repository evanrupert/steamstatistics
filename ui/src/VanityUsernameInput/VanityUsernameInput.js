import React, { Component } from 'react';
import './VanityUsernameInput.css'
import {Button, TextField, Snackbar} from '@material-ui/core'

export class VanityUsernameInput extends Component {
  constructor(props) {
    super(props);
    this.state = {
      value: '',
      error: null,
      snackBar: false
    };

    this.handleTextChange = this.handleTextChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleSnackBarClose = this.handleSnackBarClose.bind(this);
  }

  handleTextChange(event) {
    this.setState({value: event.target.value, error: null, snackBar: false});
  }

  handleSubmit(event) {
    console.log('handleSubmit: ' + this.state.value);
    let error = this.props.onSubmit(this.state.value);
    this.setState({
      value: this.state.value,
      error: error,
      snackBar: error !== null
    });

    event.preventDefault();
  }

  handleSnackBarClose() {
    this.setState({
      value: this.state.value,
      error: this.state.value,
      snackBar: false
    })
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <TextField value={this.state.value}
                     className="vanityUrlInput"
                     placeholder="Steam Vanity URL"
                     onChange={this.handleTextChange}
                     error={this.state.error !== null}
          />
          <Button variant="contained"
                  color="primary"
                  onClick={this.handleSubmit}>
            Get Data
          </Button>
        </form>
        <Snackbar open={this.state.snackBar}
                  autoHideDuration={2000}
                  message={this.state.error}
                  className="errorSnackBar"
                  onClose={this.handleSnackBarClose}
        />
      </div>
    );
  }
}
