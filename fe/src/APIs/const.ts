import { getToken } from "../Auth/firebase";

export const USERS_CONTROLLER_URL: string = import.meta.env.VITE_USERS_CONTROLLER;
export const POSTS_CONTROLLER_URL: string = import.meta.env.VITE_POSTS_CONTROLLER;

export const UPLOAD_REQUEST = `${POSTS_CONTROLLER_URL}/upload`;

export enum HTTP_CODES {
    OK = 200,
    CREATED = 201,
    BAD_REQUEST = 400,
    UNAUTHORIZED = 401,
    NOT_FOUND = 404,
    TOO_LARGE = 413,
    SERVER_ERROR = 500,
}

export async function authHeader(): Promise<Record<string, string>> {
    const token = await getToken();
    return token ? { Authorization: `Bearer ${token}` } : {};
}
