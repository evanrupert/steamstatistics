import React, { Component } from 'react';
export class VanityUsernameInput extends Component {
  constructor(props) {
    super(props)
    this.state = {value: ''};

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
          <label>
            Vanity Username:
            <input type="text" 
                   name="vanityUsername" 
                   value={this.state.value}
                   onChange={this.handleTextChange}
            />
          </label>
          <input type="submit" value="Get Data" />
        </form>
      </div>
    );
  }
}
