import axios from "axios";

import { getToken } from "../Auth/firebase";
import { ApiError, readMessage } from "../Utils/exceptionHandler";

// Plain instance: the URLs already come from the env (see const.ts), so there
// is no baseURL, and no transforms beyond the two interceptors below.
const api = axios.create();

// FirebaseMiddleware wants `Authorization: Bearer <firebase id token>`.
api.interceptors.request.use(async (config) => {
    const token = await getToken();
    if (token)
        config.headers.Authorization = `Bearer ${token}`;

    return config;
});

// Every failure leaves here as an ApiError, so callers and the toast in
// main.tsx see one error type with a message the Go side actually sent.
api.interceptors.response.use(
    (response) => response,
    (error: unknown) => {
        if (!axios.isAxiosError(error))
            return Promise.reject(error);

        const method = error.config?.method?.toUpperCase() ?? 'request';
        const path = error.config?.url ?? '';

        // No response at all: offline, DNS, CORS, timeout.
        if (!error.response)
            return Promise.reject(new ApiError(0, `${method} ${path} failed (${error.message})`));

        const { status, data } = error.response;
        return Promise.reject(new ApiError(status, readMessage(data) ?? `${method} ${path} failed (${status})`));
    },
);

export default api;
