import React from "react/addons";
import AppActions from "../actions/AppActions.js";
import Input from "react-bootstrap/lib/Input";
import {Layout, Item} from "react-polymer-layout";
import {Colors} from "./Styles.jsx";
import moment from "moment";

const Builds = React.createClass({
  mixins: [React.addons.LinkedStateMixin],

  getInitialState() {
    return {
      builds: [],
      selected: null,
      timer: null,
      search: "",
      log: "",
      last: 0
    };
  },

  componentDidMount() {
    let that = this;
    this.state.timer = setInterval(() => {
      that._receive_builds();
      that._receive_log();
    }, 1000);
  },

  componentWillUnmount() {
    clearInterval(this.state.timer);
  },

  _receive_builds(flag) {
    let search = this.state.search.trim();
    if (search && !flag) return;
    let params = search ? {search: search} : {};
    AppActions.receive_builds(params, data => {
      let found = false;
      for (let i = 0; i < data.length; i++) {
        if (data[i].id === this.state.selected) {
          found = true;
        }
      }
      if (!found) {
        this.state.selected = data[0] ? data[0].id : null;
        this.state.last = 0;
        this.state.log = "";
      }
      this.setState({builds: data});
    });
  },

  _receive_log() {
    let build = this.state.builds.filter(b => b.id === this.state.selected)[0];
    if (build && build.status !== null) return;
    AppActions.receive_log({
      id: this.state.selected,
      last: this.state.last
    }, data => {
      if (this.state.last === 0) {
        this.setState({
          log: data.log,
          last: data.last
        });
      } else {
        this.setState({
          log: this.state.log + data.log,
          last: data.last
        });
      }
    });
  },

  _handleClick(id) {
    if (id === this.state.selected) return;
    this.setState({
      selected: id,
      last: 0,
      log: ""
    });
  },

  _getBuildList(builds) {
    return builds.map(b => {
      return (
        <Layout vertical style={{padding: 10, borderLeft: `5px solid ${Colors.GREEN_700}`, marginBottom: 10}} onClick={() => this._handleClick(b.id)}>
          <Item style={{fontSize: "1.1em", color: Colors.GREEN_900}}>{`${b.owner}/${b.repository}`}</Item>
          <Item style={{color: "#999"}}>Commit: {b.commit}</Item>
          <Item style={{color: "#999"}}>
            Elapsed: {moment.duration((b.finished && new Date(b.finished).getTime()) || Date.now() - new Date(b.dispatched).getTime()).humanize()}
          </Item>
        </Layout>
      );
    });
  },

  render() {
    if (this.state.selected === null && this.state.builds[0]) {
      this.state.selected = this.state.builds[0].id;
    }
    let build = this.state.builds.filter(b => b.id === this.state.selected)[0];
    let rightPanel = null;
    if (build) {
      rightPanel = (
        <Layout vertical flex style={{paddingLeft: 20}}>
          <div style={{fontSize: "1.6em"}}>{build && `${build.owner}/${build.repository}`}</div>
          <Layout horizontal style={{color: "#999", marginTop: 20}} wrap>
            <Layout vertical flex="1" style={{minWidth: 250}}>
              <Item>Branch: {build.branch}</Item>
              <Item>Commit: {build.commit}</Item>
              <Item>Committer: {build.committer}</Item>
            </Layout>
            <Layout vertical flex="1" style={{minWidth: 250}}>
              <Item>Build: {build.id}</Item>
              <Item>Worker: {build.worker}</Item>
              <Item>Elapsed: {moment.duration((build.finished && new Date(build.finished).getTime()) || Date.now() - new Date(build.dispatched).getTime()).humanize()}</Item>
            </Layout>
          </Layout>
          <pre style={{marginTop: 20, padding: 10, background: "#000", color: "#fafafa"}}>
            <code>
              {this.state.log}
            </code>
          </pre>
        </Layout>
      );
    }
    return (
      <Layout horizontal style={{padding: 20}}>
        <Layout vertical style={{width: 250}}>
          <Layout vertical style={{border: `1px solid #ccc`}}>
            <Item style={{padding: 10}}>
              <Input type="text" className="sidebar-input" placeholder="Search"
                     valueLink={this.linkState("search")} />
            </Item>
            <Layout vertical>
              {this._getBuildList(this.state.builds)}
            </Layout>
          </Layout>
        </Layout>
        {rightPanel}
      </Layout>
    );
  }
});

export default Builds;
