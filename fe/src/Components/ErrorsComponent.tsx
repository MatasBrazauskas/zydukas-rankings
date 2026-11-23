import { useDispatch, useSelector } from "react-redux";

import { type RootState } from "../Store/store";
import { removeError } from "../Store/errorState";

function ErrorsComponent() {
    const { errors } = useSelector((state: RootState) => state.ERRORS_INFO);
    const dispatch = useDispatch();

    return (
        <div className='errorToasts'>
            {errors.map((error, i) => {
                return (
                    <div key={i} className='alert alert-danger d-flex align-items-center gap-2 shadow-sm py-2'>
                        <span className='me-auto small'>{error}</span>
                        <button type='button' className='btn-close' onClick={() => dispatch(removeError(i))}/>
                    </div>
                )
            })}
        </div>
    )
}

export default ErrorsComponent;
