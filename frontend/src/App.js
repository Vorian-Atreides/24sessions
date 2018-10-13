import React, { Component } from 'react';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      ip: '',
      country: '',
      city: '',
      error: '',
    };
    this.onGetLocationClicked = this.onGetLocationClicked.bind(this);
  }

  /**
   * Events
   */

  async downloadLocation(ip) {
    try {
      const response = await fetch(`http://localhost:8080/geolocation/${ip}`);
      const body = await response.json();
      if (response.ok) {
        this.setState({ city: body.city, country: body.country, error: '' });
      } else {
        this.setState({ error: body.reason });
      }
    } catch(err) {
      this.setState({ error: err.message });
    }
  }

  async componentDidMount() {
    try {
      const response = await fetch('https://api.ipify.org?format=json');
      const body = await response.json();
      if (response.ok) {
        this.setState({ ip: body.ip, error: '' });
      }
    } catch(err) {
      this.setState({ error: err.message });
    }
  }

  onGetLocationClicked() {
    this.downloadLocation(this.state.ip);
  }

  /**
   * UI
   */

  render() {
    return (
      <div className="App">
        <p>
          Your IP address is: <span>{this.state.ip}</span>
        </p>
        {
          this.state.error.length > 0 &&
          <p>
            Oups, an error occured: <span>{this.state.error}</span>.
          </p>
        }
        { this.state.city.length > 0 && this.state.country.length > 0 &&
          <p>
            And you live in <span>{this.state.city}</span>, <span>{this.state.country}</span>.
          </p>
        }
        <button onClick={this.onGetLocationClicked}>
          Get location
        </button>
      </div>
    );
  }
}

export default App;
