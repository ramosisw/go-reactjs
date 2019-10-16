import { createStore, applyMiddleware } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { createLogger } from 'redux-logger';
import rootReducer from './reducers';
import { config } from '../config';

const loggerMiddleware = config.development ? createLogger() : null;

export const store = createStore(
    rootReducer,
    config.development ? applyMiddleware(thunkMiddleware, loggerMiddleware) : applyMiddleware(thunkMiddleware)
);