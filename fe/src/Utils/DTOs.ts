export type UserInfoDTO = {
    username: string,
    firebase_id: string,
    created_at: string,
    updated_at: string,
}

export type CreateUserDTO = {
    username: string,
}

export type UploadedFileDTO = {
    message: string,
}

export type PostDTO = {
    id: number,
    username: string,
    image_url: string,
    created_at: string,
    // Optional: a post can be posted without one, and the Go side does not
    // send the field at all yet.
    caption?: string,
}
