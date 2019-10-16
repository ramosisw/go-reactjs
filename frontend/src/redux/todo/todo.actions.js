import { todoConstants } from './todo.constants'
import { todoService} from './todo.service'


export const todoActions = {
    get
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