require("./actions/Mockjax.js");

import React from "react";
import {Layout} from "react-polymer-layout";
import Router from "react-router";
import Builds from "./components/Builds.jsx";
import Images from "./components/Images.jsx";
import Projects from "./components/Projects.jsx";
import Workers from "./components/Workers.jsx";
import {Colors} from "./components/Styles.jsx";

const {DefaultRoute, Link, Route, RouteHandler} = Router;

const App = React.createClass({
  render() {
    let tabStyle = {
      height: 48,
      lineHeight: "48px",
      padding: "0 20px",
      fontSize: "1em",
      color: "#eee",
      display: "inline-block"
    };
    let tabs = [
      <div style={tabStyle} className="tab-link" key="logo">
        <span style={{fontWeight: 600, fontSize: "1.2em", color: "white"}}>ArchCI</span>
      </div>
    ];
    tabs.push(["Builds", "Projects", "Images", "Workers"].map(n =>
        <Link to={n.toLowerCase()} style={tabStyle} className="tab-link" key={n}>
          {n}
        </Link>
    ));

    return (
      <Layout vertical style={{height: "100%"}}>
        <Layout horizontal style={{background: Colors.GREEN_700, height: 48, width: "100%"}}>
          {tabs}
        </Layout>
        
        <RouteHandler />
      </Layout>
    );
  }
});

let routes = (
  <Route name="app" path="/" handler={App}>
    <Route name="builds" handler={Builds} />
    <Route name="images" handler={Images} />
    <Route name="projects" handler={Projects} />
    <Route name="workers" handler={Workers} />
    <DefaultRoute handler={Builds} />
  </Route>
);

Router.run(routes, Handler => {
  React.render(<Handler />, document.body);
});
