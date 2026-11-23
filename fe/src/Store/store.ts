import { configureStore } from "@reduxjs/toolkit";

import errorsReducer from './errorState';

const store = configureStore({
    reducer: {
        ERRORS_INFO: errorsReducer,
    }
});

export default store;
export type RootState = ReturnType<typeof store.getState>;
