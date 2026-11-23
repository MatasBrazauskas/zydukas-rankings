import { NavLink } from "react-router-dom";

function TopBar() {
    return (
        <nav className='navbar navbar-expand bg-white border-bottom shadow-sm'>
            <div className='container'>
                <span className='navbar-brand fw-bold'>zydukas</span>

                <div className='navbar-nav me-auto'>
                    <NavLink className='nav-link' to='/' end>Feed</NavLink>
                    <NavLink className='nav-link' to='/profile'>Profile</NavLink>
                </div>
            </div>
        </nav>
    )
}

export default TopBar;
