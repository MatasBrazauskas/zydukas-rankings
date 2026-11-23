import { useState } from "react";
import { useDispatch } from "react-redux";
import { useQueryClient } from "@tanstack/react-query";

import { logIn, isPopupCancelled } from "../Auth/firebase";
import { getUser } from "../APIs/usersAPI";
import { ApiError } from "../Utils/exceptionHandler";
import { addError } from "../Store/errorState";
import { USER_INFO_QUERY_KEY } from "../Utils/queryConsts";

type Phase = 'idle' | 'firebase' | 'profile';

function LoginPage() {
    const dispatch = useDispatch();
    const queryClient = useQueryClient();
    const [phase, setPhase] = useState<Phase>('idle');

    async function signIn() {
        setPhase('firebase');

        try {
            // Firebase first: the popup is what mints the id token every
            // backend call needs, and its uid is the firebase_id the users row
            // is keyed by.
            const { user } = await logIn();

            // Then the row, for the username. It lands under the key the shell
            // reads, so the app comes up with the name already in hand instead
            // of flashing a second loading state.
            setPhase('profile');
            const profile = await getUser();
            queryClient.setQueryData([USER_INFO_QUERY_KEY, user.uid], profile);
        } catch (error) {
            // A failed lookup is left to the shell's own query: it runs the
            // moment the session lands, and its toast would double this one.
            if (!isPopupCancelled(error) && !(error instanceof ApiError))
                dispatch(addError('Google sign-in failed.'));
        }

        setPhase('idle');
    }

    return (
        <div className='card shadow-sm p-4 text-center'>
            <h1 className='h3'>zydukas</h1>
            <p className='text-secondary'>Sign in to upload pictures and see posts</p>

            <button className='btn btn-primary' onClick={signIn} disabled={phase !== 'idle'}>
                {phase === 'firebase' && 'Opening Google…'}
                {phase === 'profile' && 'Loading your profile…'}
                {phase === 'idle' && 'Continue with Google'}
            </button>
        </div>
    )
}

export default LoginPage;
