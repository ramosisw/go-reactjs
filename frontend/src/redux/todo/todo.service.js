import { config } from '../../config'
import { handleResponse } from '../../helpers/handleResponse'

export const todoService = {
    get
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