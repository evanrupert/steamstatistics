import React, { Component } from 'react';
import './VanityUsernameInput.css'
import {Button, TextField} from '@material-ui/core'

export class VanityUsernameInput extends Component {
  constructor(props) {
    super(props);
    this.state = {value: '', error: props.error};

    this.handleTextChange = this.handleTextChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleTextChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    console.log('handleSubmit: ' + this.state.value);
    this.props.onSubmit(this.state.value);
    event.preventDefault();
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
      </div>
    );
  }
}
