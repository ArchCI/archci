import React from "react";
import Moment from "moment";
import {Layout, Item} from "react-polymer-layout";
import {Colors} from "./Styles.jsx";
import AppActions from "../actions/AppActions.js";

const Workers = React.createClass({
  getInitialState() {
    return {
      workers: []
    };
  },

  componentDidMount() {
    AppActions.receive_workers({}, workers => {
      this.setState({workers: workers});
    });
  },

  _labelStyle(color) {
    return {
      background: color,
      width: 250,
      padding: 10,
      textAlign: "right",
      fontSize: "1em",
      color: "white"
    };
  },

  _workerList(workers) {
    return workers.map(w =>
      <Layout justified style={{padding: 10, margin: "0 0 10px 20px", background: "#fafafa"}}>
        <Item>{w.ip}</Item>
        <Item style={{color: "#999"}}>
          {Moment(new Date(w.contacted)).fromNow()}
        </Item>
      </Layout>
    );
  },

  render() {
    return (
      <Layout vertical style={{padding: 20}}>
        <Layout horizontal start>
          <div style={this._labelStyle(Colors.YELLOW_A700)}>Busy</div>
          <Layout vertical flex>
            {this._workerList(this.state.workers.filter(w => {
              return Date.now() - new Date(w.contacted).getTime() < 86400000 && w.busy;
            }))}
          </Layout>
        </Layout>
        
        <div><hr /></div>

        <Layout horizontal start>
          <div style={this._labelStyle(Colors.GREEN_A700)}>Idle</div>
          <Layout vertical flex>
            {this._workerList(this.state.workers.filter(w => {
              return Date.now() - new Date(w.contacted).getTime() < 86400000 && !w.busy;
            }))}
          </Layout>
        </Layout>

        <div><hr /></div>

        <Layout horizontal start>
          <div style={this._labelStyle(Colors.BLUE_A700)}>Silent</div>
          <Layout vertical flex>
            {this._workerList(this.state.workers.filter(w => {
              return Date.now() - new Date(w.contacted).getTime() > 86400000;
            }))}
          </Layout>
        </Layout>
      </Layout>
    );
  }
});

export default Workers;
