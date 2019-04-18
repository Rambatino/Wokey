import React, { Component } from 'react';

export default class BaseCard extends Component {
  render() {
    console.log(this.props.children)
    return (
      <div className="BaseCard" >
        <div className="alert alert-primary" role="alert" >
          This is a primary alert—check it out!
          </div>
        Base Card
        {this.props.children}
        <div className="card" style={{ width: '18rem;' }}>
          <div className="card-body">
            <h5 className="card-title">Card title</h5>
            <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
            <a href="#" className="btn btn-primary">Go somewhere</a>
          </div>
        </div>
      </div >
    );
  }
}
