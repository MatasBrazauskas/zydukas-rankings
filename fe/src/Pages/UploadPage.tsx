import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";

import uploadPost from "../APIs/postsAPI";
import { CAPTION_MAX_LENGTH, FEED_QUERY_KEY } from "../Utils/queryConsts";

const MAX_FILE_SIZE = 50 << 20; // PostsController.CreatePost rejects anything bigger.

function UploadPage() {
    const [file, setFile] = useState<File | null>(null);
    const [caption, setCaption] = useState('');
    const queryClient = useQueryClient();
    const navigate = useNavigate();

    const { mutate, isPending, reset } = useMutation({
        mutationFn: () => uploadPost(file!, caption.trim()),
        // Straight back to the feed on a successful upload; the invalidation is
        // left unawaited so the feed refetches while it is already on screen.
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: [FEED_QUERY_KEY] });
            navigate('/');
        },
    });

    const tooLarge = !!file && file.size > MAX_FILE_SIZE;

    return (
        <div className='card shadow-sm p-4'>
            <h1 className='h4'>Create a post</h1>

            <textarea
                className='form-control'
                rows={1}
                value={caption}
                maxLength={CAPTION_MAX_LENGTH}
                placeholder='Write a caption (optional)'
                disabled={isPending}
                onChange={(e) => {
                    setCaption(e.target.value);
                    reset();
                }}
            />

            <div className='form-text mb-3 text-end'>
                {caption.length}/{CAPTION_MAX_LENGTH}
            </div>

            {!file &&
                <input
                    type='file'
                    className='form-control mb-3'
                    accept='image/*'
                    disabled={isPending}
                    onChange={(e) => {
                        setFile(e.target.files?.[0] ?? null);
                        reset();
                    }}
                />
            }

            {file && !tooLarge && <img className='img-fluid rounded mb-3' src={URL.createObjectURL(file)} alt={file.name}/>}

            {tooLarge && <div className='alert alert-warning py-2 small'>That file is over 50 MB.</div>}

            <button className='btn btn-primary' onClick={() => mutate()} disabled={isPending || !file || tooLarge}>
                {isPending ? 'Posting…' : 'Post'}
            </button>

            <button className='btn btn-link text-secondary' onClick={() => navigate('/')}>
                Back to the feed
            </button>
        </div>
    )
}

export default UploadPage;
