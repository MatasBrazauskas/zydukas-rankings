import api from "./client";
import { USERS_CONTROLLER_URL, HTTP_CODES } from "./const";
import { ApiError } from "../Utils/exceptionHandler";

import { type UserInfoDTO } from "../Utils/DTOs";
import { type CreateUserDTO } from "../Utils/DTOs";

export async function getUser(): Promise<UserInfoDTO | null> {
    try {
        const response = await api.get<{ data: UserInfoDTO }>(USERS_CONTROLLER_URL);
        return response.data.data;
    } catch (error) {
        // No row yet (or the Go side failed to look for one) is not a toast.
        if (error instanceof ApiError &&
            (error.status === HTTP_CODES.NOT_FOUND || error.status === HTTP_CODES.SERVER_ERROR))
            return null;

        throw error;
    }
}


export async function createUser(payload: CreateUserDTO) {
    const response = await api.post(USERS_CONTROLLER_URL, payload);
    return response.data;
}

// PATCH /users/:new_username — the name travels in the path, the firebase id
// in the Authorization header the request interceptor adds.
export async function updateUsername(newUsername: string) {
    const response = await api.patch(`${USERS_CONTROLLER_URL}/${encodeURIComponent(newUsername)}`);
    return response.data;
}

export async function deleteUser() {
    const response = await api.delete(USERS_CONTROLLER_URL);
    return response.data;
}
