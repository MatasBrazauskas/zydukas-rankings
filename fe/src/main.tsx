import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { MutationCache, QueryCache, QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider } from "react-redux";

import FeedPage, { Feed } from './Pages/FeedPage';
import UploadPage from './Pages/UploadPage';
import ProfilePage from './Pages/ProfilePage';
import store from './Store/store';
import { addError } from './Store/errorState';

import 'bootstrap/dist/css/bootstrap.min.css';
import './mainStyle.css';

const onError = (error: Error) => store.dispatch(addError(error.message));

export const queryClient = new QueryClient({
    queryCache: new QueryCache({ onError }),
    mutationCache: new MutationCache({ onError }),
    defaultOptions: {
        queries: {
            retry: false,
        }
    }
});

createRoot(document.getElementById('root')!).render(
    <QueryClientProvider client={queryClient}>
        <Provider store={store}>
            <BrowserRouter>
                <Routes>
                    <Route path='/' element={<FeedPage/>}>
                        <Route index element={<Feed/>}/>
                        <Route path='profile' element={<ProfilePage/>}/>
                        <Route path='upload' element={<UploadPage/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </Provider>
    </QueryClientProvider>
);
