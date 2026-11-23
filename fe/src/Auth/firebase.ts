import { initializeApp } from "firebase/app";
import {
    GoogleAuthProvider,
    deleteUser,
    getAuth,
    onAuthStateChanged,
    reauthenticateWithPopup,
    signInWithPopup,
    signOut,
    type User,
} from "firebase/auth";

const app = initializeApp({
    apiKey: import.meta.env.VITE_FIREBASE_API_KEY,
    authDomain: import.meta.env.VITE_FIREBASE_AUTH_DOMAIN,
    projectId: import.meta.env.VITE_FIREBASE_PROJECT_ID,
    storageBucket: import.meta.env.VITE_FIREBASE_STORAGE_BUCKET,
    messagingSenderId: import.meta.env.VITE_FIREBASE_MESSAGING_SENDER_ID,
    appId: import.meta.env.VITE_FIREBASE_APP_ID,
});

export const auth = getAuth(app);

const provider = new GoogleAuthProvider();

export const logIn = () => signInWithPopup(auth, provider);

export const logOut = () => signOut(auth);

export const onAuthChange = (listener: (user: User | null) => void) => onAuthStateChanged(auth, listener);

export async function getToken(): Promise<string | null> {
    return auth.currentUser ? auth.currentUser.getIdToken() : null;
}

// Firebase refuses deleteUser on an old session, so ask for the popup once.
export async function deleteAuthUser() {
    const user = auth.currentUser;
    
    if (!user)
        return;

    try {
        await deleteUser(user);
    } catch (error) {
        if ((error as { code?: string }).code !== 'auth/requires-recent-login')
            throw error;

        await reauthenticateWithPopup(user, provider);
        await deleteUser(user);
    }
}

export const isPopupCancelled = (error: unknown) =>
    ['auth/popup-closed-by-user', 'auth/cancelled-popup-request'].includes((error as { code?: string }).code ?? '');
