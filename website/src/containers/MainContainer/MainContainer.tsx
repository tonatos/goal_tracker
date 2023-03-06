import Style from './MainContainer.module.css';
import { Timeline } from '../../components/Timeline';
import { AddContribution } from '../../components/AddContribution';

export const MainContainer = () => {
    return (
        <div className={Style.App}>
            <Timeline />
            <AddContribution />
        </div>
    )
}