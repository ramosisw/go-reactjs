import { todoConstants } from './todo.constants'
import { todoService } from './todo.service'


export const todoActions = {
    get,
    post,
    put
};

function get() {
    return dispatch => {
        dispatch(request());
        todoService.get().then(
            data => {
                dispatch(success(data));
            },
            error => {
                dispatch(failure(error.toString()));
            }
        );
    };

    function request() { return { type: todoConstants.GET_REQUEST } }
    function success(data) { return { type: todoConstants.GET_SUCCESS, data } }
    function failure(error) { return { type: todoConstants.GET_FAILURE, error } }
}

function post(todo) {
    return dispatch => {
        dispatch(request());
        todoService.post(todo).then(
            data => {
                dispatch(success(data));
            },
            error => {
                dispatch(failure(error.toString()));
            }
        );
    };

    function request() { return { type: todoConstants.POST_REQUEST } }
    function success(data) { return { type: todoConstants.POST_SUCCESS, data } }
    function failure(error) { return { type: todoConstants.POST_FAILURE, error } }
}

function put(todo) {
    return dispatch => {
        dispatch(request());
        todoService.put(todo).then(
            data => {
                dispatch(success(data));
            },
            error => {
                dispatch(failure(error.toString()));
            }
        );
    };

    function request() { return { type: todoConstants.PUT_REQUEST } }
    function success(data) { return { type: todoConstants.PUT_SUCCESS, data } }
    function failure(error) { return { type: todoConstants.PUT_FAILURE, error } }
}