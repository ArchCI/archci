"use strict";

Object.defineProperty(exports, "__esModule", {
    value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var _react = require("react");

var _react2 = _interopRequireDefault(_react);

var Item = _react2["default"].createClass({
    displayName: "Item",

    // See Polymer layout attributes
    propTypes: {
        flex: _react2["default"].PropTypes.oneOfType([_react2["default"].PropTypes.bool, _react2["default"].PropTypes.string]),
        layout: _react2["default"].PropTypes.bool,
        wrap: _react2["default"].PropTypes.bool,
        reverse: _react2["default"].PropTypes.bool,
        horizontal: _react2["default"].PropTypes.bool,
        vertical: _react2["default"].PropTypes.bool,
        center: _react2["default"].PropTypes.bool,
        start: _react2["default"].PropTypes.bool,
        end: _react2["default"].PropTypes.bool,
        stretch: _react2["default"].PropTypes.bool,
        startJustified: _react2["default"].PropTypes.bool,
        centerJustified: _react2["default"].PropTypes.bool,
        endJustified: _react2["default"].PropTypes.bool,
        justified: _react2["default"].PropTypes.bool,
        aroundJustified: _react2["default"].PropTypes.bool,
        selfStart: _react2["default"].PropTypes.bool,
        selfCenter: _react2["default"].PropTypes.bool,
        selfEnd: _react2["default"].PropTypes.bool,
        selfStretch: _react2["default"].PropTypes.bool,
        relative: _react2["default"].PropTypes.bool,
        fit: _react2["default"].PropTypes.bool,
        hidden: _react2["default"].PropTypes.bool
    },

    render: function render() {
        var props = this.props;
        var style = props.layout ? { display: "flex" } : {};
        // flex
        if (typeof props.flex === "string") {
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

        for (var k in props.style) {
            style[k] = props.style[k];
        }return _react2["default"].createElement(
            "div",
            _extends({}, props, { style: style }),
            props.children
        );
    }
});

var Layout = _react2["default"].createClass({
    displayName: "Layout",

    render: function render() {
        return _react2["default"].createElement(
            Item,
            _extends({ layout: true }, this.props),
            this.props.children
        );
    }
});

exports["default"] = {
    Layout: Layout,
    Item: Item
};
module.exports = exports["default"];
