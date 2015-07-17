"use strict";

import React from "react";

let Item = React.createClass({
    // See Polymer layout attributes
    propTypes: {
        flex: React.PropTypes.oneOfType([
            React.PropTypes.bool,
            React.PropTypes.string
        ]),
        layout: React.PropTypes.bool,
        wrap: React.PropTypes.bool,
        reverse: React.PropTypes.bool,
        horizontal: React.PropTypes.bool,
        vertical: React.PropTypes.bool,
        center: React.PropTypes.bool,
        start: React.PropTypes.bool,
        end: React.PropTypes.bool,
        stretch: React.PropTypes.bool,
        startJustified: React.PropTypes.bool,
        centerJustified: React.PropTypes.bool,
        endJustified: React.PropTypes.bool,
        justified: React.PropTypes.bool,
        aroundJustified: React.PropTypes.bool,
        selfStart: React.PropTypes.bool,
        selfCenter: React.PropTypes.bool,
        selfEnd: React.PropTypes.bool,
        selfStretch: React.PropTypes.bool,
        relative: React.PropTypes.bool,
        fit: React.PropTypes.bool,
        hidden: React.PropTypes.bool
    },

    render() {
        let props = this.props;
        let style = props.layout ? {display: "flex"} : {};
        // flex
        if (typeof(props.flex) === "string") {
            style.flex = props.flex;
        } else if (props.flex) {
            style.flex = "1 1 1e-9px";
        }
        // flex-wrap
        if (props.wrap) {
            style.flexWrap = "wrap";
        }
        // flex-direction
        if (props.vertical) {
            style.flexDirection = style.WebkitFlexDirection = props.reverse ? "column-reverse" : "column";
        } else {
            style.flexDirection = style.WebkitFlexDirection = props.reverse ? "row-reverse" : "row";
        }
        // align-items
        if (props.center) {
            style.alignItems = "center";
        } else if (props.start) {
            style.alignItems = "flex-start";
        } else if (props.end) {
            style.alignItems = "flex-end";
        } else if (props.stretch) {
            style.alignItems = "stretch";
        }
        // justify-content
        if (props.startJustified) {
            style.justifyContent = "flex-start";
        } else if (props.centerJustified) {
            style.justifyContent = "center";
        } else if (props.endJustified) {
            style.justifyContent = "flex-end";
        } else if (props.justified) {
            style.justifyContent = "space-between";
        } else if (props.aroundJustified) {
            style.justifyContent = "space-around";
        }
        // align-self
        if (props.selfStart) {
            style.alignSelf = "flex-start";
        } else if (props.selfCenter) {
            style.alignSelf = "center";
        } else if (props.selfEnd) {
            style.alignSelf = "flex-end";
        } else if (props.selfStretch) {
            style.alignSelf = "stretch";
        }
        // other
        if (props.relative) {
            style.position = "relative";
        } else if (props.fit) {
            style.position = "absolute";
            style.top = style.bottom = style.left = style.right = 0;
        }
        if (props.hidden) {
            style.display = "none";
        }
        
        for (let k in props.style) style[k] = props.style[k];
        return <div {...props} style={style}>{props.children}</div>;
    }
});

let Layout = React.createClass({
    render() {
        return <Item layout {...this.props}>{this.props.children}</Item>;
    }
});

export default {
    Layout: Layout,
    Item: Item
};
