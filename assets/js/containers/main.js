import React, {Component} from 'react'
import Logo from "../components/logo";
import {connect} from "react-redux";

class Main extends Component {
    render() {
        return (
            <div>
                <Logo/>
            </div>
        )
    }
}

const mapSateToProps = state => {
    return {
        main: state.main
    }
};

export default connect(mapSateToProps)(Main)