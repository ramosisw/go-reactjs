import { todoConstants } from './todo.constants'

const initialState = {
    get: {
        data: []
    }
}

export function todo(state = {}, action) {
    switch (action.type) {
        case todoConstants.GET_SUCCESS: return {
            ...state,
            get: {
                ...state.get,
                data: action.data
            }
        }
        default: return initialState;
    }
}