var _createClass = function () {function defineProperties(target, props) {for (var i = 0; i < props.length; i++) {var descriptor = props[i];descriptor.enumerable = descriptor.enumerable || false;descriptor.configurable = true;if ("value" in descriptor) descriptor.writable = true;Object.defineProperty(target, descriptor.key, descriptor);}}return function (Constructor, protoProps, staticProps) {if (protoProps) defineProperties(Constructor.prototype, protoProps);if (staticProps) defineProperties(Constructor, staticProps);return Constructor;};}();function _classCallCheck(instance, Constructor) {if (!(instance instanceof Constructor)) {throw new TypeError("Cannot call a class as a function");}}function _possibleConstructorReturn(self, call) {if (!self) {throw new ReferenceError("this hasn't been initialised - super() hasn't been called");}return call && (typeof call === "object" || typeof call === "function") ? call : self;}function _inherits(subClass, superClass) {if (typeof superClass !== "function" && superClass !== null) {throw new TypeError("Super expression must either be null or a function, not " + typeof superClass);}subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } });if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass;}var App = function (_React$Component) {_inherits(App, _React$Component);
  function App(props) {_classCallCheck(this, App);var _this = _possibleConstructorReturn(this, (App.__proto__ || Object.getPrototypeOf(App)).call(this,
    props));
    _this.state = {
      linkInput: "",
      codeInput: "",
      link: "",
      shareCode: 0,
      alreadyShared: false,
      alreadyGotCode: false };

    _this.submitLink = _this.submitLink.bind(_this);
    _this.submitNumber = _this.submitNumber.bind(_this);return _this;
  }_createClass(App, [{ key: "validateNumInput", value: function validateNumInput(
    num) {
      num = num + "";
      if (num.length > 4) return false;
      if (isNaN(num)) return false;
      return true;
    } }, { key: "handleLinkChange", value: function handleLinkChange(
    event) {
      this.setState({
        linkInput: event.target.value });

    } }, { key: "submitLink", value: function submitLink()
    {
      xmlhttp = new XMLHttpRequest();
      xmlhttp.open("GET", "http://ktj.st/linkshare/addlink/?link=" + this.state.linkInput, false);
      xmlhttp.send();
      var data = xmlhttp.responseText;
      this.setState({
        alreadyShared: true,
        shareCode: data });

    } }, { key: "submitNumber", value: function submitNumber()

    {
      xmlhttp = new XMLHttpRequest();
      xmlhttp.open("GET", "http://ktj.st/linkshare/getcode/?code=" + this.state.codeInput, false);
      xmlhttp.send();
      var data = xmlhttp.responseText;
      this.setState({
        alreadyGotCode: true,
        link: data });

    } }, { key: "handleNumberChange", value: function handleNumberChange(

    event) {
      this.setState({
        codeInput: this.validateNumInput(event.target.value) ?
        event.target.value :
        this.state.codeInput });

    } }, { key: "render", value: function render()
    {
      return (
        React.createElement("div", null,
          React.createElement("div", { className: "chooseSide" },
            React.createElement("div", { className: "accessSide" },
              React.createElement("h1", { className: "title" }, this.state.alreadyGotCode ? "HERE IS YOUR LINK!" : "ACCESS SHARED LINK"),
              !this.state.alreadyGotCode && React.createElement("input", {
                placeholder: "1234",
                onChange: this.handleNumberChange.bind(this),
                value: this.state.codeInput,
                className: "inputBoxNumber" }),

              this.state.alreadyGotCode ? React.createElement("h2", { className: "linkText" }, this.state.link) : React.createElement("button", { onClick: this.submitNumber, className: "submitButton" }, "GET")),

            React.createElement("div", { className: "shareSide" },
              React.createElement("h1", { className: "title" },
                this.state.alreadyShared ? "SHARE THIS CODE" : "INPUT A LINK TO SHARE"),

              !this.state.alreadyShared &&
              React.createElement("input", {
                value: this.state.linkInput,
                onChange: this.handleLinkChange.bind(this),
                className: "inputBox" }),


              this.state.alreadyShared ?
              React.createElement("h2", { className: "shareCode" }, this.state.shareCode) :

              React.createElement("button", { className: "submitButton", onClick: this.submitLink }, "SHARE")))));







    } }]);return App;}(React.Component);


var appContainer = React.createElement(App, null);

ReactDOM.render(appContainer, document.getElementById("app"));
