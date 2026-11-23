import { Outlet } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";

import TopBar from "./Components/TopBar";
import LoginPage from "./Pages/LoginPage";
import RegisterPage from "./Pages/RegisterPage";
import ErrorsComponent from "./Components/ErrorsComponent";
import useAuthUser from "./Auth/useAuthUser";
import { getUser } from "./APIs/usersAPI";
import { USER_INFO_QUERY_KEY, USER_INFO_STALE_TIME } from "./Utils/queryConsts";

function MainPage() {
    const { user, loading } = useAuthUser();

    const { data: profile, isLoading } = useQuery({
        queryKey: [USER_INFO_QUERY_KEY, user?.uid],
        queryFn: getUser,
        enabled: !!user,
        staleTime: USER_INFO_STALE_TIME,
    });

    return (
        <div className='mainContainer'>
            {(!loading && !!user && !!profile) && <TopBar/>}

            <div className='container py-4' style={{ maxWidth: 520 }}>
                {(loading || isLoading) && <div className='text-center text-secondary'>Loading…</div>}
                {(!loading && !user) && <LoginPage/>}
                {(!loading && !!user && !isLoading && !profile) && <RegisterPage/>}
                {(!!user && !!profile) && <Outlet context={profile}/>}
            </div>

            <ErrorsComponent/>
        </div>
    )
}

export default MainPage;
