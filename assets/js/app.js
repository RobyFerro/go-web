import React from 'react'
import ReactDOM from 'react-dom'
import {Provider} from 'react-redux'
import store from './store.js'
import Main from "./containers/main";

class App extends React.Component {
    render() {
        return (
            <Provider store={store}>
                <Main/>
            </Provider>
        );
    }
}

ReactDOM.render(<App/>, document.getElementById('app'));