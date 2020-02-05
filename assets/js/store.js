import {applyMiddleware, combineReducers, createStore} from 'redux';
import mainReducer from "./reducers/mainReducer";
import thunk from 'redux-thunk';

const store = applyMiddleware(thunk)(createStore)(combineReducers({
    main: mainReducer
}), {
    main: {}
});

export default store;