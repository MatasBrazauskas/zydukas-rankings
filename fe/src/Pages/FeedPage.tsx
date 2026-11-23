import { Outlet, useNavigate, useOutletContext } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";

import TopBar from "../Components/TopBar";
import LoginPage from "./LoginPage";
import RegisterPage from "./RegisterPage";
//import ErrorsComponent from "../Components/ErrorsComponent";
import ErrorsComponent from "../Components/ErrorsComponent";
import useAuthUser from "../Auth/useAuthUser";
import { getUser } from "../APIs/usersAPI";
import { getFeed } from "../APIs/postsAPI";
import { type PostDTO, type UserInfoDTO } from "../Utils/DTOs";
import {
    FEED_QUERY_KEY,
    FEED_STALE_TIME,
    USER_INFO_QUERY_KEY,
    USER_INFO_STALE_TIME,
} from "../Utils/queryConsts";

// The app shell, which MainPage used to be: the login / register / app gate
// plus the chrome every page sits in. `/` fills the Outlet with <Feed/> below;
// `/profile` and `/upload` replace it.
function FeedPage() {
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

// The feed itself: everyone's posts, newest first, and the way to add one.
export function Feed() {
    const profile = useOutletContext<UserInfoDTO>();
    const navigate = useNavigate();

    const { data: posts, isLoading } = useQuery({
        queryKey: [FEED_QUERY_KEY],
        queryFn: getFeed,
        staleTime: FEED_STALE_TIME,
    });

    return (
        <>
            <div className='d-flex align-items-center gap-3 mb-3'>
                <h1 className='h4 mb-0 me-auto'>Feed</h1>

                <button className='btn btn-primary' onClick={() => navigate('/upload')}>
                    Create a post
                </button>
            </div>

            {isLoading && <div className='text-center text-secondary'>Loading the feed…</div>}

            {(!isLoading && !posts?.length) && (
                <div className='card shadow-sm p-4 text-center text-secondary'>
                    Nothing here yet, {profile.username}. Post the first picture.
                </div>
            )}

            {posts?.map((post) => <PostCard key={post.id} post={post}/>)}
        </>
    )
}

function PostCard({ post }: { post: PostDTO }) {
    return (
        <article className='card shadow-sm mb-3'>
            <img className='card-img-top' src={post.image_url} alt={`Post by ${post.username}`}/>

            <div className='card-body py-2'>
                <div className='d-flex align-items-center gap-2'>
                    <span className='fw-semibold'>{post.username}</span>
                    <span className='text-secondary small'>{new Date(post.created_at).toLocaleDateString()}</span>
                </div>

                {/* Optional, and whitespace in it is the poster's own. */}
                {!!post.caption?.trim() && (
                    <p className='mb-0 mt-1' style={{ whiteSpace: 'pre-wrap' }}>{post.caption}</p>
                )}
            </div>
        </article>
    )
}

export default FeedPage;
