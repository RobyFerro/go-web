import React, {Component} from 'react'

class Logo extends Component {
    render() {
        return (
            <div className="container">
                <h1>Welcome to Go-Web framework!</h1>
                <img className="logo" src="public/img/logo.png" alt="Go-Web Framework"/>
            </div>
        )
    }
}

export default Logo