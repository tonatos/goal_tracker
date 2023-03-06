import StartPointImage from '../../assets/StartPoint.png';
import FinishPointImage from '../../assets/FinishPoint.png';
import CursorIcon from '../../assets/CursorIcon.svg';


const StartPoint = ({}) => {
    return (
        <div>
            <img src={StartPointImage} />
        </div>
    )
}

const FinishPoint = ({}) => {
    return (
        <div>
            <img src={FinishPointImage} />
        </div>
    )
}

const Cursor = ({}) => {
    return (
        <div>
            <img src={CursorIcon} />
        </div>
    )
}

export const Timeline = ({ }) => {
    return (
        <>
            <StartPoint />
            <FinishPoint />
            <Cursor />
        </>
    )
}