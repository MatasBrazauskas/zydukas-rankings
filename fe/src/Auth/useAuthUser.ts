import { useEffect, useState } from "react";
import { type User } from "firebase/auth";

import { onAuthChange } from "./firebase";

type AuthState = {
    user: User | null,
    loading: boolean,
}

function useAuthUser(): AuthState {
    const [state, setState] = useState<AuthState>({ user: null, loading: true });

    useEffect(() => onAuthChange((user) => setState({ user, loading: false })), []);

    return state;
}

export default useAuthUser;
