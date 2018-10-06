import React, { Component } from 'react';
import './VanityUsernameInput.css'
import {Button, TextField, Tooltip} from '@material-ui/core'

const styles = {
  errorTooltip: {
    fontSize: 14,
    background: "#F31431"
  }
};

export class VanityUsernameInput extends Component {
  constructor(props) {
    super(props);
    this.state = {
      value: '',
      error: null
    };

    this.handleTextChange = this.handleTextChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
    this.getTextField = this.getTextField.bind(this);
  }

  handleTextChange(event) {
    this.setState({value: event.target.value, error: null});
  }

  handleSubmit(event) {
    console.log('handleSubmit: ' + this.state.value);
    let error = this.props.onSubmit(this.state.value);
    this.setState({value: this.state.value, error: error});
    event.preventDefault();
  }
  
  getTextField() {
    let textField =
      <TextField value={this.state.value}
                 className="vanityUrlInput"
                 placeholder="Steam Vanity URL"
                 onChange={this.handleTextChange}
                 error={this.state.error !== null}
      />;

    if (this.state.error === null)
      return textField;

    return (
      <Tooltip title={this.state.error}
               placement="left"
               classes={{tooltip: styles.errorTooltip}}
      >
        {textField}
      </Tooltip>
    )
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          {this.getTextField()}
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
