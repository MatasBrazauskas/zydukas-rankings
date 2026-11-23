import api from "./client";
import { POSTS_CONTROLLER_URL, UPLOAD_REQUEST, HTTP_CODES } from "./const";
import { ApiError } from "../Utils/exceptionHandler";
import { type PostDTO, type UploadedFileDTO } from "../Utils/DTOs";

export async function getFeed(): Promise<PostDTO[]> {
    try {
        const response = await api.get<{ data: PostDTO[] | null }>(POSTS_CONTROLLER_URL);
        return response.data.data ?? [];
    } catch (error) {
        // `GET /posts` is not routed yet (be/app/app.go only has the upload), so
        // an empty feed is not a toast — same deal as getUser on a missing row.
        if (error instanceof ApiError &&
            (error.status === HTTP_CODES.NOT_FOUND || error.status === HTTP_CODES.SERVER_ERROR))
            return [];

        throw error;
    }
}

// The caption rides along in the same multipart body as the file, under
// `caption`. PostsController.CreatePost only reads `file` today, so it is
// dropped server-side until a models.Post exists to store it on.
async function uploadPost(file: File, caption: string): Promise<UploadedFileDTO> {
    const form = new FormData();

    form.append('file', file);
    form.append('caption', caption);

    const response = await api.post<UploadedFileDTO>(UPLOAD_REQUEST, form);
    return response.data;
}

export default uploadPost;
