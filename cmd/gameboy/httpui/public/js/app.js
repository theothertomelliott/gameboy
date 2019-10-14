'use strict';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      cpu: {
        Paused: false,
        Registers: {
          PC: 0
        }
      }
    };
  }

  componentDidMount() {
    this.timer = setInterval(() => this.getCPU(), 1000);
    this.getCPU();
  }

  componentWillUnmount() {
    this.timer = null;
  }

  getCPU() {
    fetch('/api/cpu')
      .then(result => {
        if (!result.ok) {
          throw ("error loading cpu state");
        }
        return result.json()
      })
      .then(result => this.setState({ cpu: result }))
      .catch(error => {
        console.log(error);
        clearInterval(this.timer)
        this.timer = null;
      });
  }


  render() {
    return (
      <div>
        <DebugMenu paused={this.state.cpu.Paused}></DebugMenu>
        <div className="debugspacer"></div>
        <div className="rightbar">
          <div className="registers">
            <h2>Registers</h2>
            <Registers
              registers={this.state.cpu.Registers}
            ></Registers>
          </div>
          <div className="serial">
            <h2>Serial Output</h2>

          </div>
          <div className="stack">
            <h2>Stack</h2>
          </div>
        </div>
        <Decompile
          paused={this.state.cpu.Paused}
          pc={this.state.cpu.Registers.PC}
        ></Decompile>
      </div>
    );
  }
}

function x(str) {
  if (!str) {
    return str
  }
  return str.toString(16).toUpperCase();
}

class Registers extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <table border="0">
        <tbody>
          <tr><td>A</td><td>0x{x(this.props.registers.A)}</td><td>F</td><td>0x{x(this.props.registers.F)}</td></tr>
          <tr><td>B</td><td>0x{x(this.props.registers.B)}</td><td>C</td><td>0x{x(this.props.registers.C)}</td></tr>
          <tr><td>D</td><td>0x{x(this.props.registers.D)}</td><td>E</td><td>0x{x(this.props.registers.E)}</td></tr>
          <tr><td>H</td><td>0x{x(this.props.registers.H)}</td><td>L</td><td>0x{x(this.props.registers.L)}</td></tr>
          <tr><td colSpan="2">PC</td><td colSpan="2">0x{x(this.props.registers.PC)}</td></tr>
          <tr><td colSpan="2">SP</td><td colSpan="2">0x{x(this.props.registers.SP)}</td></tr>
        </tbody>
      </table>
    );
  }
}

class DebugMenu extends React.Component {
  constructor(props) {
    super(props);
  }

  performAction(path) {
    fetch(path).then(result => {
      if (!result.ok) {
        throw ("error calling" + path);
      }
    })
      .catch(error => {
        console.log(error);
      });
  }

  render() {
    var paused = "Pause";
    if (this.props.paused) {
      paused = "Resume";
    }
    return (
      <div className="debugmenu">
        <ul>
          <li><button onClick={e => { this.performAction("/debug/togglepaused") }}>{paused}</button></li>
          <li><button onClick={e => { this.performAction("/debug/step") }}>Step</button></li>
          <li><button onClick={e => { this.performAction("/reset") }}>Reset</button></li>
        </ul>
      </div>
    );
  }
}

class Decompile extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      decompile: [],
      breakpoints: [],
      error: null
    };
  }

  componentDidMount() {
    this.timer = setInterval(() => this.updateData(), 1000);
    this.updateData();
  }

  componentWillUnmount() {
    this.timer = null;
  }

  updateData() {
    this.getDecompilation();
    this.getBreakpoints();
  }

  getDecompilation() {
    fetch('/api/decompile')
      .then(result => {
        if (!result.ok) {
          throw ("error loading decompile");
        }
        return result.json()
      })
      .then(result => this.setState({ decompile: result }))
      .catch(error => {
        console.log(error);
        clearInterval(this.timer)
        this.timer = null;
      });
  }

  getBreakpoints() {
    fetch('/api/breakpoints')
      .then(result => {
        if (!result.ok) {
          throw ("error loading breakpoints");
        }
        return result.json()
      })
      .then(result => this.setState({ breakpoints: result }))
      .catch(error => {
        // Degrade gracefully
        this.setState({ breakpoints: [] })
      });
  }

  setBreakpoint(index) {
    if (!index) {
      return;
    }
    let indexUC = index.toString(16).toUpperCase();
    fetch('/debug/togglebreakpoint/' + indexUC).then(result => {
      if (!result.ok) {
        throw ("error toggling breakpoint for " + indexUC);
      }
      console.log("Toggled breakpoint at " + indexUC);
      this.getBreakpoints();
    })
      .catch(error => {
        console.log(error);
      });
  }

  componentWillReceiveProps(nextProps) {
    // You don't have to do this check first, but it can help prevent an unneeded render
    if (nextProps.pc !== this.state.pc) {
      this.setState({ pc: nextProps.pc });
      if (nextProps.paused) {
        this.setState({ shouldScroll: true });
      }
    }
  }

  componentDidUpdate() {
    if (this.state.shouldScroll) {
      scrollToPC();
      this.setState({ shouldScroll: false });
    }
  }

  render() {
    return (
      <table border="0" cellPadding="0" cellSpacing="0">
        <tbody>
          {this.state.decompile.map((value, index) => {
            var bp = <td>&nbsp;</td>;
            if (this.state.breakpoints && this.state.breakpoints.includes(value.Index)) {
              bp = <td>•</td>;
            }
            let hexIndex = value.Index.toString(16).toUpperCase();
            var identifier = hexIndex;
            if (this.props.pc == value.Index) {
              identifier = 'PC';
            }
            return (
              <tr id={identifier} key={value.Index} onClick={(e) => this.setBreakpoint(value.Index)} className="oprow">
                {bp}
                <td>&nbsp;</td>
                <td>0x{hexIndex}</td>
                <td>&nbsp;</td>
                <td>{value.Description}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    );
  }
}

const domContainer = document.querySelector('#app_container');
ReactDOM.render(React.createElement(App), domContainer);