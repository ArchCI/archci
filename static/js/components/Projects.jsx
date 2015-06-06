import React from "react";
import AppActions from "../actions/AppActions.js";
import AppStores from "../stores/AppStores.js";
import DropdownButton from "react-bootstrap/lib/DropdownButton";
import Button from "react-bootstrap/lib/Button";
import Input from "react-bootstrap/lib/Input";
import MenuItem from "react-bootstrap/lib/MenuItem";
import {Layout, Item} from "react-polymer-layout";
import {Colors} from "./Styles.jsx";
import moment from "moment";

const Projects = React.createClass({
  getInitialState() {
    return {
      projects: [],
      builds: []
    };
  },

  componentDidMount() {
    AppActions.receive_projects({});
    AppStores.addChangeListener(this._onChange);
  },

  componentWillUnmount() {
    AppStores.removeChangeListener(this._onChange);
  },

  _onChange() {
    let projects = AppStores.getProjects();
    this.setState({
      projects: Object.keys(projects).map(k => projects[k])
    });
  },

  _handleClick(id) {
    AppActions.receive_project_builds({id: id}, data => {
      this.setState({
        builds: data
      });
    });
  },

  _getBuildList(builds) {
    return builds.map(b => {
      return (
        <Layout style={{padding: "10px 0"}}>
          <Item flex="1" style={{marginLeft: 20}}>{b.id}</Item>
          <Item flex="3">{b.commit} ({b.branch})</Item>
          <Item flex="2">{b.committer}</Item>
          <Item flex="2">
            {moment.duration(new Date(b.finished).getTime() - new Date(b.dispatched).getTime()).humanize()}</Item>
        </Layout>
      );
    });
  },

  _getProjectList(projects) {
    return projects.map(p => {
      return (
        <Layout vertical style={{padding: 10, borderLeft: `5px solid ${Colors.GREEN_700}`, marginBottom: 10}} onClick={() => this._handleClick(p.id)}>
          <Item style={{fontSize: "1.1em", color: Colors.GREEN_900}}>{`${p.owner}/${p.repository}`}</Item>
          <Layout horizontal justified center>
            <Item style={{color: "#999"}}>Domain: {p.domain}</Item>
            <Item style={{color: "#999"}}>{p.public ? "Public" : "Private"}</Item>
          </Layout>
        </Layout>
      );
    });
  },

  render() {
    return (
      <Layout horizontal style={{padding: 20}}>
        <Layout vertical style={{width: 250}}>
          <Layout vertical style={{border: `1px solid #ccc`, padding: 10}}>
            <div style={{fontWeight: 600, color: "#999"}}>New Project</div>
            <Item>
              <Input type="text" className="sidebar-input" placeholder="Repository URL" ref="url" />
            </Item>
            <Item>
              <DropdownButton title="Repository type" className="sidebar-drop">
                <MenuItem eventKey="1">GitHub (Public)</MenuItem>
                <MenuItem eventKey="2">GitHub (Private)</MenuItem>
                <MenuItem eventKey="3">GitLab (Private)</MenuItem>
              </DropdownButton>
            </Item>
            <Item>
            </Item>
          </Layout>

          <Layout vertical style={{border: `1px solid #ccc`, marginTop: 20}}>
            <Item style={{padding: 10}}>
              <Input type="text" className="sidebar-input" placeholder="Filter by name" ref="filter" />
            </Item>
            <Layout vertical>
              {this._getProjectList(this.state.projects)}
            </Layout>
          </Layout>
        </Layout>

        <Layout style={{marginLeft: 20}} vertical flex>
          <Layout>
            <Item flex="1" style={{marginLeft: 20}}>Build ID</Item>
            <Item flex="3">Commit</Item>
            <Item flex="2">Committer</Item>
            <Item flex="2">Elapsed</Item>
          </Layout>
          <div style={{color: "#999", marginTop: -10}}><hr /></div>
          <Layout vertical>
            {this._getBuildList(this.state.builds)}
          </Layout>
        </Layout>
      </Layout>
    );
  }
});

export default Projects;
