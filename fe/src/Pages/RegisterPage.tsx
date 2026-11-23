import { useMutation, useQueryClient } from "@tanstack/react-query";
import { type User } from "firebase/auth";

import { createUser } from "../APIs/usersAPI";
import { logOut } from "../Auth/firebase";
import useAuthUser from "../Auth/useAuthUser";
import { USER_INFO_QUERY_KEY } from "../Utils/queryConsts";

// The username is Google's, not the user's to invent. It is the same name the
// id token carries in its `name` claim, so the row and the JWT the backend
// verifies agree from the first request. Google always sets displayName; the
// email local part is only there for a provider that does not.
export function firebaseUsername(user: User | null): string {
    const name = user?.displayName?.trim();

    if (name)
        return name;

    const email = user?.email?.trim();

    return email ? email.split('@')[0] : '';
}

function RegisterPage() {
    const { user } = useAuthUser();
    const queryClient = useQueryClient();

    const username = firebaseUsername(user);

    const { mutate, isPending } = useMutation({
        mutationFn: () => createUser({ username }),
        onSuccess: () => queryClient.invalidateQueries({ queryKey: [USER_INFO_QUERY_KEY] }),
    });

    return (
        <div className='card shadow-sm p-4'>
            <h1 className='h4'>Finish signing up</h1>
            <p className='text-secondary small'>
                Your posts are signed with the name on your Google account. You can change it later
                from your profile.
            </p>

            <div className='d-flex align-items-center gap-3 mb-3'>
                {!!user?.photoURL && (
                    <img className='rounded-circle' src={user.photoURL} alt='' width={48} height={48}/>
                )}

                <div>
                    <div className='fw-semibold'>{username || 'No name on this account'}</div>
                    {!!user?.email && <div className='text-secondary small'>{user.email}</div>}
                </div>
            </div>

            <button className='btn btn-primary' onClick={() => mutate()} disabled={isPending || !username}>
                {isPending ? 'Creating…' : 'Create account'}
            </button>

            <button className='btn btn-link text-secondary' onClick={() => logOut()}>
                Use a different Google account
            </button>
        </div>
    )
}

export default RegisterPage;
