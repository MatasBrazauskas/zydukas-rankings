import { useState } from "react";
import { useOutletContext } from "react-router-dom";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { deleteUser, updateUsername } from "../APIs/usersAPI";
import { deleteAuthUser, logOut } from "../Auth/firebase";
import { type UserInfoDTO } from "../Utils/DTOs";
import {
    USERNAME_MIN_LENGTH,
    USERNAME_MAX_LENGTH,
    USER_INFO_QUERY_KEY,
} from "../Utils/queryConsts";

// Go sends time.Time as RFC3339; anything else is not worth guessing at.
function formatDate(value: string): string {
    const date = new Date(value);
    return Number.isNaN(date.getTime()) ? '—' : date.toLocaleDateString();
}

function ProfilePage() {
    // Everything shown here is the `GET /users` row from the Outlet, not the
    // Firebase session — the session's name, picture and email are Google's
    // copy and drift from the database.
    const profile = useOutletContext<UserInfoDTO>();
    const queryClient = useQueryClient();

    const [username, setUsername] = useState(profile.username);

    const rename = useMutation({
        mutationFn: () => updateUsername(username.trim()),
        onSuccess: () => queryClient.invalidateQueries({ queryKey: [USER_INFO_QUERY_KEY] }),
    });

    // The row goes first: if Firebase then refuses, the session is still good
    // enough to retry with, while the other order would orphan the row.
    const remove = useMutation({
        mutationFn: async () => {
            await deleteUser();
            await deleteAuthUser();
        },
        onSettled: () => queryClient.invalidateQueries({ queryKey: [USER_INFO_QUERY_KEY] }),
        onError: () => logOut(),
    });

    const trimmed = username.trim();
    const tooShort = trimmed.length < USERNAME_MIN_LENGTH;
    const unchanged = trimmed === profile.username;

    function confirmRemove() {
        if (window.confirm(`Delete ${profile.username}? This cannot be undone.`))
            remove.mutate();
    }

    return (
        <>
            <div className='card shadow-sm overflow-hidden mb-3'>
                <div className='profileCover'/>

                <div className='card-body text-center pt-0'>
                    <div className='profileAvatar'>{profile.username.charAt(0).toUpperCase()}</div>

                    <h1 className='h4 mt-3 mb-1'>{profile.username}</h1>
                    <p className='text-secondary small mb-0'>Joined {formatDate(profile.created_at)}</p>

                    <button className='btn btn-outline-primary btn-sm mt-3' onClick={() => logOut()}>
                        Sign out
                    </button>
                </div>
            </div>

            <div className='card shadow-sm p-4 mb-3'>
                <h2 className='h6'>Username</h2>
                <p className='text-secondary small mb-3'>
                    This is the name shown on your posts. {USERNAME_MIN_LENGTH}–{USERNAME_MAX_LENGTH} characters.
                </p>

                <div className='d-flex gap-2'>
                    <input
                        className='form-control'
                        value={username}
                        maxLength={USERNAME_MAX_LENGTH}
                        disabled={rename.isPending}
                        onChange={(e) => {
                            setUsername(e.target.value);
                            rename.reset();
                        }}
                    />

                    <button
                        className='btn btn-primary flex-shrink-0'
                        onClick={() => rename.mutate()}
                        disabled={rename.isPending || tooShort || unchanged}
                    >
                        {rename.isPending ? 'Saving…' : 'Save'}
                    </button>
                </div>

                {(tooShort && !!trimmed.length) && (
                    <div className='form-text'>At least {USERNAME_MIN_LENGTH} characters.</div>
                )}
                {rename.isSuccess && <div className='alert alert-success py-2 small mb-0 mt-3'>Username updated.</div>}
            </div>

            <div className='card shadow-sm px-4 py-2 mb-3'>
                <div className='profileRow border-top-0'>
                    <span className='text-secondary small'>Joined</span>
                    <span className='fw-semibold'>{formatDate(profile.created_at)}</span>
                </div>

                <div className='profileRow'>
                    <span className='text-secondary small'>Last updated</span>
                    <span className='fw-semibold'>{formatDate(profile.updated_at)}</span>
                </div>
            </div>

            <div className='card shadow-sm p-4'>
                <h2 className='h6 text-danger'>Danger zone</h2>
                <p className='text-secondary small mb-3'>
                    Deleting your account removes your profile and your sign-in. This cannot be undone.
                </p>

                <button className='btn btn-outline-danger' onClick={confirmRemove} disabled={remove.isPending}>
                    {remove.isPending ? 'Deleting…' : 'Delete my account'}
                </button>
            </div>
        </>
    )
}

export default ProfilePage;
