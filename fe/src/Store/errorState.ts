import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import { ERRORS_INFO, type ErrorsInformation } from "./utils";

const initialState: ErrorsInformation = {
    errors: [],
}

const MAX_ERRORS = 3;

const errorSlice = createSlice({
    name: ERRORS_INFO,
    initialState,
    reducers: {
        addError: (state: ErrorsInformation, action: PayloadAction<string>) => {
            if (state.errors.includes(action.payload))
                return;

            state.errors.unshift(action.payload);
            state.errors.length = Math.min(state.errors.length, MAX_ERRORS);
        },
        removeError: (state: ErrorsInformation, action: PayloadAction<number>) => {
            state.errors.splice(action.payload, 1);
        },
    }
})

export const { addError, removeError } = errorSlice.actions;
export default errorSlice.reducer;
