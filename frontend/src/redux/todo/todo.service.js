import { config } from '../../config'
import { handleResponse } from '../../helpers/handleResponse'

export const todoService = {
    get,
    post,
    put
};

const TODO_URI = '/todo';

/**
 * list all tasks
 */
function get() {
    const requestOptions = {
        method: 'GET'
    };
    return fetch(`${config.apiUrl}${TODO_URI}`, requestOptions).then(handleResponse);
}

function post(todo) {
    const requestOptions = {
        method: 'POST',
        body: JSON.stringify(todo)
    };
    return fetch(`${config.apiUrl}${TODO_URI}`, requestOptions).then(handleResponse);
}

function put(todo) {
    const requestOptions = {
        method: 'PUT',
        body: JSON.stringify(todo)
    };
    return fetch(`${config.apiUrl}${TODO_URI}/${todo.ID}`, requestOptions).then(handleResponse);
}